package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/db"
	"github.com/wothing/log"
)

type StockServer struct{}

func (s *StockServer) UpdateMapRow(ctx context.Context, req *pb.MapRowBatch) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ReduceMapRow", "%#v", req))

	for _, map_row := range req.Data {
		g := &pb.Goods{Stock: map_row.Stock}

		err := db.UpdateMapRow(map_row)
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}

		// update goods's stock
		g.GoodsId = map_row.GoodsId
		err = db.UpdateGoods(g)
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}

		// add input/output record
		var operate_type string
		if g.Stock > 0 {
			operate_type = "load"
		} else {
			operate_type = "unload"
		}

		goods_shift_record := &pb.GoodsShiftRecord{
			GoodsId:     map_row.GoodsId,
			LocationId:  map_row.LocationId,
			Stock:       g.Stock,
			UserId:      req.UserId,
			OperateType: operate_type,
		}

		log.JSON(goods_shift_record)

		err = db.AddGoodsShiftRecord(goods_shift_record)
		if err != nil {
			log.Error(err)
		}
	}

	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

func (s *StockServer) ListGoodsAllLocations(ctx context.Context, req *pb.Goods) (*pb.GoodsResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ListGoodsLocations", "%#v", req))

	data, total, err := db.ListGoodsAllLocations(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.GoodsResp{Code: errs.Ok, Message: "ok", Data: data, Total: total}, nil
}

func (s *StockServer) GetLocationId(ctx context.Context, req *pb.Location) (*pb.LocationResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocationId", "%#v", req))

	err := db.GetLocationId(req)

	if err != nil {
		if err.Error() == notFound {
			// create location
			err = db.CreateLocation(req)
			if err != nil {
				log.Error(err)
				return nil, errs.Wrap(errors.New(err.Error()))
			}
		} else {
			// met other error
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
	}

	return &pb.LocationResp{Code: errs.Ok, Message: "ok", Data: req}, nil
}

func (s *StockServer) SearchLocation(ctx context.Context, req *pb.Location) (*pb.SearchLocationResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SearchGoods", "%#v", req))

	data, err := db.SearchLocation(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.SearchLocationResp{Code: errs.Ok, Message: "ok", Data: data}, nil
}

func (s *StockServer) SaveMapRow(ctx context.Context, req *pb.MapRow) (*pb.MapRowResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveSingleGoods", "%#v", req))

	// update goods's stock
	var operate_stock int64 = req.Stock
	defer func() {
		// modify Goods record
		g := &pb.Goods{Stock: operate_stock, GoodsId: req.GoodsId}
		err := db.UpdateGoods(g)
		if err != nil {
			log.Error(err)
		}

		// add input/output record
		var operate_type string
		if operate_stock > 0 {
			operate_type = "load"
		} else {
			operate_type = "unload"
		}
		goods_shift_record := &pb.GoodsShiftRecord{
			GoodsId:     req.GoodsId,
			LocationId:  req.LocationId,
			Stock:       operate_stock,
			UserId:      req.UserId,
			OperateType: operate_type,
		}
		err = db.AddGoodsShiftRecord(goods_shift_record)
		if err != nil {
			log.Error(err)
		}
	}()

	err := db.GetMapRow(req)
	if err != nil {
		if err.Error() == notFound {
			// create this map row
			err = db.SaveMapRow(req)
			if err != nil {
				log.Error(err)
				return nil, errs.Wrap(errors.New(err.Error()))
			}
			return &pb.MapRowResp{Code: errs.Ok, Message: "ok", Data: req}, nil
		} else {
			// met other error
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
	}

	// update map row
	err = db.UpdateMapRow(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.MapRowResp{Code: errs.Ok, Message: "ok", Data: req}, nil
}

// get the location stock
func (s *StockServer) GetLocationStock(ctx context.Context, req *pb.Location) (*pb.LocationBatchResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocationStock", "%#v", req))
	locations, err, totalCount, totalStock := db.GetLocationStock(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.LocationBatchResp{Code: errs.Ok, Message: "ok", Data: locations, TotalCount: totalCount, Total: totalStock}, nil
}

// get the location stock
func (s *StockServer) UpdateLocation(ctx context.Context, req *pb.Location) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocationStock", "%#v", req))
	err := db.UpdateLocation(req)
	if err != nil {
		if err.Error() == "exists" {
			return &pb.NormalResp{Code: errs.Ok, Message: "exists"}, nil
		}
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// get goods'g shift record
func (s *StockServer) GetGoodsShiftRecord(ctx context.Context, req *pb.GoodsShiftRecord) (*pb.GoodsShiftRecordListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetGoodsShiftRecord", "%#v", req))
	models, err, totalCount := db.GetGoodsShiftRecord(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.GoodsShiftRecordListResp{Code: errs.Ok, Message: "ok", TotalCount: totalCount, Data: models}, nil
}

//获取入库出库记录的导出时间
func (s *StockServer) GetShiftRocordExportDate(ctx context.Context, req *pb.User) (*pb.UserResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetShiftRocordExportDate", "%#v", req))
	err := db.GetShiftRocordExportDate(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.UserResp{Code: errs.Ok, Message: "ok", Data: req}, nil
}

//获取入库出库记录的导出时间
func (s *StockServer) UpdateShiftRocordExportDate(ctx context.Context, req *pb.User) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetShiftRocordExportDate", "%#v", req))
	err := db.UpdateShiftRocordExportDate(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}
