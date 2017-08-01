package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

// func GetGoodsLocationMap(g *pb.Goods) error {
// 	query := "select id sto"
// }

func UpdateStock(g *pb.Goods) error {
	query := "update goods_location_map set stock = stock + $1 where id = $2"
	_, err := DB.Exec(query, g.Stock, g.MapId)
	return err
}

func SaveMap(g_l_map *pb.Goods) error {
	query := "insert into goods_location_map(location_id, goods_id, stock) values($1, $2, $3) returning id, extract(epoch from create_at)::bigint"

	return DB.QueryRow(query, g_l_map.LocationId, g_l_map.GoodsId, g_l_map.Stock).Scan(&g_l_map.MapId, &g_l_map.CreateAt)
}

func SaveGoods(g *pb.Goods) error {
	query := "insert into goods(book_id, user_id, remark, status) values($1, $2, $3, $4) returning id"

	log.Debugf("insert into goods(book_id, user_id, remark, status) values('%s', '%s', '%s', %d) returning id", g.BookId, g.UserId, g.Remark, g.Status)
	return DB.QueryRow(query, g.BookId, g.UserId, g.Remark, g.Status).Scan(&g.GoodsId)
}

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

// create location
func CreateLocation(l *pb.Location) error {
	query := "insert into location(warehouse, shelf, floor, user_id) values($1, $2, $3, $4) returning id"
	log.Debugf("insert into location(warehouse, shelf, floor, user_id) values('%s', '%s', '%s', '%s') returning id", l.Warehouse, l.Shelf, l.Floor, l.UserId)

	return DB.QueryRow(query, l.Warehouse, l.Shelf, l.Floor, l.UserId).Scan(&l.LocationId)
}

// get location
func GetLocationId(l *pb.Location) error {
	query := "select id from location where warehouse = $1 and shelf = $2 and floor = $3 and user_id = $4"

	log.Debugf("select id from location where warehouse = '%s' and shelf = '%s' and floor = '%s' and user_id = '%s'", l.Warehouse, l.Shelf, l.Floor, l.UserId)

	err := DB.QueryRow(query, l.Warehouse, l.Shelf, l.Floor, l.UserId).Scan(&l.LocationId)

	if err == sql.ErrNoRows {
		log.Debug("Location not found")
		return errors.New("not_found")
	}

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func LocationExist(goods *pb.Goods) (bool, error) {
	query := "select count(1) from location where warehouse = $1 and shelf = $2 and floor = $3"
	var total int64

	err := DB.QueryRow(query, goods.Warehouse, goods.Shelf, goods.Floor).Scan(&total)
	log.Debugf("select count(1) from location where warehouse = '%s' and shelf = '%s' and floor = '%s'", goods.Warehouse, goods.Shelf, goods.Floor)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return total > 0, nil
}

func SearchLocation(l *pb.Location) ([]*pb.Location, error) {
	query := "select warehouse, shelf, floor, id from location where user_id = $1"

	if l.Warehouse != "" {
		query += fmt.Sprintf(" and warehouse = '%s'", l.Warehouse)
	}

	if l.Shelf != "" {
		query += fmt.Sprintf(" and shelf = '%s'", l.Shelf)
	}

	if l.Floor != "" {
		query += fmt.Sprintf(" and floor = '%s'", l.Floor)
	}

	rows, err := DB.Query(query, l.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var locations []*pb.Location
	query_total_stock := "select sum(stock) from goods_location_map where location_id = $1"
	if rows.Next() {
		location := &pb.Location{}
		err = rows.Scan(&location.Warehouse, &location.Shelf, &location.Floor, &location.LocationId)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		// compute total stock at location
		DB.QueryRow(query_total_stock, location.LocationId).Scan(&location.TotalStock)
		locations = append(locations, location)
	}

	return locations, nil
}
