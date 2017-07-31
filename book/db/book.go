package db

import (
	"database/sql"
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
