package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/goushuyun/log"
	"github.com/wothing/worpc"

	"github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/sms/service"
)

const (
	svcName = "sms"
	port    = 10015
)

func main() {
	db.InitRedis(svcName)
	defer db.CloseRedis()

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
