package db

import (
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

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
