package service

import (
	"errors"
	"goushuyun/books/db"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"

	"github.com/wothing/log"
)

type BookServer struct {
}

//获取图书信息
func (s *BookServer) GetBookInfo(ctx context.Context, in *pb.Book) (*pb.BookResp, error) {
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
			return &pb.BookResp{Code: errs.Ok, Message: "errParam"}, nil
		} else {
			return &pb.BookResp{Code: errs.Ok, Message: "ok", Data: book}, nil
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
				return &pb.BookResp{Code: errs.Ok, Message: "ok", Data: book}, nil
			}
			// second :if local db don't has this book info ,just get it from internet (dangdang ,jd ,book uu)
			GetBookInfoBySpider(in.Isbn, "")
			//finally : insert a new data just has one field "isbn" and return

		}
	}
	return &pb.BookResp{Code: errs.Ok, Message: "errParam"}, nil
}

//更改图书信息
func (s *BookServer) UpdateBookInfo(ctx context.Context, in *pb.Book) (*pb.BookResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "UpdateBookInfo", "%#v", in))

	return &pb.BookResp{Code: errs.Ok, Message: "ok"}, nil
}

//管理员新增图书信息
func (s *BookServer) SaveBook(ctx context.Context, in *pb.Book) (*pb.BookResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveBook", "%#v", in))

	return &pb.BookResp{Code: errs.Ok, Message: "ok"}, nil
}
