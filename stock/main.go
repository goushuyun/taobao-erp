package main

import (
	"github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/register"
	"github.com/robfig/cron"
	"github.com/wothing/worpc"
	"google.golang.org/grpc"

	"github.com/goushuyun/taobao-erp/stock/service"
)

const (
	svcName = "stock"
	port    = 10017
)

var svcNames = []string{
	"book",
	"mediastore",
}

func main() {
	m := db.NewMicro(svcName, port)
	m.RegisterPG()
	m.ReferServices(svcNames...)
	s := grpc.NewServer(grpc.UnaryInterceptor(worpc.UnaryInterceptorChain(worpc.Recovery, worpc.Logging)))
	pb.RegisterStockServiceServer(s, &service.StockServer{})

	//注册时间轮询
	c := cron.New()
	register.RegisterBookPolling(c)
	c.Start()
	defer c.Stop()

	s.Serve(m.CreateListener())
}
