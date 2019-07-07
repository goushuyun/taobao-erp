package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/stock/db"

	"golang.org/x/net/context"

	"github.com/goushuyun/log"
	"github.com/goushuyun/taobao-erp/pb"
)

func (s *StockServer) DelLocation(ctx context.Context, req *pb.Location) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ReduceMapRow", "%#v", req))

	if err := db.DelLocation(req.LocationId); err != nil {
		log.Error(err)
		return nil, errs.Wrap(err)
	}

	return &pb.NormalResp{Code: errs.Ok, Message: "The location was deleted successfully"}, nil
}

func (s *StockServer) LocationFazzyQuery(ctx context.Context, req *pb.Location) (*pb.LocationBatchResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ReduceMapRow", "%#v", req))

	data, total, err := db.LocationFazzyQuery(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.LocationBatchResp{Code: errs.Ok, Message: "ok", Data: data, Total: total}, nil
}
