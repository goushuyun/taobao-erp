package db

import (
	"fmt"

	"github.com/wothing/log"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
)

/*
   save the book audit
*/
func SaveBookAudit(record *pb.BookAuditRecord) error {
	query := "insert into book_audit_record(book_id,isbn,book_cate,title,publisher,author,edition,image,price,apply_user_id,apply_user_name) values(%s) return id ,extract(epoch from create_at)::bigint"
	var condition string
	condition += fmt.Sprintf("'%s','%s','%s','%s','%s','%s','%s','%s',%d,'%s','%s'", record.BookId, record.Isbn, record.BookCate, record.Title, record.Publisher, record.Author, record.Edition, record.Image, record.Price, record.ApplyUserId, record.ApplyUserName)
	query = fmt.Sprintf(query, condition)
	log.Debug(query)
	err := DB.QueryRow(query).Scan(&record.Id, &record.CreateAt)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

/*
   get book audit list for Auditor
*/
func GetBookAuditList(record *pb.BookAuditRecord) {
	query := "select id,book_id,isbn,book_cate,title,publisher,author,edition,image,price,apply_user_id,apply_user_name,check_user_id,check_user_name,apply_reason,status,feedback,extract(epoch from create_at)::bigint from book_audit_record where 1=1"
	queryCount := "select count(*) from book_audit_record where 1=1"
	var condition string
	// condition 1 : apply_user_id
	if record.ApplyUserId != "" {
		condition += fmt.Sprintf(" and apply_user_id='%s'", record.ApplyUserId)
	}
	// condition 2 : check_user_id
	if record.CheckUserId != "" {
		condition += fmt.Sprintf(" and check_user_id='%s'", record.CheckUserId)
	}
	// condition 3 : id
	if record.Id != "" {
		condition += fmt.Sprintf(" and id='%s'", record.Id)
	}
	// condition 4 : book_id
	if record.BookId != "" {
		condition += fmt.Sprintf(" and book_id='%s'", record.BookId)
	}
	// condition 5 : isbn
	if record.Isbn != "" {

	}
	// condition 6 : status
	if record.Status != 0 {
		condition += fmt.Sprintf(" and status=%d", record.Status)
	}
	queryCount += condition
	log.Debug(queryCount)
	query += condition
	log.Debug(query)

}

/*
   update book audit
*/
func UpdateBookAudit() {

}
