package service

import (
	"github.com/goushuyun/taobao-erp/pb"
	"golang.org/x/net/context"
)

type UsersServer struct {
}

func (s *UsersServer) Register(ctx context.Context, req *pb.User) (*pb.UserResp, error) {
	return nil, nil
}
