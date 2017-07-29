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

	"github.com/goushuyun/weixin-golang/seller/role"
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

func TestSignAppToken(t *testing.T) {
	tokenStr = SignUserToken(AppToken, "0000000001", "170405000004")
	t.Log(tokenStr)

	claims, err := Check(tokenStr)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", claims)
}

func TestSign(t *testing.T) {
	tokenStr = SignSellerToken(InterToken, "00000004", "18817953402", "170424000006", role.InterAdmin)
	t.Log(">>>>>>>>>>>>>>>token>>>>>>>>>>>>>>")
	t.Logf("%s\n", tokenStr)
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	claims, err := Check(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", claims)
}

func TestCheck(t *testing.T) {
	tokenStr := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMDAwMDAwMDAxIiwic2VzIjoiZm40MCIsInNjcCI6MSwiaWF0IjoxNDkzMTEyNjA2LCJzdG9yZV9pZCI6IjE3MDQxMTAwMDAwMiIsImlzcyI6IjEuMCJ9.z43qCeJ6L3rNjgIWRHNRPavGZYtdni2xKyPeZjzyxXmm580YOaLGLkaPOMXxhusZMuyooh6suuS1b3KPPUq5SlGq3fVIKfkOJzEl-zHQHmAOMnkYVTXBl3Niq0hHZjaKxCKK1_wEQOIreP0GaYvmHTa7aDbdaLBviJUf7TaeF27uAS0i23-jKl6dFfleOKufyBQKkAuUfXcRXlmduLBK7RBUJ2p4CXF-vYNOZREiRmzuO2KPXORk-u0LQrpgmirK46uUD6xQvfD_OwUSmHs7rHDG1i2bco5m1lmlcKGIOhYO--XNrRSsQ8spWGDPCtwItX-xzuw7lDpGr3SAth9bzQ`

	c, err := Check(tokenStr)

	t.Logf("%+v\n", c)

	if err != nil {
		t.Error("check failed")
	}

	if c.VerifyIsExpired() != false {
		t.Error("should not expire")
	}

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
