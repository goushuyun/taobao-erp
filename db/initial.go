/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/15 14:12
 */

package db

import (
	"flag"
	"os"
	"time"

	"github.com/wothing/log"
	wonaming "github.com/wothing/wonaming/etcd"

	"github.com/goushuyun/weixin-golang/ver"
)

var (
	_host         = flag.String("h", "0.0.0.0", "service listening host")
	_port         = flag.Int("p", 0, "service listening port")
	_etcd         = flag.String("etcd", "http://127.0.0.1:2379", "etcd host list, split by ','")
	_buildversion = flag.Bool("v", false, "show build version")
)

var _DEBUG = "FALSE"

func DEBUGMODE() bool {
	return _DEBUG != "FALSE"
}

func init() {
	flag.Parse()

	if DEBUGMODE() {
		log.SetOutputLevel(int(log.Ldebug))
		log.Debug("DEBUG MODE")
	}

	if *_buildversion {
		log.Infof("Git commit: %s", ver.GitCommit)
		log.Infof("Build time: %s", ver.BuildDate)
		os.Exit(0)
	}
}

func GetEtcd() string {
	return *_etcd
}

func GetPort(defaultPort int) int {
	if *_port == 0 {
		return defaultPort
	}
	return *_port
}

func RegisterService(svcName string, port int) error {
	return wonaming.Register(svcName, *_host, GetPort(port), GetEtcd(), time.Second*15, 20)
}
