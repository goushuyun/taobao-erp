/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/16 10:47
 */

package db

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/wothing/log"
	wonaming "github.com/wothing/wonaming/etcd"
	"github.com/wothing/worc"
)

type micro struct {
	name          string
	port          int
	closeFuncChan chan closeFunc
	sync.WaitGroup
}

func NewMicro(name string, port int) *micro {
	err := wonaming.Register(name, *_host, GetPort(port), *_etcd, time.Second*15, 20)
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	m := &micro{name: name, port: port, closeFuncChan: make(chan closeFunc, 8)}
	m.Add(1)
	go m.destroy()
	return m
}

// SetLogLevel sets log level for log package, this should init as early as possible
func (m *micro) SetLogLevel(lvl int) {
	log.SetOutputLevel(lvl)
}

func (m *micro) RegisterPG() {
	InitPG(m.name)
	m.closeFuncChan <- DB.Close
}

func (m *micro) RegisterPGX() {
	InitPGX(m.name)
	m.closeFuncChan <- ClosePGX
}

func (m *micro) RegisterRedis() {
	InitRedis(m.name)
	m.closeFuncChan <- CloseRedis
}

func (m *micro) RegisterNSQPub() {
	InitNSQPub(m.name)
	m.closeFuncChan <- CloseNSQPub
}

func (m *micro) RegisterNSQSub() {
	InitNSQSub(m.name)
	m.closeFuncChan <- CloseNSQSub
}

func (m *micro) ReferServices(svcNames ...string) {
	worc.StartServiceConns(*_etcd, svcNames)
	m.closeFuncChan <- func() error {
		worc.CloseServiceConns()
		return nil
	}
}

func (m *micro) CreateListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", m.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("%s start listening at :%d", m.name, m.port)
	return lis
}

func (m *micro) Wait() {
	m.Wait()
}

type closeFunc func() error

func (m *micro) destroy() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-ch
	for cf := range m.closeFuncChan {
		cf()
	}
	m.Done()
}
