package main

import (
	"strings"

	"google.golang.org/grpc"

	"github.com/goushuyun/weixin-golang/mediastore/service"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/goushuyun/weixin-golang/db"

	"github.com/wothing/worpc"
)

const svcName = "bc_mediastore"
const port = 8852

func main() {
	m := db.NewMicro(svcName, port)

	s := grpc.NewServer(grpc.UnaryInterceptor(worpc.UnaryInterceptorChain(worpc.Recovery, worpc.Logging)))
	test := strings.ToLower(db.GetValue(svcName, "mode", "test")) != "live"
	pb.RegisterMediastoreServer(s, &service.MediastoreServer{Test: test})
	s.Serve(m.CreateListener())
}
