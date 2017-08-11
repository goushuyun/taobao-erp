package main

import (
	"github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/robfig/cron"
	"github.com/wothing/worpc"
	"google.golang.org/grpc"

	"github.com/goushuyun/taobao-erp/book/register"
	"github.com/goushuyun/taobao-erp/book/service"
)

const (
	svcName = "book"
	port    = 10018
)

var svcNames = []string{
	"mediastore",
}

func main() {
	m := db.NewMicro(svcName, port)
	m.RegisterPG()
	m.ReferServices(svcNames...)
	m.RegisterRedis()

	s := grpc.NewServer(grpc.UnaryInterceptor(worpc.UnaryInterceptorChain(worpc.Recovery, worpc.Logging)))
	pb.RegisterBookServiceServer(s, &service.BookServer{})
	//注册时间轮询
	c := cron.New()
	register.RegisterBookPolling(c)
	c.Start()
	defer c.Stop()

	s.Serve(m.CreateListener())
	service.HandlePendingBook()
}
