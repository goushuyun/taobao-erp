// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package misc

import (
	"fmt"

	"github.com/goushuyun/weixin-golang/errs"
)

// Result 所有请求都返回这个类型，其中 data 是返回的实际内容
type Result struct {
	// 应答码
	Code string `json:"code"`

	// 错误消息，如果请求成功，不返回这个字段
	Message string `json:"message,omitempty"`

	// 错误提示，用于客户端提醒用户
	Alert string `json:"alert,omitempty"`

	// 分页时返回总条数
	Total int64 `json:"total,omitempty"`

	// 返回数据，如果请求错误，不返回这个字段
	Data interface{} `json:"data,omitempty"`
}

func (r *Result) Error() string {
	return fmt.Sprintf(`code: %s, message: %s`, r.Code, r.Message)
}

func (r *Result) String() string {
	return r.Error()
}

var OkResult = &Result{Code: errs.Ok}

// NewErrResult 创建错误 Result
func NewErrResult(code, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Alert:   errs.ErrAlerts[code],
	}
}
