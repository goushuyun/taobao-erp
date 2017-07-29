package service

import (
	"errors"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/wothing/log"
	"golang.org/x/net/context"
)

type MediastoreServer struct {
	Test bool
}

// extract image from weixin server and upload to qiniu
func (s *MediastoreServer) ExtractImageFromWeixinToQiniu(ctx context.Context, req *pb.ExtractReq) (*pb.ExtractResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "ExtractImageFromWeixinToQiniu", "%#v", req))

	qiniu_keys := []string{}

	for _, url := range req.WeixinMediaUrls {

		key, err := upload(url, req.Zone, req.Appid)
		if err != nil {
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}

		qiniu_keys = append(qiniu_keys, key)
	}

	return &pb.ExtractResp{QiniuKeys: qiniu_keys}, nil
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
