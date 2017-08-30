package service

import (
	"testing"

	"github.com/goushuyun/taobao-erp/db"
)

func TestRedis(t *testing.T) {
	db.InitRedis("sms")

	conn := db.GetRedisConn()
	val, err := conn.Do("get", "namdde")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", val)
}
