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
		if book.BookCate != "" {
			condition += fmt.Sprintf(" and book_cate='%s'", book.BookCate)
		}
		if book.BookNo != "" {
			condition += fmt.Sprintf(" and book_no='%s'", book.BookNo)
		}
	}
	if book.Id != "" {
		condition += fmt.Sprintf(" and id='%s'", book.Id)
	}

	condition += " order by id"
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

//change the book info
func UpdateBookInfo(book *pb.Book) (updateContent string, err error) {
	query := "update book set update_at=now()"
	var condition string

	if book.BookNo != "" {
		condition += fmt.Sprintf(",book_no='%s'", book.BookNo)
		updateContent += fmt.Sprintf(" 图书编号：'%s'", book.BookNo)
	}
	if book.BookCate != "" {
		condition += fmt.Sprintf(",book_cate='%s'", book.BookCate)
		updateContent += fmt.Sprintf(" 图书类型：'%s'", book.BookCate)
	}
	if book.Title != "" {
		condition += fmt.Sprintf(",title='%s'", book.Title)
		updateContent += fmt.Sprintf(" 书名：'%s'", book.Title)
	}
	if book.Publisher != "" {
		condition += fmt.Sprintf(",publisher='%s'", book.Publisher)
		updateContent += fmt.Sprintf(" 出版社：'%s'", book.Publisher)
	}
	if book.Author != "" {
		condition += fmt.Sprintf(",author='%s'", book.Author)
		updateContent += fmt.Sprintf(" 作者：'%s'", book.Author)
	}
	if book.Edition != "" {
		condition += fmt.Sprintf(",edition='%s'", book.Edition)
		updateContent += fmt.Sprintf(" 版本：'%s'", book.Edition)
	}
	if book.Pubdate != "" {
		condition += fmt.Sprintf(",pubdate='%s'", book.Pubdate)
		updateContent += fmt.Sprintf(" 出版日期：'%s'", book.Pubdate)
	}
	if book.SeriesName != "" {
		condition += fmt.Sprintf(",series_name='%s'", book.SeriesName)
		updateContent += fmt.Sprintf(" 丛书名：'%s'", book.SeriesName)
	}
	if book.Image != "" {
		condition += fmt.Sprintf(",image='%s'", book.Image)
		updateContent += fmt.Sprintf(" 图片地址：'%s'", book.Image)
	}
	if book.Price != 0 {
		condition += fmt.Sprintf(",price=%d", book.Price)
		updateContent += fmt.Sprintf(" 价格：%d", book.Price)
	}
	if book.Catalog != "" {
		condition += fmt.Sprintf(",catalog='%s'", book.Catalog)
		updateContent += fmt.Sprintf(" 目录：'%s'", book.Catalog)
	}
	if book.Abstract != "" {
		condition += fmt.Sprintf(",abstract='%s'", book.Abstract)
		updateContent += fmt.Sprintf(" 书本简介：'%s'", book.Abstract)
	}
	if book.Page != "" {
		condition += fmt.Sprintf(",page='%s'", book.Page)
		updateContent += fmt.Sprintf(" 页数：'%s'", book.Page)
	}
	if book.Packing != "" {
		condition += fmt.Sprintf(",packing='%s'", book.Packing)
		updateContent += fmt.Sprintf(" 包装：'%s'", book.Packing)
	}
	if book.AuthorIntro != "" {
		condition += fmt.Sprintf(",author_intro='%s'", book.AuthorIntro)
		updateContent += fmt.Sprintf(" 作者简介：'%s'", book.AuthorIntro)
	}
	if book.SourceInfo != "" {
		condition += fmt.Sprintf(",source_info='%s'", book.SourceInfo)
		updateContent += fmt.Sprintf(" 书本来源：'%s'", book.SourceInfo)
	}
	condition += fmt.Sprintf(" where id='%s'", book.Id)
	if updateContent == "" {
		err = errors.New("没任何信息更新呦～")
		return
	}
	query += condition
	log.Debug(query)
	_, err = DB.Exec(query)
	if err != nil {
		log.Error(err)
	}
	return
}
