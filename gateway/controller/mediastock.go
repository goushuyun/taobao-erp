package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

func GetUpToken(w http.ResponseWriter, r *http.Request) {
	req := &pb.UpLoadReq{}
	misc.CallWithResp(w, r, "mediastore", "GetUpToken", req, "key")
}
