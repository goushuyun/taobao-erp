package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetUserInfoReq{}

	misc.CallWithResp(w, r, "bc_user", "GetUserInfo", req, "store_id", "appid")
}
