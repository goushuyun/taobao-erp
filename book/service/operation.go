package service

import (
	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"

	"github.com/wothing/log"
)

//submit a audit about the book
func (s *BookServer) SubmitBookAudit(ctx context.Context, in *pb.BookAuditRecord) (*pb.BookAuditRecordListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SubmitBookAudit", "%#v", in))
	return &pb.BookAuditRecordListResp{Code: errs.Ok, Message: "ok"}, nil
}
