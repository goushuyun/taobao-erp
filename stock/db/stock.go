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

	if g.LocationId != "" {
		query += fmt.Sprintf(" and m.location_id='%s'", g.LocationId)
	}
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

	target := "m.stock, g.status, g.remark, l.warehouse, l.shelf, l.floor, m.id, l.id"
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
		order_condition = "l.warehouse, l.shelf, l.floor"
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
		err = rows.Scan(&tmp.Stock, &tmp.Status, &tmp.Remark, &tmp.Warehouse, &tmp.Shelf, &tmp.Floor, &tmp.MapRowId, &tmp.LocationId)
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

// get the location stock
func GetLocationStock(location *pb.Location) (locations []*pb.Location, err error, totalCount, totalStock int64) {

	queryCount := "select count(distinct l.id),sum(stock) from location l join goods_location_map m on l.id::uuid=m.location_id::uuid where 1=1 %s  having sum(stock)>0"
	query := "select l.id,warehouse,shelf,floor,sum(stock) from location l join goods_location_map m on l.id::uuid=m.location_id::uuid where 1=1 %s group by l.id,warehouse,shelf,floor having sum(stock)>0 order by warehouse,shelf,floor "
	var condition string
	if location.Warehouse != "" {
		condition += fmt.Sprintf(" and warehouse='%s'", location.Warehouse)
	}
	if location.Shelf != "" {
		condition += fmt.Sprintf(" and shelf='%s'", location.Shelf)
	}
	if location.Floor != "" {
		condition += fmt.Sprintf(" and floor='%s'", location.Floor)
	}
	if location.UserId != "" {
		condition += fmt.Sprintf(" and l.user_id='%s'", location.UserId)
	}
	queryCount = fmt.Sprintf(queryCount, condition)
	log.Debug(queryCount)
	err = DB.QueryRow(queryCount).Scan(&totalCount, &totalStock)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		log.Error(err)
		return
	}
	if totalCount <= 0 {
		return
	}
	query = fmt.Sprintf(query, condition)
	if location.Page <= 0 {
		location.Page = 1
	}
	if location.Size <= 0 {
		location.Size = 15
	}
	query += fmt.Sprintf(" offset %d limit %d", (location.Page-1)*location.Size, location.Size)
	log.Debug(query)
	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	for rows.Next() {
		model := &pb.Location{}
		locations = append(locations, model)
		err = rows.Scan(&model.LocationId, &model.Warehouse, &model.Shelf, &model.Floor, &model.TotalStock)
		if err != nil {
			log.Error(err)
			return
		}
	}
	return
}

// update location
func UpdateLocation(location *pb.Location) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Error(err)
		return err
	}
	defer tx.Rollback()

	query := "update location set update_at=now(),warehouse='%s',shelf='%s',floor='%s' where not exists(select * from location where warehouse='%s' and shelf='%s' and floor='%s' and user_id='%s') and id='%s'"
	query = fmt.Sprintf(query, location.Warehouse, location.Shelf, location.Floor, location.Warehouse, location.Shelf, location.Floor, location.UserId, location.LocationId)
	log.Debug(query)
	result, err := tx.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	counts, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if counts != 1 {
		return errors.New("exists")
	}
	tx.Commit()
	return nil
}

// add a record about the goods shift record
func AddGoodsShiftRecord(model *pb.GoodsShiftRecord) error {
	// first get the location detail by location id
	query := fmt.Sprintf("select warehouse,shelf,floor from location where id='%s'", model.LocationId)
	log.Debug(query)
	err := DB.QueryRow(query).Scan(&model.Warehouse, &model.Shelf, &model.Floor)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("参数错误")
		}
		log.Error(err)
		return err
	}
	// the save the record
	query = "insert into goods_shift_record(goods_id,location_id,warehouse,shelf,floor,user_id,stock,operate_type) values('%s','%s','%s','%s','%s','%s','%d','%s')"
	query = fmt.Sprintf(query, model.GoodsId, model.LocationId, model.Warehouse, model.Shelf, model.Floor, model.UserId, model.Stock, model.OperateType)
	log.Debug(query)
	_, err = DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// get goods shift record
func GetGoodsShiftRecord(model *pb.GoodsShiftRecord) (models []*pb.GoodsShiftRecord, err error, totalCount int64) {
	query := "select count(*) from goods_shift_record gs left join goods g on gs.goods_id::uuid=g.id left join book b on g.book_id=b.id left join users u on u.id=gs.user_id where 1=1"
	var condition string
	if model.StartAt != 0 && model.EndAt != 0 {
		condition += fmt.Sprintf(" and gs.create_at between to_timestamp(%d) and to_timestamp(%d)", model.StartAt, model.EndAt)
	}
	if model.Isbn != "" {
		condition += fmt.Sprintf(" and b.isbn ='%s'", model.Isbn)
	}
	if model.OperateType != "" {
		condition += fmt.Sprintf(" and operate_type='%s'", model.OperateType)
	}
	if model.UserId != "" {
		condition += fmt.Sprintf(" and gs.user_id='%s'", model.UserId)

	}
	query += condition
	log.Debug(query)
	err = DB.QueryRow(query).Scan(&totalCount)
	if err != nil {
		log.Error(err)
		return
	}
	if totalCount <= 0 {
		return
	}
	query = "select %s from  goods_shift_record gs left join goods g on gs.goods_id::uuid=g.id left join book b on g.book_id=b.id left join users u on u.id=gs.user_id where 1=1"
	param := " gs.id, gs.goods_id,gs.location_id,gs.warehouse,gs.shelf,gs.floor,gs.user_id,gs.stock,extract(epoch from gs.create_at)::bigint,gs.operate_type,b.isbn,b.book_no,b.book_cate,b.title,u.name"
	query = fmt.Sprintf(query, param)
	if model.Page <= 0 {
		model.Page = 1
	}
	if model.Size <= 0 {
		model.Size = 15
	}
	condition += fmt.Sprintf(" order by gs.create_at desc,gs.id offset %d limit %d", (model.Page-1)*model.Size, model.Size)
	query += condition
	log.Debug(query)
	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		result := &pb.GoodsShiftRecord{}
		models = append(models, result)
		//b.isbn,b.book_no,b.book_cate,b.title,u.name
		var isbn, book_no, book_cate, title, name sql.NullString
		err = rows.Scan(&result.Id, &result.GoodsId, &result.LocationId, &result.Warehouse, &result.Shelf, &result.Floor, &result.UserId, &result.Stock, &result.CreateAt, &result.OperateType, &isbn, &book_no, &book_cate, &title, &name)
		if err != nil {
			log.Error(err)
			return
		}
		if isbn.Valid {
			result.Isbn = isbn.String
		}
		if book_no.Valid {
			result.BookNo = book_no.String
		}
		if book_cate.Valid {
			result.BookCate = book_cate.String
		}
		if title.Valid {
			result.BookTitle = title.String
		}
		if name.Valid {
			result.UserName = name.String
		}
	}
	return
}
