// Copyright (c) 2015 Wothing Co., Ltd. All rights reserved.
//
// Description:
//   API appway(gateway) redis file of 17mei.cn.
//
// Authors:
//   likun@wothing.com, 2015.11.37, initial creating.
//
package db

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/wothing/log"
)

var pool *redis.Pool

func newPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				//log.Errorf("connect to redis error : %s", err)
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					log.Fatalf("AUTH error: %s", err)
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				c.Close()
				log.Errorf("PING error: %s", err)
			}
			return nil
		},
	}
}

func InitRedis(svcName string) {
	redishost := GetValue(svcName, "redis/host", "127.0.0.1")
	redisport := GetValue(svcName, "redis/port", "6379")
	redispwd := GetValue(svcName, "redis/password", "")
	redis := fmt.Sprintf("%s:%s", redishost, redisport)
	setupRedis(redis, redispwd)
}

func setupRedis(server string, password string) {
	pool = newPool(server, password)
	err := pool.TestOnBorrow(pool.Get(), time.Now()) // time.Now() this parameter no use

	if err != nil {
		log.Fatalf("redis connect error: %s", err.Error())
	}
}

func CloseRedis() error {
	if pool != nil {
		return pool.Close()
	}
	return nil
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}
