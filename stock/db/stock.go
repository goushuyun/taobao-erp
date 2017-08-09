package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func LocationFazzyQuery(l *pb.Location) ([]*pb.Location, int64, error) {
	query := "select %s from location where user_id = $1 %s"
	target := "id, warehouse, shelf, floor"
	var condition string

	if l.Warehouse != "" {
		condition += fmt.Sprintf("and warehouse like '%s'", misc.FazzyQuery(l.Warehouse))
	}
	if l.Shelf != "" {
		condition += fmt.Sprintf("and shelf like '%s'", misc.FazzyQuery(l.Shelf))
	}
	if l.Floor != "" {
		condition += fmt.Sprintf("and floor like '%s'", misc.FazzyQuery(l.Floor))
	}
	var total int64
	data := []*pb.Location{}
	err := DB.QueryRow(fmt.Sprintf(query, "count(1)", condition), l.UserId).Scan(&total)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	if total == 0 {
		return data, 0, nil
	}

	log.Debug(fmt.Sprintf(query, target, condition))
	rows, err := DB.Query(fmt.Sprintf(query, target, condition), l.UserId)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	for rows.Next() {
		tmp := &pb.Location{}
		err = rows.Scan(&tmp.LocationId, &tmp.Warehouse, &tmp.Shelf, &tmp.Floor)
		if err != nil {
			log.Error(err)
			return nil, 0, err
		}

		log.JSON(tmp)

		data = append(data, tmp)
	}

	return data, total, nil
}

func ListGoodsAllLocations(g *pb.Goods) ([]*pb.Goods, int64, error) {
	query := `
	select %s from goods_location_map as m
		left join goods as g on m.goods_id = g.id
		left join location as l on m.location_id = l.id
		where m.stock > 0 and m.goods_id = $1 and m.user_id = $2
	`
	var (
		total int64
		data  []*pb.Goods
	)

	log.Debug(fmt.Sprintf(query, "count(*)"))
	err := DB.QueryRow(fmt.Sprintf(query, "count(*)"), g.GoodsId, g.UserId).Scan(&total)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	if total == 0 {
		// not found
		return data, total, nil
	}

	target := "m.stock, g.status, g.remark, l.warehouse, l.shelf, l.floor, m.id"
	// join "order by" condition
	condition := " order by %s limit %d offset %d"
	var order_condition string

	switch g.OrderBy {
	case pb.ListOrderBy_LocationReverse:
		order_condition = "m.stock desc"
	case pb.ListOrderBy_StockForward:
		order_condition = "m.stock asc"
	case pb.ListOrderBy_UpdateAtReverse:
		order_condition = "m.update_at desc"
	case pb.ListOrderBy_UpdateAtForward:
		order_condition = "m.update_at asc"
	default:
		order_condition = "l.floor asc, l.shelf asc, l.warehouse asc"
	}

	condition = fmt.Sprintf(condition, order_condition, g.Size, (g.Page-1)*g.Size)

	log.Debug(fmt.Sprintf(query, target) + condition)
	rows, err := DB.Query(fmt.Sprintf(query, target)+condition, g.GoodsId, g.UserId)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	for rows.Next() {
		tmp := &pb.Goods{}
		err = rows.Scan(&tmp.Stock, &tmp.Status, &tmp.Remark, &tmp.Warehouse, &tmp.Shelf, &tmp.Floor, &tmp.MapRowId)
		if err != nil {
			log.Error(err)
			return nil, 0, err
		}
		data = append(data, tmp)
	}

	return data, total, nil
}

func UpdateStock(g *pb.Goods) error {
	query := "update goods_location_map set stock = stock + $1 where id = $2"
	_, err := DB.Exec(query, g.Stock, g.MapId)
	return err
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
