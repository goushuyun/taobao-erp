package service

import (
	"github.com/goushuyun/taobao-erp/misc"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func (s *StockServer) LocationFazzyQuery(ctx context.Context, req *pb.Location) (*pb.LocationBatchResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ReduceMapRow", "%#v", req))

	return &pb.LocationBatchResp{}, nil
}
