package main

import (
	"github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
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
	s.Serve(m.CreateListener())
}
