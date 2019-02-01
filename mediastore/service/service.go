package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"

	"github.com/goushuyun/taobao-erp/pb"

	"github.com/goushuyun/log"
	"golang.org/x/net/context"
)

type MediastoreServer struct {
	Test bool
}

// refresh urls
func (s *MediastoreServer) RefreshUrls(ctx context.Context, req *pb.RefreshReq) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "RefreshUrls", "%#v", req))

	// to refresh urls
	err := RefreshURLCache(req.Urls)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

func (s *MediastoreServer) GetUpToken(ctx context.Context, req *pb.UpLoadReq) (*pb.GetUpTokenResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetUpToken", "%#v", req))

	token, url := makeToken(req.Zone, req.Key)
	return &pb.GetUpTokenResp{Token: token, Url: url}, nil
}

func (s *MediastoreServer) FetchImage(ctx context.Context, req *pb.FetchImageReq) (*pb.FetchImageResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "FetchImage", "%#v", req))

	url, err := FetchImg(req.Zone, req.Url, req.Key)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.FetchImageResp{QiniuUrl: url}, nil
}
