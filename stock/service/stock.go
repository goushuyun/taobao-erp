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

func (s *StockServer) GetLocationId(ctx context.Context, req *pb.Location) (*pb.LocationResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocationId", "%#v", req))

	err := db.GetLocationId(req)

	if err != nil {
		if err.Error() == "not_found" {
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

func (s *StockServer) SaveSingleGoods(ctx context.Context, req *pb.MapRow) (*pb.MapRowResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveSingleGoods", "%#v", req))

	err := db.GetMapRow(req)
	if err != nil {
		if err.Error() == "not_found" {
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
