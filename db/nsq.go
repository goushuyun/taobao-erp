/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/14 16:34
 */

package db

import (
	"strings"

	nsq "github.com/nsqio/go-nsq"
	"github.com/wothing/log"
)

type nsqPub struct {
	*nsq.Producer
}

type nsqSub struct {
	nsqlAddr string
	config   *nsq.Config
	conns    []*nsq.Consumer
}

func (s *nsqSub) GetAddr() string {
	return s.nsqlAddr
}

func (s *nsqSub) GetConfig() *nsq.Config {
	return s.config
}

func (s *nsqSub) AddConn(conn *nsq.Consumer) {
	s.conns = append(s.conns, conn)
}

var NSQLogger nsqlogger

type nsqlogger struct {
}

func (nsqlogger) Output(calldepth int, s string) error {
	if i := strings.Index(s, "["); i != -1 {
		log.Error(s[i:])
	} else {
		log.Error(s)
	}
	return nil
}

var NSQPub *nsqPub
var NSQSub *nsqSub

func InitNSQPub(svcName string) {
	host := GetValue(svcName, "nsqd/host", "127.0.0.1")
	port := GetValue(svcName, "nsqd/port", "4150")
	nsqdAddr := host + ":" + port

	conf := nsq.NewConfig()

	p, err := nsq.NewProducer(nsqdAddr, conf)
	if err != nil {
		log.Fatal("InitNSQPub", nsqdAddr, err)
	}

	p.SetLogger(NSQLogger, nsq.LogLevelError)

	// we have NSQLogger, so do not care about ping error
	p.Ping()

	NSQPub = &nsqPub{
		p,
	}
}

func CloseNSQPub() error {
	if NSQPub != nil {
		NSQPub.Stop()
	}
	return nil
}

func InitNSQSub(svcName string) {
	host := GetValue(svcName, "nsql/host", "127.0.0.1")
	port := GetValue(svcName, "nsql/port", "4161")
	nsqlAddr := host + ":" + port

	conf := nsq.NewConfig()
	conf.MaxAttempts = 20 // MAX is 20 but we will give up if it reached 12

	NSQSub = &nsqSub{
		nsqlAddr: nsqlAddr,
		config:   conf,
	}
}

func CloseNSQSub() error {
	if NSQSub != nil && len(NSQSub.conns) != 0 {
		for i := range NSQSub.conns {
			NSQSub.conns[i].Stop()
		}
	}
	return nil
}
