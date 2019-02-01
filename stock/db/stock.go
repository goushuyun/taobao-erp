package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/log"
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
	query := fmt.Sprintf("select warehouse,shelf,floor,user_id from location where id='%s'", model.LocationId)
	log.Debug(query)
	var userId string
	err := DB.QueryRow(query).Scan(&model.Warehouse, &model.Shelf, &model.Floor, &userId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("参数错误")
		}
		log.Error(err)
		return err
	}

	model.UserId = userId

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

	if model.SizeLimit == "none" {
		condition += fmt.Sprintf(" order by gs.warehouse,gs.shelf,gs.floor,b.isbn,gs.id")
	} else {
		if model.Page <= 0 {
			model.Page = 1
		}
		if model.Size <= 0 {
			model.Size = 15
		}
		condition += fmt.Sprintf(" order by gs.create_at desc,gs.id offset %d limit %d", (model.Page-1)*model.Size, model.Size)
	}
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

//更改上架记录导日期
func UpdateShiftRocordExportDate(user *pb.User) error {
	query := "update users set export_start_at=%d,export_end_at=%d where id='%s'"
	query = fmt.Sprintf(query, user.ExportStartAt, user.ExportEndAt, user.Id)
	log.Debug(query)
	_, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//获取导出时间
func GetShiftRocordExportDate(user *pb.User) error {
	query := fmt.Sprintf("select export_start_at,export_end_at from users where id='%s'", user.Id)
	log.Debug(query)
	err := DB.QueryRow(query).Scan(&user.ExportStartAt, &user.ExportEndAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Error(err)
		return err
	}
	return nil
}

// get user taobao setting
func GetUserTaobaoSetting(model *pb.TaobaoCsvRecord) error {
	query := "select id ,user_id,discount,supplemental_fee,province,city,express_template,pingyou_fee,express_fee,ems_fee,reduce_stock_style,extract(epoch from create_at)::bigint,product_title,product_describe from users_taobao_setting where user_id='%s'"
	query = fmt.Sprintf(query, model.UserId)
	log.Debug(query)
	err := DB.QueryRow(query).Scan(&model.Id, &model.UserId, &model.Discount, &model.SupplementalFee, &model.Province, &model.City, &model.ExpressTemplate, &model.PingyouFee, &model.ExpressFee, &model.EmsFee, &model.ReduceStockStyle, &model.CreateAt, &model.ProductTitle, &model.ProductDescribe)
	if err != nil {
		if err == sql.ErrNoRows {
			query = fmt.Sprintf("insert into users_taobao_setting(user_id) values('%s') returning id", model.UserId)
			err = DB.QueryRow(query).Scan(&model.Id)
			if err != nil {
				log.Error(err)
				return err
			}
		} else {
			log.Error(err)
			return err
		}
	}
	return nil
}

// update user's taobao setting
func UpdateUserTaobaoSetting(model *pb.TaobaoCsvRecord) error {
	query := "update users_taobao_setting set update_at=now()"
	var conditon string
	conditon += fmt.Sprintf(",discount=%d", model.Discount)

	conditon += fmt.Sprintf(",supplemental_fee=%d", model.SupplementalFee)

	if model.Province != "" {
		conditon += fmt.Sprintf(",province='%s'", model.Province)
	}
	if model.City != "" {
		conditon += fmt.Sprintf(",city='%s'", model.City)
	}
	if model.ExpressTemplate != "" {
		conditon += fmt.Sprintf(",express_template='%s'", model.ExpressTemplate)
	}
	if model.PingyouFee != 0 {
		conditon += fmt.Sprintf(",pingyou_fee=%d", model.PingyouFee)
	}
	if model.ExpressFee != 0 {
		conditon += fmt.Sprintf(",express_fee=%d", model.ExpressFee)
	}
	if model.EmsFee != 0 {
		conditon += fmt.Sprintf(",ems_fee=%d", model.EmsFee)
	}
	if model.ReduceStockStyle != "" {
		conditon += fmt.Sprintf(",reduce_stock_style='%s'", model.ReduceStockStyle)
	}
	if model.ProductTitle != "" {
		conditon += fmt.Sprintf(",product_title='%s'", model.ProductTitle)
	}
	if model.ProductDescribe != "" {
		conditon += fmt.Sprintf(",product_describe='%s'", model.ProductDescribe)
	}
	conditon += fmt.Sprintf(" where id='%s'", model.Id)
	query += conditon
	log.Debug(query)
	_, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// add a record about taobaocsv
func AddExportCsvRecord(tx *sql.Tx, model *pb.TaobaoCsvRecord) error {
	query := "insert into taobao_csv_record(user_id,discount,supplemental_fee,province,city,express_template,pingyou_fee,express_fee,ems_fee,reduce_stock_style,product_title,product_describe,search_isbn,search_title,search_publisher,search_compare,search_stock,search_author) values('%s',%d,%d,'%s','%s','%s',%d,%d,%d,'%s','%s','%s','%s','%s','%s','%s',%d,'%s') returning id"
	query = fmt.Sprintf(query, model.UserId, model.Discount, model.SupplementalFee, model.Province, model.City, model.ExpressTemplate, model.PingyouFee, model.ExpressFee, model.EmsFee, model.ReduceStockStyle, model.ProductTitle, model.ProductDescribe, model.Isbn, model.Title, model.Publisher, model.Compare, model.Stock, model.Author)
	log.Debug(query)
	err := tx.QueryRow(query).Scan(&model.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// add item to record
func AddExportCsvRecordRelatedItems(tx *sql.Tx, model *pb.TaobaoCsvRecord) error {
	query := "insert into taobao_csv_record_item(goods_id,taobao_csv_record_id) select g.id,'%s' from goods g join book b on g.book_id=b.id where 1=1"
	query = fmt.Sprintf(query, model.Id)
	var condition string
	if model.Isbn != "" {
		condition += fmt.Sprintf(" and b.isbn='%s'", model.Isbn)
	}
	if model.BookNo != "" {
		condition += fmt.Sprintf(" and b.book_no='%s'", model.BookNo)
	}

	if model.Title != "" {
		condition += fmt.Sprintf(" and b.title like '%s'", misc.FazzyQuery(model.Title))
	}
	if model.Publisher != "" {
		condition += fmt.Sprintf(" and b.publisher='%s'", model.Publisher)
	}
	if model.Author != "" {
		condition += fmt.Sprintf(" and b.author='%s'", model.Author)
	}
	if model.Compare != "" {
		if model.Compare == "less" {
			condition += fmt.Sprintf(" and g.stock<%d", model.Stock)
		} else if model.Compare == "greater" {
			condition += fmt.Sprintf(" and g.stock>=%d", model.Stock)
		}
	}
	if model.UserId != "" {
		condition += fmt.Sprintf(" and g.user_id='%s'", model.UserId)
	}

	query += condition

	log.Debug(query)
	_, err := tx.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//获取淘宝导出记录
func GetTaobaoCsvExportRecord(req *pb.TaobaoCsvRecord) (models []*pb.TaobaoCsvRecord, err error, totalCount int64) {
	query := "select id,user_id,discount,supplemental_fee,province,city,express_template,pingyou_fee,express_fee,ems_fee,reduce_stock_style,status,file_url,extract(epoch from create_at)::bigint,extract(epoch from create_at)::bigint,summary,product_title,product_describe,search_isbn,search_title,search_publisher,search_compare,search_stock,search_author from taobao_csv_record where 1=1"
	queryCount := "select count(*) from  taobao_csv_record where 1=1"
	if req.UserId != "" {
		query += fmt.Sprintf(" and user_id='%s'", req.UserId)
		queryCount += fmt.Sprintf(" and user_id='%s'", req.UserId)

	}
	if req.Id != "" {
		query += fmt.Sprintf(" and id='%s'", req.Id)
		queryCount += fmt.Sprintf(" and id='%s'", req.Id)

	}
	if req.Status != 0 {
		query += fmt.Sprintf(" and status=%d", req.Status)
		queryCount += fmt.Sprintf(" and status=%d", req.Status)
	}
	log.Debug(queryCount)
	err = DB.QueryRow(queryCount).Scan(&totalCount)
	if err != nil {
		log.Error(err)
		return
	}
	if totalCount <= 0 {
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 15
	}

	query += fmt.Sprintf(" order by create_at desc,id desc offset %d limit %d", req.Size*(req.Page-1), req.Size)

	log.Debug(query)

	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		model := &pb.TaobaoCsvRecord{}
		models = append(models, model)
		//id,user_id,discount,supplemental_fee,province,city,express_template,pingyou_fee,express_fee,ems_fee,reduce_stock_style,status,file_url,extract(epoch from complete_at)::bigint,extract(epoch from create_at)::bigint
		var completeAt sql.NullInt64
		err = rows.Scan(&model.Id, &model.UserId, &model.Discount, &model.SupplementalFee, &model.Province, &model.City, &model.ExpressTemplate, &model.PingyouFee, &model.ExpressFee, &model.EmsFee, &model.ReduceStockStyle, &model.Status, &model.FileUrl, &completeAt, &model.CreateAt, &model.Summary, &model.ProductTitle, &model.ProductDescribe, &model.Isbn, &model.Title, &model.Publisher, &model.Compare, &model.Stock, &model.Author)
		if err != nil {
			log.Error(err)
			return
		}
		if completeAt.Valid {
			model.CompleteAt = completeAt.Int64
		}
	}
	return

}

// get items about taobao exported record
func GetTaobaoCsvExportRecordItems(req *pb.TaobaoCsvRecord) (models []*pb.TaobaoCsvRecordItem, err error) {
	query := "select t.id,t.goods_id,g.stock,b.isbn,b.book_no,b.title,b.publisher,b.author,b.edition,b.pubdate,b.series_name,b.image,b.price,b.catalog,b.abstract,b.page,b.packing,b.format,b.author_intro,b.taobao_category from taobao_csv_record_item t join goods g on t.goods_id::uuid=g.id join book b on g.book_id=b.id where taobao_csv_record_id='%s' order by id desc"
	query = fmt.Sprintf(query, req.Id)
	log.Debug(query)
	rows, err := DB.Query(query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		model := &pb.TaobaoCsvRecordItem{}
		models = append(models, model)
		err = rows.Scan(&model.Id, &model.GoodsId, &model.Stock, &model.Isbn, &model.BookNo, &model.Title, &model.Publisher, &model.Author, &model.Edition, &model.Pubdate, &model.SeriesName, &model.Image, &model.Price, &model.Catalog, &model.Abstract, &model.Page, &model.Packing, &model.Format, &model.AuthorIntro, &model.TaobaoCategory)
		if err != nil {
			log.Error(err)
			return
		}
	}
	return
}

// get items about taobao exported record
func UpdatTaobaoCsvExportRecordItems(req *pb.TaobaoCsvRecord) error {
	query := "update taobao_csv_record set update_at=now()"
	if req.Summary != "" {
		query += fmt.Sprintf(",summary='%s'", req.Summary)
	}
	if req.Status != 0 {
		query += fmt.Sprintf(",status=%d", req.Status)
	}
	if req.FileUrl != "" {
		query += fmt.Sprintf(",file_url='%s'", req.FileUrl)
	}
	query += fmt.Sprintf(" where id='%s'", req.Id)
	log.Debug(query)
	_, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// del taobao csv record items
func DelTaobaoCsvExportRecordItems(req *pb.TaobaoCsvRecord) error {
	query := fmt.Sprintf("delete from taobao_csv_record_item where taobao_csv_record_id='%s'", req.Id)
	log.Debug(query)
	_, err := DB.Exec(query)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
