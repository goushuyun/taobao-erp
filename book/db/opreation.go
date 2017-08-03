package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/wothing/log"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
)

/*
   save the book audit
*/
func SaveBookAudit(record *pb.BookAuditRecord) error {
	query := "insert into book_audit_record(book_id,isbn,book_cate,title,publisher,author,edition,image,price,apply_user_id,apply_user_name) values(%s) returning id ,extract(epoch from create_at)::bigint"
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
func GetBookAuditList(record *pb.BookAuditRecord) (records []*pb.BookAuditRecord, err error, totalCount int64) {
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
		condition += fmt.Sprintf(" and isbn='%s'", record.Isbn)
	}

	// condition 6 : status
	if record.Status != 0 {
		condition += fmt.Sprintf(" and status=%d", record.Status)
	}
	//condition 7 : book_cate
	if record.BookCate != "" {
		condition += fmt.Sprintf(" and book_cate='%s'", record.BookCate)
	}
	queryCount += condition
	log.Debug(queryCount)

	// find out how rows tyey are
	err = DB.QueryRow(queryCount).Scan(&totalCount)
	if err != nil {
		log.Error(err)
		return
	}
	if totalCount <= 0 {
		return
	}
	if record.Page <= 0 {
		record.Page = 1
	}
	if record.Size <= 0 {
		record.Size = 15
	}
	// ergodic the rows
	condition += fmt.Sprintf(" order by create_at,id offset %d limit %d", (record.Page-1)*record.Size, record.Size)
	query += condition
	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		model := &pb.BookAuditRecord{}
		records = append(records, model)
		//id,book_id,isbn,book_cate,title,publisher,author,edition,image,price,apply_user_id,apply_user_name,check_user_id,check_user_name,apply_reason,status,feedback,extract(epoch from create_at)::bigint
		err = rows.Scan(&model.Id, &model.BookId, &model.Isbn, &model.BookCate, &model.Title, &model.Publisher, &model.Author, &model.Edition, &model.Image, &model.Price, &model.ApplyUserId, &model.ApplyUserName, &model.CheckUserId, &model.CheckUserName, &model.ApplyReason, &model.Status, &model.Feedback, &model.CreateAt)
		if err != nil {
			log.Error(err)
			return
		}
	}

	return

}

/*
   update book audit
*/
func UpdateBookAudit(record *pb.BookAuditRecord) error {
	query := "update book_audit_record set update_at=now()"
	var condition string
	if record.CheckUserId != "" {
		condition += fmt.Sprintf(",check_user_id='%s'", record.CheckUserId)
	}
	if record.CheckUserName != "" {
		condition += fmt.Sprintf(",check_user_name='%s'", record.CheckUserName)
	}
	if record.Status != 0 {
		condition += fmt.Sprintf(",status=%d", record.Status)
	}
	if record.Feedback != "" {
		condition += fmt.Sprintf(",feedback='%s'", record.Feedback)
	}
	condition += fmt.Sprintf(" where id='%s'", record)
	query += condition
	log.Debug(query)
	result, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	rowsCount, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	if rowsCount <= 0 {
		return errors.New("没修改任何数据呦～")
	}
	return nil
}

/*
   update book audit
*/
func BatchUpdateBookAudit(record *pb.BookAuditRecord) error {
	query := "update book_audit_record set update_at=now()"
	var condition string
	if record.CheckUserId != "" {
		condition += fmt.Sprintf(",check_user_id='%s'", record.CheckUserId)
	}
	if record.CheckUserName != "" {
		condition += fmt.Sprintf(",check_user_name='%s'", record.CheckUserName)
	}
	if record.Status != 0 {
		condition += fmt.Sprintf(",status=%d", record.Status)
	}
	if record.Feedback != "" {
		condition += fmt.Sprintf(",feedback='%s'", record.Feedback)
	}

	stmt := strings.Repeat(",'%s'", len(record.Ids))
	var ids []interface{}
	for _, value := range record.Ids {
		ids = append(ids, value)
	}

	stmt = fmt.Sprintf(stmt, ids...)
	stmt = string([]byte(stmt)[1:])
	condition += fmt.Sprintf(" where id in(%s)", stmt)

	query += condition
	log.Debug(query)
	result, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	rowsCount, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	if rowsCount <= 0 {
		return errors.New("没修改任何数据呦～")
	}
	return nil
}

// get book audit organize list
func GetOrganizedBookAuditList(record *pb.BookAuditRecord) (models []*pb.OrganizedBookAudit, err error, totalCount int64) {

	var query, queryCount string
	if record.Page <= 0 {
		record.Page = 1
	}
	if record.Size <= 0 {
		record.Size = 15
	}
	if record.SearchType == 0 {
		// update
		query = fmt.Sprintf("select book_id,count(*) from book_audit_record where status=1 group by book_id order by count(*) desc,book_id offset %d limit %d", (record.Page-1)*record.Size, record.Size)
		queryCount = "select count(distinct book_id) from book_audit_record where status=1"
		log.Debug(queryCount)
		err = DB.QueryRow(queryCount).Scan(&totalCount)
		if err != nil {
			log.Error(err)
			return
		}
		if totalCount <= 0 {
			return
		}
		log.Debug(query)
		rows, errNw := DB.Query(query)
		if errNw != nil {
			log.Error(errNw)
			err = errNw
			return
		}
		defer rows.Close()
		for rows.Next() {
			model := &pb.OrganizedBookAudit{}
			models = append(models, model)
			err = rows.Scan(&model.BookId, &model.ParticipateNum)
			if err != nil {
				log.Error(err)
				return
			}
		}

	} else {
		// insert
		query = fmt.Sprintf("select isbn,count(*) from book_audit_record where status=1 and book_cate<>'' group by isbn order by count(*) desc,isbn offset %d limit %d", (record.Page-1)*record.Size, record.Size)
		queryCount = "select count(distinct isbn) from book_audit_record where status=1 and book_cate<>''"
		log.Debug(queryCount)
		err = DB.QueryRow(queryCount).Scan(&totalCount)
		if err != nil {
			log.Error(err)
			return
		}
		if totalCount <= 0 {
			return
		}

		log.Debug(query)
		rows, errNw := DB.Query(query)
		if errNw != nil {
			log.Error(errNw)
			err = errNw
			return
		}
		defer rows.Close()
		for rows.Next() {
			model := &pb.OrganizedBookAudit{}
			models = append(models, model)
			err = rows.Scan(&model.Isbn, &model.ParticipateNum)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}

	//organize the result and return
	for i := 0; i < len(models); i++ {
		model := models[i]
		book := &pb.Book{Isbn: model.Isbn, Id: model.BookId}
		bookList, err := GetBookInfo(book)
		if err != nil {
			log.Error(err)
			return models, err, totalCount
		}
		if len(bookList) > 0 {
			book = bookList[1]
			model.BookId = book.Id
			model.Isbn = book.Isbn
			model.BookCate = book.BookCate
			model.BookNo = book.BookNo
			model.Image = book.Image
			model.Title = book.Title
			model.Price = book.Price
			model.Publisher = book.Publisher
			model.Author = book.Author
			model.Edition = book.Edition
		}
	}

	return
}
