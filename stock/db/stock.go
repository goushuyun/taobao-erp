package db

import (
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

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
