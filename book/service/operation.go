package service

import (
	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"

	"github.com/wothing/log"
)

// submit a audit about the book
func (s *BookServer) SubmitBookAudit(ctx context.Context, in *pb.BookAuditRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SubmitBookAudit", "%#v", in))

	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// get book audit record
func (s *BookServer) GetBookAuditRecord(ctx context.Context, in *pb.BookAuditRecord) (*pb.BookAuditRecordListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetBookAuditRecord", "%#v", in))

	return &pb.BookAuditRecordListResp{Code: errs.Ok, Message: "ok"}, nil
}

// handle the audit request : accept or reject (when reject ,a reason better for user)
func (s *BookServer) UpdateAuditRecord(ctx context.Context, in *pb.BookAuditRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "UpdateAuditRecord", "%#v", in))

	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}
