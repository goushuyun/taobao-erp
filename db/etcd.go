/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Rod6 on 2016/06/15 09:08
 */

package db

import (
	"fmt"
	"strings"
	"sync"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

const (
	prefix = "bookcloud"
)

type etcdConn struct {
	api client.KeysAPI
}

var etcdconn *etcdConn = nil

var once sync.Once

func GetEtcdConn() client.KeysAPI {
	if etcdconn != nil {
		return etcdconn.api
	} else {
		once.Do(func() {
			endpoints := strings.Split(GetEtcd(), ",")
			conf := client.Config{
				Endpoints: endpoints,
			}
			cli, _ := client.New(conf)
			if cli != nil {
				etcdconn = &etcdConn{api: client.NewKeysAPI(cli)}
			}
		})
		return etcdconn.api
	}
}

// service - service name
// key - kv's key
// defaultValue - if cannot get any value, return it
func GetValue(service string, key string, defaultValue string) string {
	getKey := fmt.Sprintf("/%s/%s/%s", prefix, service, key)
	defaultKey := fmt.Sprintf("/%s/%s", prefix, key)

	v, err := GetEtcdConn().Get(context.Background(), getKey, nil)
	if err == nil && v.Node != nil && v.Node.Value != "" {
		return v.Node.Value
	}

	v, err = GetEtcdConn().Get(context.Background(), defaultKey, nil)
	if err == nil && v.Node != nil && v.Node.Value != "" {
		return v.Node.Value
	}

	return defaultValue
}
