package service

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/goushuyun/taobao-erp/book/constant"
	"github.com/goushuyun/taobao-erp/book/db"
	"github.com/goushuyun/taobao-erp/misc/bookspider"
	"github.com/pborman/uuid"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"

	"github.com/goushuyun/log"
)

type BookServer struct {
}

//获取图书信息
func (s *BookServer) GetBookInfo(ctx context.Context, in *pb.Book) (*pb.BookListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetBookInfo", "%#v", in))
	/*
	   check if need precision search by book'id if id not null ,just search book info from local db
	*/
	if in.Id != "" {
		// get book info from local db
		books, err := db.GetBookInfo(in)
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
		if len(books) <= 0 {
			return &pb.BookListResp{Code: errs.Ok, Message: "errParam"}, nil
		} else {
			return &pb.BookListResp{Code: errs.Ok, Message: "ok", Data: books}, nil
		}
	} else {
		if in.Isbn != "" {
			// first : get get book info from local db
			books, err := db.GetBookInfo(in)
			if err != nil {
				log.Error(err)
				return nil, errs.Wrap(errors.New(err.Error()))
			}
			if len(books) > 0 {
				return &pb.BookListResp{Code: errs.Ok, Message: "ok", Data: books}, nil
			}
			// second :if local db don't has this book info ,just get it from internet (dangdang ,jd ,book uu)
			// this function distinguish the upload mode: default or speed
			book, err := insertByUploadMode(in.Isbn, in.UploadMode)
			if err != nil {
				log.Error(err)
				return nil, errs.Wrap(errors.New(err.Error()))
			}
			bookresp := &pb.BookListResp{Code: errs.Ok, Message: "ok"}
			bookresp.Data = append(bookresp.Data, book)
			return bookresp, nil
		}
	}
	return &pb.BookListResp{Code: errs.Ok, Message: "errParam"}, nil
}

//获取图书信息
func (s *BookServer) GetLocalBookInfo(ctx context.Context, in *pb.Book) (*pb.BookListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocalBookInfo", "%#v", in))

	// get book info from local db
	books, err := db.GetBookInfo(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	if len(books) <= 0 {
		return &pb.BookListResp{Code: errs.Ok, Message: "errParam"}, nil
	}
	return &pb.BookListResp{Code: errs.Ok, Message: "ok", Data: books}, nil

}

//change the book info
func (s *BookServer) UpdateBookInfo(ctx context.Context, in *pb.Book) (*pb.BookResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "UpdateBookInfo", "%#v", in))
	updateContent, err := db.UpdateBookInfo(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	log.Debug(updateContent)
	return &pb.BookResp{Code: errs.Ok, Message: "ok"}, nil
}

//insert new data to book
func (s *BookServer) SaveBook(ctx context.Context, in *pb.Book) (*pb.BookResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveBook", "%#v", in))
	book_no := "00"
	if in.BookCate == "poker" {
		//首先查看有没有未分配的book_no
		books, err := db.GetBookInfo(&pb.Book{Isbn: in.Isbn})
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
		for i := 0; i < len(books); i++ {
			book := books[i]
			if book.BookCate == "" {
				book.BookCate = "poker"
				if book.BookNo == "" {
					book.BookNo = book_no
					book_no = constant.FindNextNo(book_no)
				} else {
					book_no = constant.FindNextNo(book.BookNo)
				}
				db.UpdateBookInfo(book)
			} else {
				book_no = constant.FindNextNo(book.BookNo)
			}

		}

	}
	in.BookNo = book_no
	err := db.InsertBookInfo(in)
	if err != nil {
		//check the err reason if equal 'exists' in particular.if yes ,return specially identification str
		if err.Error() == "exists" {
			return &pb.BookResp{Code: errs.Ok, Message: "exists", Data: in}, nil
		}
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.BookResp{Code: errs.Ok, Message: "ok", Data: in}, nil
}

/*
	private function: handle the book info
	 1: download and upload the book image to qiniu
	 2: handle the book title
*/
func handleBookInfos(book *pb.Book, ctx context.Context) error {
	t := time.Now()
	timestamp := t.Format("20060102030405")
	if strings.HasPrefix(book.Image, "http") {
		fetchImageReq := &pb.FetchImageReq{
			Zone: pb.MediaZone_Test,
			Url:  book.Image,
			Key:  timestamp + book.Isbn + filepath.Ext(book.Image),
		}
		mediaResp := &pb.FetchImageResp{}
		err := misc.CallSVC(ctx, "mediastore", "FetchImage", fetchImageReq, mediaResp)
		if err != nil {
			log.Error(err)
			return err
		}
		book.Image = fetchImageReq.Key
	}
	return nil
}

/*
	private function: in order to resolve the condition that user wantna upload book quick
	 if upload mode is 1 , so speed upload ,just omit the wait time about the search book info on internet
	 if upload mode is 0, wait until search over
*/

func insertByUploadMode(isbn string, uploadMode int64) (book *pb.Book, err error) {
	ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", uuid.New()))
	if uploadMode == 0 {
		book, err = bookspider.GetBookInfoBySpider(isbn, "")
		if err != nil {
			log.Error(err)
			return
		}
		if book == nil {
			book, err = bookspider.GetBookInfoBySpider(isbn, "")
			if err != nil {
				log.Error(err)
				return
			}
		}
		if book != nil {
			err = handleBookInfos(book, ctx) //handle the book info
			if err != nil {
				log.Error(err)
				return
			}
		} else {
			//if book is not found from internet just init a book struct with one field 'isbn'
			book = &pb.Book{Isbn: isbn}
		}
		//finally : insert a new data and return
		err = db.InsertBookInfo(book)
		if err != nil {
			log.Error(err)
			return
		}
		//爬取图书的分类
		go db.InsertBookCategoryPendingGatherData(&pb.BookPendingGather{Id: book.Id, Source: "taobao"})
		return
	} else {
		book = &pb.Book{Isbn: isbn}
		err = db.InsertBookInfo(book)
		if err != nil {
			log.Error(err)
			return
		}
		go func() {
			db.InsertBookPendingGatherData(&pb.BookPendingGather{BookId: book.Id})
			//爬取图书的分类
			db.InsertBookCategoryPendingGatherData(&pb.BookPendingGather{BookId: book.Id, Source: "taobao"})
		}()

	}

	return
}
