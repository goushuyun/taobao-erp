package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

func Login(w http.ResponseWriter, r *http.Request) {
	req := &pb.User{}

	misc.CallWithResp(w, r, "users", "Login", req, "mobile", "password")
}

func Register(w http.ResponseWriter, r *http.Request) {
	req := &pb.User{}

	misc.CallWithResp(w, r, "users", "Register", req, "mobile", "password", "name", "checkcode")
}

func CheckUserExist(w http.ResponseWriter, r *http.Request) {
	req := &pb.User{}

	misc.CallWithResp(w, r, "users", "UserExist", req, "mobile")
}
