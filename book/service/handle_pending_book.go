package service

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/goushuyun/taobao-erp/book/db"
	"github.com/goushuyun/taobao-erp/misc/bookspider"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/pborman/uuid"

	"github.com/wothing/log"
)

var IsRunOn = false

/*
	first,find the book that info is complete
	then ,push the book to the table book_pending_gather
*/
func GetAllIncompleteBook() {
	books, err := db.GetIncompleteBook()
	if err != nil {
		return
	}
	if books != nil && len(books) > 1 {
		for i := 0; i < len(books); i++ {
			book := books[i]
			book.SearchTime = 1
			_, err = db.UpdateBookInfo(book)
			if err != nil {
				log.Error(err)
				continue
			}
			err = db.InsertBookPendingGatherData(&pb.BookPendingGather{BookId: book.Id})
			if err != nil {
				log.Error(err)
				continue
			}

		}
	}
}

func HandlePendingBook() {
	if !IsRunOn {
		// open the switch
		IsRunOn = true
		// get pending data
		pendingBooks, err := db.GetPendingGatherBooks()
		if err != nil {
			log.Error(err)
			log.Debugf("发生错误，停留1分钟")
			time.Sleep(1 * time.Minute)
			IsRunOn = false
			HandlePendingBook()
		}
		// handle pending data
		if len(pendingBooks) <= 0 {
			log.Debugf("没数据，停留1分钟")
			time.Sleep(1 * time.Minute)
			IsRunOn = false
			HandlePendingBook()
		} else {
			if len(pendingBooks) <= 100 {
				corePendingGatherHandler(pendingBooks)
			} else {
				var weight int
				mod := len(pendingBooks) / 100

				if mod > 0 {
					weight = len(pendingBooks)/100 + 1
				} else {
					weight = len(pendingBooks) / 100
				}
				log.Debug(mod)
				log.Debug(weight)

				for i := 0; i < weight; i = i + 100 {
					var end int
					end = i + 100
					if end >= len(pendingBooks) {
						end = len(pendingBooks)
					}
					log.Debugf("开始位置：%d 结束位置：%d", i, end)
					corePendingGatherHandler(pendingBooks[i:end])
				}
			}

			IsRunOn = false
			HandlePendingBook()
		}
	}
}

func corePendingGatherHandler(pendingBooks []*pb.BookPendingGather) {
	ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", uuid.New()))
	var wg sync.WaitGroup
	statisticChan := make(chan int)
	for i := 0; i < len(pendingBooks); i++ {
		pendingBook := pendingBooks[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			book, err := bookspider.GetBookInfoBySpider(pendingBook.Isbn, "")
			if book == nil {
				for index := 0; index < 2; index++ {
					var num int32
					num = rand.Int31n(3)
					log.Debugf("========停留秒数：%d", num)
					time.Sleep(time.Duration(num) * time.Second)
					book, err = bookspider.GetBookInfoBySpider(pendingBook.Isbn, "")
					if book != nil {
						break
					}
				}
			}
			if err != nil {
				log.Error(err)
				statisticChan <- 1
				return
			}
			if book != nil {
				err = handleBookInfos(book, ctx) //handle the book info
				if err != nil {
					log.Error(err)
					statisticChan <- 1
					return
					//continue
				}
				book.Id = pendingBook.BookId
				db.UpdateBookInfo(book)
				db.DelBookPendingGatherData(pendingBook)
				statisticChan <- 1
				return
			} else {
				if pendingBook.SearchTime >= 70 {
					db.DelBookPendingGatherData(pendingBook)
					statisticChan <- 1
					return
				}
				pendingBook.SearchTime = 1
				db.UpdateBookPendingGatherData(pendingBook)
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

	log.Debug("uploadOver")

}
