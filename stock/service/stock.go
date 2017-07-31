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

func (s *StockServer) SaveSingleGoods(ctx context.Context, req *pb.Goods) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveSingleGoods", "%#v", req))

	// to get book_id

	// save stock

	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}
