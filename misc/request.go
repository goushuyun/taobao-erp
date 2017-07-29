// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/goushuyun/weixin-golang/errs"
)

// Deprecated: using Request2Struct instead
func Request2Map(body []byte) map[string]interface{} {
	if body == nil || bytes.Equal(body, []byte("null")) {
		return make(map[string]interface{})
	}

	var j interface{}
	err := json.Unmarshal(body, &j)
	if err != nil {
		return make(map[string]interface{})
	}

	data, ok := j.(map[string]interface{})
	if !ok {
		return make(map[string]interface{})
	}

	return data
}

func ReqStringCheck(data map[string]interface{}, paramNames ...string) error {
	for _, param := range paramNames {
		if data[param] == nil {
			return NewErrResult(errs.ErrRequestFormat, fmt.Sprintf("param '%s' is missing", param))
		}

		if _, ok := data[param].(string); !ok {
			return NewErrResult(errs.ErrRequestFormat, fmt.Sprintf("param '%s' not string", param))
		}
	}
	return nil
}

func CheckParams(in ...interface{}) error {
	var err = NewErrResult(errs.ErrRequestFormat, "param is missing")
	for _, v := range in {
		if v == nil {
			return err
		}
		t := reflect.TypeOf(v)
		v := reflect.ValueOf(v)
		switch t.Kind() {
		case reflect.String:
			if v.String() == "" {
				return err
			}
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			if v.Int() == 0 {
				return err
			}
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			if v.Uint() == 0 {
				return err
			}
		case reflect.Float32, reflect.Float64:
			if v.Float() == 0 {
				return err
			}
		case reflect.Slice:
			if v.Len() == 0 {
				return err
			}
		}
	}
	return nil
}
