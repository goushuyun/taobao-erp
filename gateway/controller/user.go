package controller

import (
	"goushuyun/errs"
	"goushuyun/misc/token"
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

func ChangePwd(w http.ResponseWriter, r *http.Request) {
	req := &pb.User{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.Id = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "users", "ChangePwd", req, "mobile", "checkcode", "password")
}

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
