package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/wothing/log"
	"github.com/wothing/worpc"

	"github.com/goushuyun/weixin-golang/db"
	"github.com/goushuyun/weixin-golang/pb"
	"github.com/goushuyun/weixin-golang/sms/service"
)

const svcName = "bc_sms"
const port = 8850

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", db.GetPort(port)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infof("starting to listen at :%d", db.GetPort(port))

	err = db.RegisterService(svcName, port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(worpc.UnaryInterceptorChain(worpc.Recovery, worpc.Logging)))
	pb.RegisterSMSServiceServer(s, &service.SMSServer{})
	s.Serve(lis)
}
