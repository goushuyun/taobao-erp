/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/07 17:34
 */

package misc

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/wothing/log"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

//returnNotToken 返回没找到token的错误提示
func ReturnNotToken(w http.ResponseWriter, r *http.Request) {

	RespondMessage(w, r, map[string]interface{}{
		"code":    errs.ErrTokenNotFound,
		"message": "need token but not found",
	})
}

// 随机字符串
func GenCheckCode(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	log.Debugf(">>>>>>>>>%s", result)
	return string(result)
}

func Md5String(objs ...interface{}) string {
	text := ""
	for i := range objs {
		text += fmt.Sprint(objs[i])
	}

	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func Contains(array []string, element string) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}
	return false
}

var reg = regexp.MustCompile("1\\d{10}")

func MobileFormat(mobile string) (string, error) {
	if !reg.MatchString(mobile) {
		return "", errs.NewError(errs.ErrMobileFormat, `The mobile should match 1\d{10}`)
	}

	return mobile, nil
}

func SuperPrint(x interface{}) string {
	buff := bytes.NewBuffer([]byte{})
	if err := encode(buff, reflect.ValueOf(x)); err != nil {
		return err.Error()
	}
	return buff.String()
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Ptr:
		buf.WriteByte('&')
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Struct:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "%s:", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Interface:
		return encode(buf, v.Elem())

	default: // complex, chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//FazzyQuery 模糊查询
func FazzyQuery(value string) string {
	var fazzy_value = "%"
	for _, char := range value {
		char := fmt.Sprintf("%c", char)
		if char != " " {
			fazzy_value += (char + "%")
		}
	}
	return fazzy_value
}

//记录日志
func LogErr(err error) {
	log.Debug("<<<<<<< new error <<<<<<<<<<<")
	log.Debugf("%+v", err)
	log.Debug("<<<<<<<           <<<<<<<<<<<")
}

//截取字符串
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
