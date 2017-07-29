package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//SellerLogin 手机号登录
func SellerLogin(w http.ResponseWriter, r *http.Request) {
	req := &pb.LoginModel{}
	misc.CallWithResp(w, r, "bc_seller", "SellerLogin", req, "mobile", "password")
}

//SellerRegister 商家注册
func SellerRegister(w http.ResponseWriter, r *http.Request) {
	req := &pb.RegisterModel{}
	misc.CallWithResp(w, r, "bc_seller", "SellerRegister", req, "mobile", "password", "message_code", "username")
}

//UpdatePasswordAndLogin 修改登录密码
func UpdatePasswordAndLogin(w http.ResponseWriter, r *http.Request) {
	req := &pb.RegisterModel{}
	misc.CallWithResp(w, r, "bc_seller", "UpdatePasswordAndLogin", req, "mobile", "password", "message_code")
}

//CheckMobileExist 检验手机号是否注册过
func CheckMobileExist(w http.ResponseWriter, r *http.Request) {
	req := &pb.CheckMobileReq{}
	misc.CallWithResp(w, r, "bc_seller", "CheckMobileExist", req, "mobile")
}

//GetTelCode 获取手机验证码
func GetTelCode(w http.ResponseWriter, r *http.Request) {
	req := &pb.CheckMobileReq{}
	misc.CallWithResp(w, r, "bc_seller", "GetTelCode", req, "mobile")
}

//GetUpdateTelCode 获取手机验证码
func GetUpdateTelCode(w http.ResponseWriter, r *http.Request) {
	req := &pb.CheckMobileReq{}
	misc.CallWithResp(w, r, "bc_seller", "GetUpdateTelCode", req, "mobile")
}

//SelfStores 获取商家所关联的店铺
func SelfStores(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	req := &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}
	misc.CallWithResp(w, r, "bc_seller", "SelfStores", req, "id")
}
