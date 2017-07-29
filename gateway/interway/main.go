package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/urfave/negroni"
	"github.com/wothing/log"

	"github.com/goushuyun/weixin-golang/db"
	"github.com/goushuyun/weixin-golang/gateway/interway/router"
	m "github.com/goushuyun/weixin-golang/gateway/middleware"
)

const (
	svcName = "bc_interway"
	port    = 8848
)

var serviceNames = []string{
	"bc_seller",
	"bc_store",
	"bc_mediastore",
	"bc_school",
	"bc_location",
	"bc_books",
	"bc_goods",
	"bc_topic",
	"bc_weixin",
	"bc_circular",
	"bc_order",
	"bc_retail",
	"bc_statistic",
	"bc_account",
	"bc_payment",
	"bc_groupon",
}

func main() {
	defer log.Infof("%s stopped, bye bye !", svcName)
	runtime.GOMAXPROCS(runtime.NumCPU())

	micro := db.NewMicro(svcName, port)
	micro.ReferServices(serviceNames...)

	n := negroni.New()
	n.Use(m.RecoveryMiddleware())
	n.Use(m.LogMiddleware())
	n.Use(m.JWTMiddleware())
	// n.Use(m.TokenRequiredMiddle())
	n.UseHandler(router.SetRouterV1())

	networkAddr := fmt.Sprintf("0.0.0.0:%d", db.GetPort(port))
	log.Infof("%s servering on %s", svcName, networkAddr)
	log.Fatal(http.ListenAndServe(networkAddr, n))
}
