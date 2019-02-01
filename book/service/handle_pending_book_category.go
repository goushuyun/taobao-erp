package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/goushuyun/taobao-erp/book/db"
	"github.com/goushuyun/taobao-erp/misc/bookspider"
	"github.com/goushuyun/taobao-erp/pb"

	"github.com/goushuyun/log"
)

var CateIsRunOn = false
var cate_data_length = 50

func HandlePendingBookCategory() {
	if !CateIsRunOn {
		// open the switch
		CateIsRunOn = true
		// get pending data
		pendingBooks, err := db.GetPendingGatherBooksCategory()
		if err != nil {
			log.Error(err)
			log.Debugf("发生错误，停留1分钟")
			time.Sleep(1 * time.Minute)
			CateIsRunOn = false
			HandlePendingBook()
		}
		// handle pending data
		if len(pendingBooks) <= 0 {
			log.Debugf("没数据，停留1分钟")
			time.Sleep(1 * time.Minute)
			CateIsRunOn = false
			HandlePendingBook()
		} else {
			if len(pendingBooks) <= data_length {
				coreBookCategoryPendingGatherHandler(pendingBooks)
			} else {
				var weight int
				mod := len(pendingBooks) / data_length

				if mod > 0 {
					weight = len(pendingBooks)/data_length + 1
				} else {
					weight = len(pendingBooks) / data_length
				}
				log.Debug(mod)
				log.Debug(weight)

				for i := 0; i < weight; i = i + data_length {
					var end int
					end = i + data_length
					if end >= len(pendingBooks) {
						end = len(pendingBooks)
					}
					log.Debugf("开始位置：%d 结束位置：%d", i, end)
					coreBookCategoryPendingGatherHandler(pendingBooks[i:end])
				}
			}

			CateIsRunOn = false
			HandlePendingBookCategory()
		}
	}
}

func coreBookCategoryPendingGatherHandler(pendingBooks []*pb.BookPendingGather) {
	var wg sync.WaitGroup
	statisticChan := make(chan int)
	for i := 0; i < len(pendingBooks); i++ {
		pendingBook := pendingBooks[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			category := bookspider.GetBookTaobaoCategory(pendingBook.Isbn)

			if category == "" {
				for index := 0; index < 2; index++ {
					var num int32
					num = rand.Int31n(3)
					log.Debugf("========停留秒数：%d", num)
					time.Sleep(time.Duration(num) * time.Second)
					category = bookspider.GetBookTaobaoCategory(pendingBook.Isbn)
					if category == "" {
						break
					}
				}
			}

			if category != "" {
				book := &pb.Book{}
				book.Id = pendingBook.BookId
				book.TaobaoCategory = category
				db.UpdateBookInfo(book)
				db.DelBookCategoryPendingGatherData(pendingBook)
				statisticChan <- 1
				return
			} else {
				if pendingBook.SearchTime >= 50 {
					db.DelBookCategoryPendingGatherData(pendingBook)
					statisticChan <- 1
					return
				}
				pendingBook.SearchTime = 1
				db.UpdateBookCategoryPendingGatherData(pendingBook)
				statisticChan <- 1
				return
			}

		}()

	}
	var currentCompleteNum int
	var singleValue int
	for {

		select {
		case singleValue, _ = <-statisticChan:
			currentCompleteNum += singleValue
		}

		if currentCompleteNum == len(pendingBooks) {
			fmt.Println(currentCompleteNum)
			close(statisticChan)
			//完成统计
			break
		}
	}

	wg.Wait()

	log.Debug("searchCategoryOver")

}
