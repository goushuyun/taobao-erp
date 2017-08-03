package service

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/book/db"
	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

// submit a audit about the book
func (s *BookServer) SubmitBookAudit(ctx context.Context, in *pb.BookAuditRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SubmitBookAudit", "%#v", in))
	if in.BookId != "" {
		_, err, totalCount := db.GetBookAuditList(in)
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
		if totalCount > 0 {
			return nil, errs.Wrap(errors.New("你有该书未审核的图书信息，请勿重复提交"))
		}
	}

	err := db.SaveBookAudit(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// get book audit record
func (s *BookServer) GetBookAuditRecord(ctx context.Context, in *pb.BookAuditRecord) (*pb.BookAuditRecordListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetBookAuditRecord", "%#v", in))
	models, err, totalCount := db.GetBookAuditList(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.BookAuditRecordListResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}

// handle the audit request : accept or reject (when reject ,a reason better for user)
func (s *BookServer) UpdateAuditRecord(ctx context.Context, in *pb.BookAuditRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "UpdateAuditRecord", "%#v", in))
	err := db.UpdateBookAudit(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// get book audit organize list
func (s *BookServer) GetOrganizedBookAuditList(ctx context.Context, in *pb.BookAuditRecord) (*pb.OrganizedBookAuditListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetOrganizedBookAuditList", "%#v", in))
	models, err, totalCount := db.GetOrganizedBookAuditList(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.OrganizedBookAuditListResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}

//  handle the book audit
func (s *BookServer) HandleBookAudit(ctx context.Context, in *pb.BookAuditRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "HandleBookAudit", "%#v", in))
	err := db.BatchUpdateBookAudit(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}
