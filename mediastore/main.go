package main

import (
	"strings"

	"google.golang.org/grpc"

	"github.com/goushuyun/taobao-erp/mediastore/service"

	"github.com/goushuyun/taobao-erp/pb"

	"github.com/goushuyun/taobao-erp/db"

	"github.com/wothing/worpc"
)

const svcName = "mediastore"
const port = 10016

func main() {
	m := db.NewMicro(svcName, port)

	s := grpc.NewServer(grpc.UnaryInterceptor(worpc.UnaryInterceptorChain(worpc.Recovery, worpc.Logging)))
	test := strings.ToLower(db.GetValue(svcName, "mode", "test")) != "live"
	pb.RegisterMediastoreServer(s, &service.MediastoreServer{Test: test})
	s.Serve(m.CreateListener())
}
