package service

import (
	"context"
	"testing"

	"github.com/goushuyun/taobao-erp/misc/bookspider"
	"github.com/pborman/uuid"
	"google.golang.org/grpc/metadata"

	"github.com/goushuyun/log"
)

func TestHandleBook(t *testing.T) {
	ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", uuid.New()))

	book, err := bookspider.GetBookInfoBySpider("9787544270878", "")
	if err != nil {
		log.Error(err)
	}

	t.Log(book)

	if book != nil {
		err = handleBookInfos(book, ctx) //handle the book info
		if err != nil {
			log.Error(err)
		}
		log.Debug(book)
	}
}
