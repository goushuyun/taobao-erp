package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func GetGoodsByBookId(g *pb.Goods) error {
	query := "select id, status, remark from goods where user_id = $1 and book_id = $2"
	log.Debugf("select id, status, remark from goods where user_id = '%s' and book_id = '%s'", g.UserId, g.BookId)
	err := DB.QueryRow(query, g.UserId, g.BookId).Scan(&g.GoodsId, &g.Status, &g.Remark)

	if err == sql.ErrNoRows {
		log.Debug("Goods not fount")
		return errors.New("not_found")
	}

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func UpdateGoods(g *pb.Goods) error {
	query := "update goods %s where id = $1"

	var operation string
	if g.Stock != 0 {
		operation += fmt.Sprintf("set stock = stock + %d", g.Stock)
	}
	if g.Status != 0 {
		operation += fmt.Sprintf("set status = %d", g.Status)
	}
	if g.Remark != "" {
		operation += fmt.Sprintf("set remark = %s", g.Remark)
	}

	log.Debugf(query, operation)
	_, err := DB.Exec(fmt.Sprintf(query, operation), g.GoodsId)

	return err
}

func SaveGoods(g *pb.Goods) error {
	query := "insert into goods(book_id, user_id, remark, status) values($1, $2, $3, $4) returning id"

	log.Debugf("insert into goods(book_id, user_id, remark, status) values('%s', '%s', '%s', %d) returning id", g.BookId, g.UserId, g.Remark, g.Status)
	return DB.QueryRow(query, g.BookId, g.UserId, g.Remark, g.Status).Scan(&g.GoodsId)
}

//  search goods from local db

func SearchGoods(goods *pb.GoodsInfo) (models []*pb.GoodsInfo, err error, totalCount int64) {
	query := "select g.id,g.remark,g.stock,b.id,b.isbn,b.book_no,b.book_cate,b.title,b.publisher,b.author,b.edition,b.image,b.price from goods g  join book b on g.book_id=b.id where 1=1"
	queryCount := "select count(*) from goods g  join book b on g.book_id=b.id where 1=1"
	var condition string
	if goods.Isbn != "" {
		condition += fmt.Sprintf(" and b.isbn='%s'", goods.Isbn)
	}
	if goods.Title != "" {
		condition += fmt.Sprintf(" and b.title='%s'", goods.Title)
	}
	if goods.Publisher != "" {
		condition += fmt.Sprintf(" and b.publisher='%s'", goods.Publisher)
	}
	if goods.Author != "" {
		condition += fmt.Sprintf(" and b.author='%s'", goods.Author)
	}
	if goods.Compare != "" {
		if goods.Compare == "less" {
			condition += fmt.Sprintf(" and g.stock<%d", goods.Stock)
		} else if goods.Compare == "greater" {
			condition += fmt.Sprintf(" and g.stock>=%d", goods.Stock)
		}
	}
	if goods.UserId != "" {
		condition += fmt.Sprintf(" and g.user_id='%s'", goods.UserId)
	}
	if goods.GoodsId != "" {
		condition += fmt.Sprintf(" and g.id='%s'", goods.GoodsId)
	}

	queryCount += condition
	log.Debug(queryCount)
	err = DB.QueryRow(queryCount).Scan(&totalCount)
	if err != nil {
		log.Error(err)
		return
	}
	if totalCount <= 0 {
		return
	}
	if goods.Sequence != "asc" && goods.Sequence != "desc" {
		goods.Sequence = ""
	}
	if goods.OrderBy != "" {
		condition += fmt.Sprintf(" order by %s %s,g.update_at desc,g.id desc", goods.OrderBy, goods.Sequence)
	} else {
		condition += " order by g.update_at desc,g.id desc "
	}
	if goods.Page <= 0 {
		goods.Page = 1
	}
	if goods.Size <= 0 {
		goods.Size = 15
	}
	condition += fmt.Sprintf(" offset %d limit %d", (goods.Page-1)*goods.Size, goods.Size)
	query += condition
	log.Debug(query)
	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		model := &pb.GoodsInfo{}
		models = append(models, model)
		//g.id,g.remark,g.stock,b.id,b.isbn,b.book_no,b.book_cate,b.title,b.publisher,b.author,b.edition,b.image,b.price
		err = rows.Scan(&model.GoodsId, &model.Remark, &model.Stock, &model.BookId, &model.Isbn, &model.BookNo, &model.BookCate, &model.Title, &model.Publisher, &model.Author, &model.Edition, &model.Image, &model.Price)
		if err != nil {
			log.Error(err)
			return
		}
	}
	return
}

// update the goods info
func UpdateGoodsInfo(goods *pb.Goods) error {
	query := "update goods set update_at=now()"
	if goods.Remark != "" {
		query += fmt.Sprintf(",remark='%s'", goods.Remark)
	}
	query += fmt.Sprintf(" where id='%s' and user_id='%s'", goods.GoodsId, goods.UserId)
	log.Debug(query)
	_, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
