package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	mediastore "github.com/goushuyun/taobao-erp/mediastore/service"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/misc/bookspider"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/db"
	"github.com/goushuyun/taobao-erp/stock/export"
	"github.com/mholt/archiver"
	"github.com/wothing/log"
)

var IsRunOn = false

func HandlePendingExportedTaobaoCsv() {
	if !IsRunOn {
		// open the switch
		IsRunOn = true
		// get pending data
		records, err, _ := db.GetTaobaoCsvExportRecord(&pb.TaobaoCsvRecord{Page: 1, Size: 2, Status: 1})
		if err != nil {
			log.Error(err)
			log.Debugf("发生错误，停留1分钟")
			time.Sleep(1 * time.Minute)
			IsRunOn = false
			HandlePendingExportedTaobaoCsv()
		}
		// handle pending data
		if len(records) <= 0 {
			log.Debugf("没数据，停留1分钟")
			time.Sleep(1 * time.Minute)
			IsRunOn = false
			HandlePendingExportedTaobaoCsv()
		} else {
			IsRunOn = false
			corePendingExportCsvHandler(records)
			HandlePendingExportedTaobaoCsv()

		}
	}
}

func corePendingExportCsvHandler(records []*pb.TaobaoCsvRecord) {
	//ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", uuid.New()))
	var wg sync.WaitGroup
	statisticChan := make(chan int)
	for i := 0; i < len(records); i++ {
		record := records[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				statisticChan <- 1
			}()
			//获取所有数据
			items, err := db.GetTaobaoCsvExportRecordItems(record)
			if err != nil {
				log.Error(err)
				return
			}
			dateStr := time.Now().Format("200601021504")
			rootPath := "/var/export/bookcloud/" + record.UserId + "/" + dateStr + "/"
			log.Debug(rootPath)
			imagePath := rootPath + dateStr + "/"
			err = os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				log.Error(err)
				return
			}
			err = os.MkdirAll(imagePath, os.ModePerm)
			if err != nil {
				log.Error(err)
				return
			}
			filepath := rootPath + dateStr + ".csv"

			defer os.Remove(rootPath)

			var taobaoModels = make([]*export.TaobaoCsvModel, 0, len(items))
			for i := 0; i < len(items); i++ {
				var isbn, category, title, image, province, city, describe, reduceStockStyle, deliveryFeeTemplete string
				var stock int64
				var price, pingyou_fee, ems_fee, express_fee float64
				item := items[i]
				// before packing taobao csv necessary field,complete the field first
				// 1  mkdir  user_id/date/

				// 2 complete the toabao category
				// if taobao category is empty , then get the category from search taobao
				category = item.TaobaoCategory
				if category == "" {
					category = bookspider.GetBookTaobaoCategory(item.Isbn)
					if category == "" {
						category = bookspider.GetBookTaobaoCategory(item.Isbn)
					}
					if category == "" {
						category = "50050324"
					}
				}

				// 2 download the picture
				image = ""
				if item.Image != "" {
					rawURL := "http://taoimage.goushuyun.cn/" + item.Image
					imageName := time.Now().Format("20060102150405")
					imageName = fmt.Sprintf(imageName+"-%d_o_%d", i, i)
					imageFilePath := imagePath + imageName + ".tbi"
					err = misc.DownloadFileFromServer(imageFilePath, rawURL)
					if err == nil {
						image = imageName + ":0:0:;"
					}
				}
				// 3 packing title
				title = record.ProductTitle
				title = strings.Replace(title, "{{title}}", item.Title, -1)
				title = strings.Replace(title, "{{publisher}}", item.Publisher, -1)
				title = strings.Replace(title, "{{isbn}}", item.Isbn, -1)
				title = strings.Replace(title, "{{author}}", item.Author, -1)

				// 4 packing describe
				if item.AuthorIntro == "" {
					item.AuthorIntro = "暂无作者简介"
				} else {
					count := strings.Count(item.AuthorIntro, "") - 1
					if count > 4000 {
						item.AuthorIntro = "暂无作者简介"
					}
				}
				if item.Catalog == "" {
					item.Catalog = "暂无目录"
				} else {
					count := strings.Count(item.Catalog, "") - 1
					if count > 4000 {
						item.Catalog = "暂无目录"
					}
				}
				if item.Abstract == "" {
					item.Abstract = "暂无简介"
				} else {
					count := strings.Count(item.Abstract, "") - 1
					if count > 4000 {
						item.Abstract = "暂无简介"
					}
				}
				describe = record.ProductDescribe
				describe = strings.Replace(describe, "{{isbn}}", item.Isbn, -1)
				describe = strings.Replace(describe, "{{publisher}}", item.Publisher, -1)
				describe = strings.Replace(describe, "{{author}}", item.Author, -1)
				describe = strings.Replace(describe, "{{edition}}", item.Edition, -1)
				describe = strings.Replace(describe, "{{pubdate}}", item.Pubdate, -1)
				describe = strings.Replace(describe, "{{page}}", item.Page, -1)
				describe = strings.Replace(describe, "{{packing}}", item.Packing, -1)
				describe = strings.Replace(describe, "{{format}}", item.Format, -1)
				describe = strings.Replace(describe, "{{catalog}}", item.Catalog, -1)
				describe = strings.Replace(describe, "{{abstract}}", item.Abstract, -1)
				describe = strings.Replace(describe, "{{author_intro}}", item.AuthorIntro, -1)

				//"9787040231250", "50000182", "离散数学", "21043943-1_o_2:0:0:;", "北京", "北京", describe, 2, 10, 7, 8, 9
				//isbn, category, title, image, province, city, describe , reduceStockStyle, deliveryFeeTemplete string, stock int64, price, pingyou_fee, ems_fee, express_fee float64
				isbn = item.Isbn
				province = record.Province
				city = record.City
				stock = item.Stock
				var serviceDiscount = float64(record.Discount) / 100
				price = float64(item.Price) / 100 * (1.00 - serviceDiscount)
				price = floatTrans(price)
				if err != nil {
					price = 0
				}
				pingyou_fee = float64(record.PingyouFee) / 100
				ems_fee = float64(record.EmsFee) / 100
				express_fee = float64(record.ExpressFee) / 100
				reduceStockStyle = record.ReduceStockStyle
				deliveryFeeTemplete = record.ExpressTemplate
				// category title  image describe isbn province,city stock,price pingyou_fee ems_fee express_fee
				taobaoModel := export.PackingTaobaoParam(isbn, category, title, item.Title, image, province, city, describe, reduceStockStyle, deliveryFeeTemplete, stock, price, pingyou_fee, ems_fee, express_fee)
				taobaoModels = append(taobaoModels, taobaoModel)
			}
			log.Debug(taobaoModels)

			file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
			defer file.Close()
			//if only file implements the io.Writer interface
			var parser = export.NewCsv(file)
			err = parser.Parse(taobaoModels)
			if err != nil {
				log.Error(err)
				record.Status = 3
				record.Summary = "生成文件失败"
				exportResult(record)
				return
			}
			export.SetReadOnly(filepath)

			// create rar file
			zipName := "taobao" + dateStr + ".zip"
			zipPath := rootPath + zipName
			err = archiver.Zip.Make(zipPath, []string{filepath, imagePath})
			if err != nil {
				record.Status = 3
				record.Summary = "压缩文件失败"
				exportResult(record)
				log.Error(err)
				return
			}
			//filepath string, zone pb.MediaZone, key string
			err = mediastore.UploadLocalFile(zipPath, 0, zipName)
			if err != nil {
				record.Status = 3
				record.Summary = "上传文件失败"
				exportResult(record)
				log.Error(err)
				return
			}
			record.FileUrl = zipName
			record.Status = 2
			err = exportResult(record)
			if err != nil {
				log.Error(err)
				return
			}

			return
		}()

	}
	var currentCompleteNum int
	var singleValue int
	for {

		select {
		case singleValue, _ = <-statisticChan:
			currentCompleteNum += singleValue
		}

		if currentCompleteNum == len(records) {
			fmt.Println(currentCompleteNum)
			close(statisticChan)
			//完成统计
			break
		}
	}

	wg.Wait()

	log.Debug("uploadOver")

}

func exportResult(record *pb.TaobaoCsvRecord) error {

	err := db.UpdatTaobaoCsvExportRecordItems(record)
	if err != nil {
		log.Error(err)
		return err
	}
	err = db.DelTaobaoCsvExportRecordItems(record)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func floatTrans(value float64) float64 {
	str := fmt.Sprintf("%.2f", value)
	value, _ = strconv.ParseFloat(str, 64)
	return value
}
