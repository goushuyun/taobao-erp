package bookspider

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/wothing/log"

	"github.com/goushuyun/taobao-erp/misc/key_word_filter"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

// "github.com/goushuyun/weixin-golang/misc/bookspider"

//通过爬虫获取图书信息
func GetBookInfoBySpider(isbn, upload_way string) (book *pb.Book, err error) {
	book = &pb.Book{}
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)
	var num int32
	if upload_way == "batch" {
		num = rand.Int31n(3)
	} else {
		num = rand.Int31n(1)
	}
	log.Debugf("===上传类型：%s========停留秒数：%d", upload_way, num)
	time.Sleep(time.Duration(num) * time.Second)
	ip := getProxyIp()

	//首先从当当上获取图书信息
	book.SourceInfo = "dangdang"
	sp := spider.NewSpider(NewDangDangListProcesser(), "spiderDangDangList")
	baseURL := "http://search.dangdang.com/?key=ISBN&ddsale=1"
	url := strings.Replace(baseURL, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	log.Debug(url)
	if ip != "" {
		req.AddProxyHost(ip)
	}

	pageItems := sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("dangdang no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}

	//从京东上获取图书信息
	book.SourceInfo = "jd"
	sp = spider.NewSpider(NewJDListProcesser(), "spiderJDList")
	baseURL = "https://search.jd.com/Search?keyword=ISBN&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=ISBN&psort=1&wtype=1&click=1"
	url = strings.Replace(baseURL, "ISBN", isbn, -1)
	req = request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	if ip != "" {
		req.AddProxyHost(ip)
	}
	pageItems = sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("jd no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}

	//如果当当图书信息为空 从bookUU上获取数据
	book.SourceInfo = "bookUU"
	sp = spider.NewSpider(NewBookUUListProcesser(), "BookUUlist")
	baseURL = "http://search.bookuu.com/AdvanceSearch.php?isbn=ISBN&sm=&zz=&cbs=&dj_s=&dj_e=&bkj_s=&bkj_e=&layer2=&zk=0&cbrq_n=2017&cbrq_y=&cbrq_n1=2017&cbrq_y1=&sjsj=0&orderby=&layer1=1"
	url = strings.Replace(baseURL, "ISBN", isbn, -1)
	req = request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	if ip != "" {
		req.AddProxyHost(ip)
	}

	pageItems = sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("bookuu no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}
	return nil, nil
}

//构建返回结果
func structData(items *page_items.PageItems, book *pb.Book) {
	//获取数据
	title, _ := items.GetItem("title")
	author, _ := items.GetItem("author")
	publisher, _ := items.GetItem("publisher")
	pubdate, _ := items.GetItem("pubdate")
	price, _ := items.GetItem("price")
	isbn, _ := items.GetItem("isbn")
	edition, _ := items.GetItem("edition")
	image_url, _ := items.GetItem("image_url")
	catalog, _ := items.GetItem("catalog")
	abstract, _ := items.GetItem("abstract")
	series_name, _ := items.GetItem("series_name")
	author_info, _ := items.GetItem("author_info")
	page, _ := items.GetItem("page")
	packing, _ := items.GetItem("packing")
	format, _ := items.GetItem("format")

	// 处理数据
	priceFloat, _ := strconv.ParseFloat(price, 64)
	priceFloat = priceFloat * 100

	// 封装数据
	// 过滤数据
	title = key_word_filter.FilterKeyWords(title)
	book.Title = title
	book.Isbn = isbn
	book.Price = int64(priceFloat)
	book.Author = author
	book.Publisher = publisher
	book.Pubdate = pubdate
	book.Image = image_url
	book.Edition = edition
	book.Catalog = catalog
	book.Abstract = abstract
	book.SeriesName = series_name
	book.AuthorIntro = author_info
	book.Page = page
	book.Packing = packing
	book.Format = format
	return
}
