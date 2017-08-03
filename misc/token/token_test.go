/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/05 23:17
 */

package token

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	ro "github.com/goushuyun/taobao-erp/users/role"
)

var tokenStr string
var session string

func TestZeroValue(t *testing.T) {
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)
	t.Log(">>>>>>>>>>>>>>>>>>>>>")

	t.Logf("%+v\n", v)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Printf("%+v\n", values)

	// val := ""
	// t.Log(IsZeroOfUnderlyingType(val))
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	log.Println("------", reflect.TypeOf(x), "----------")

	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func TestSign(t *testing.T) {
	tokenStr = SignUserToken(InterToken, "u_17073000001", "18817953402", "WangKai", ro.InterAdmin)
	t.Log(">>>>>>>>>>>>>>>token>>>>>>>>>>>>>>")
	t.Logf("%s\n", tokenStr)
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	claims, err := Check(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	if claims.VerifyIsExpired() != false {
		t.Error("should not expire")
	}
	t.Logf("%+v", claims)
}

func TestCheck(t *testing.T) {
	tokenStr := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1XzE3MDgwMjAwMDAxIiwibW9iIjoiMTc2MDA5MDkwNDAiLCJuYW1lIjoi5Yav5b-g5qOuIiwic2VzIjoiZm4yaiIsInNjcCI6NTEyLCJyb2wiOjUxMiwiaWF0IjoxNTAxNzUwODQ4LCJpc3MiOiIxLjAifQ.j3HLuI0T-omdiTD0GF2TY8DtU6KBsdL5bMwHJXEs5zrsc2brz0U2LIM7UQVIGiPrNnKpJ1heCj3ZXuikGVIuxwD-bTU97b-e1sFjXwajXCd6679G0pJucw1aQ2vkekYjjeDa5QLq0hGJ7XZgodrAMMBUgH7UyejqjBlR3BbJraTbhLZpXtNdYjdvtcEHMqva3xFARAM714qxaWjyBkFfXYSWM0UvqUNl0WP1E6IRIOVdu1lTbtEWkRTU6gE-VbTs_4lNfJN_NU53V74ymNKawG62TwHwA8blaDVip6nO5ELGmUaF_uTXV1bm-IyDdZ8HIFrgUo4sghxoSWqnn-h30g`

	c, err := Check(tokenStr)

	t.Logf("%+v\n", c)

	if err != nil {
		t.Error("check failed")
	}

	if c.VerifyIsExpired() != false {
		t.Error("should not expire")
	}
	t.Log(c.VerifyIsExpired())

	if c.VerifyCanRefresh() != true {
		t.Error("can be refresh")
	}

	temp := tokenStr + "xxxx"
	_, err = Check(temp)
	if err == nil {
		t.Error("should be illegal error")
	}
}

func TestRefresh(t *testing.T) {
	//before := tokenStr
	//<-time.After(time.Second)
	//err := Refresh(&tokenStr)
	//if err != nil {
	//	t.Error("refresh failed")
	//}
	//if before == tokenStr {
	//	t.Error("refresh failed")
	//}
}
