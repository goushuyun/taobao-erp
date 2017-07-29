package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//管理员登陆
func MasterLogin(w http.ResponseWriter, r *http.Request) {
	req := &pb.Master{}
	misc.CallWithResp(w, r, "bc_master", "MasterLogin", req, "mobile", "password")
}

//提现列表
func WithdrawList(w http.ResponseWriter, r *http.Request) {
	req := &pb.StoreWithdrawalsModel{}
	misc.CallWithResp(w, r, "bc_master", "WithdrawList", req)
}

//开始处理提现
func WithdrawHandle(w http.ResponseWriter, r *http.Request) {
	req := &pb.StoreWithdrawalsModel{}
	misc.CallWithResp(w, r, "bc_master", "WithdrawHandle", req)
}

//提现完成
func WithdrawComplete(w http.ResponseWriter, r *http.Request) {
	req := &pb.StoreWithdrawalsModel{}
	misc.CallWithResp(w, r, "bc_master", "WithdrawComplete", req)
}
