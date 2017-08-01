package db

import (
	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func SaveGoods(g *pb.Goods) error {
	query := "insert into goods(book_id, user_id, remark, status) values($1, $2, $3, $4) returning id"

	log.Debugf("insert into goods(book_id, user_id, remark, status) values('%s', '%s', '%s', %d) returning id", g.BookId, g.UserId, g.Remark, g.Status)
	return DB.QueryRow(query, g.BookId, g.UserId, g.Remark, g.Status).Scan(&g.GoodsId)
}
