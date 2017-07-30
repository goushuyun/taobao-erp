package service

import (
	"fmt"

	"github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/pb"
	users_db "github.com/goushuyun/taobao-erp/users/db"
	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/misc/token"
	"github.com/goushuyun/taobao-erp/users/role"
	"github.com/wothing/log"
)

type UsersServer struct {
}

func (s *UsersServer) Register(ctx context.Context, req *pb.User) (*pb.UserResp, error) {
	// check if identifying code is ok
	conn := db.GetRedisConn()
	defer conn.Close()

	checkcode, err := conn.Do("get", req.Mobile)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// checkcode is expired
	if checkcode == nil {
		return &pb.UserResp{Code: errs.Ok, Message: "checkcode_expired"}, nil
	}

	// checkcode is error
	checkcode = fmt.Sprintf("%s", checkcode)
	if checkcode != req.Checkcode {
		return &pb.UserResp{Code: errs.Ok, Message: "checkcode_error"}, nil
	}

	// insert user to db
	err = users_db.SaveUser(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// sign token
	token := token.SignUserToken(role.InterNormalUser, req.Id)
	req.Password = "*****"

	return &pb.UserResp{Code: errs.Ok, Message: "ok", Data: req, Token: token}, nil
}
