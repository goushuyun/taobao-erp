package db

import (
	"database/sql"
	"errors"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func UpdateMapRow(m *pb.MapRow) error {
	query := "update goods_location_map set stock = stock + $1, update_at = now() where id = $2 returning stock, extract(epoch from update_at)::bigint"

	log.Debugf("update goods_location_map set stock = stock + %d, update = now() where id = '%s' returning stock, extract(epoch from update_at)::bigint", m.Stock, m.MapRowId)
	return DB.QueryRow(query, m.Stock, m.MapRowId).Scan(&m.Stock, &m.UpdateAt)
}

func SaveMapRow(m *pb.MapRow) error {
	query := "insert into goods_location_map(location_id, goods_id, stock, user_id) values($1, $2, $3, $4) returning id, extract(epoch from create_at)::bigint, extract(epoch from update_at)::bigint"

	log.Debugf("insert into goods_location_map(location_id, goods_id, stock, user_id) values('%s', '%s', %d, '%s') returning id, extract(epoch from create_at)::bigint, extract(epoch from update_at)::bigint", m.LocationId, m.GoodsId, m.Stock, m.UserId)

	return DB.QueryRow(query, m.LocationId, m.GoodsId, m.Stock, m.UserId).Scan(&m.MapRowId, &m.CreateAt, &m.UpdateAt)
}

func GetMapRow(m *pb.MapRow) error {
	query := "select id, extract(epoch from create_at)::bigint, extract(epoch from update_at)::bigint from goods_location_map where location_id = $1 and goods_id = $2"
	log.Debugf("select id, extract(epoch from create_at)::bigint, extract(epoch from update_at)::bigint from goods_location_map where location_id = '%s' and goods_id = '%s'", m.LocationId, m.GoodsId)

	err := DB.QueryRow(query, m.LocationId, m.GoodsId).Scan(&m.MapRowId, &m.CreateAt, &m.UpdateAt)
	if err == sql.ErrNoRows {
		return errors.New("not_found")
	} else if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
