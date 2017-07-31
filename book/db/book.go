package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

//get book info from db
func GetBookInfo(book *pb.Book) (books []*pb.Book, err error) {
	query := "select id,isbn,book_no,book_cate,title,publisher,author,edition,pubdate,series_name,image,price,catalog,abstract,page,packing,format,author_intro,source_info,extract(epoch from create_at)::bigint,extract(epoch from update_at)::bigint from book where 1=1"
	var condition string
	if book.Isbn != "" {
		condition += fmt.Sprintf(" and isbn='%s'", book.Isbn)
	}
	if book.Id != "" {
		condition += fmt.Sprintf(" and id='%s'", book.Id)
	}
	condition += " order by id desc"
	query += condition
	log.Debug(query)
	rows, err := DB.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Error(err)
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		model := &pb.Book{}
		books = append(books, model)
		err = rows.Scan(&model.Id, &model.Isbn, &model.BookNo, &model.BookCate, &model.Title, &model.Publisher, &model.Author, &model.Edition, &model.Pubdate, &model.SeriesName, &model.Image, &model.Price, &model.Catalog, &model.Abstract, &model.Page, &model.Packing, &model.Format, &model.AuthorIntro, &model.SourceInfo, &model.CreateAt, &model.UpdateAt)
		if err != nil {
			log.Error(err)
			return
		}
	}
	return
}

/*
   insert a new book data to db and return it's id where complete`	r
*/
func InsertBookInfo(book *pb.Book) error {
	query := "select count(*) from book where isbn='%s' and book_no='%s' and book_cate='%s'"
	query = fmt.Sprintf(query, book.Isbn, book.BookNo, book.BookCate)
	log.Debug(query)
	var totalCount int64
	err := DB.QueryRow(query).Scan(&totalCount)
	if err != nil {
		log.Error(err)
		return err
	}
	if totalCount > 0 {
		return errors.New("exists")
	}
	query = "insert into book(isbn,book_no,book_cate,title,publisher,author,edition,pubdate,series_name,image,price,catalog,abstract,page,packing,format,author_intro,source_info) values(%s) returning id,extract(epoch from create_at)::bigint"
	param := fmt.Sprintf("'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%d,'%s','%s','%s','%s','%s','%s','%s'", book.Isbn, book.BookNo, book.BookCate, book.Title, book.Publisher, book.Author, book.Edition, book.Pubdate, book.SeriesName, book.Image, book.Price, book.Catalog, book.Abstract, book.Page, book.Packing, book.Format, book.AuthorIntro, book.SourceInfo)
	query = fmt.Sprintf(query, param)
	log.Debug(query)
	err = DB.QueryRow(query).Scan(&book.Id, &book.CreateAt)
	if err != nil {
		log.Error(err)
		return err
	}
	book.UpdateAt = book.CreateAt
	return nil
}
