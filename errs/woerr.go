// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package errs

import (
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	Ok = "00000"

	// Generic error definition, 10000-10999
	ErrInternal       = "10000"
	ErrInMaintain     = "10001"
	ErrRequestFormat  = "10002"
	ErrApiDeprecated  = "10003" // Api is deprectated, need to update app
	ErrSessionExpired = "10004"
	ErrInBlacklist    = "10005"
	ErrTempBlock      = "10006" // 暂时被禁止访问
	ErrAuth           = "10007"
	ErrUnimplemented  = "10008"
	ErrUnreachable    = "10009" // 第三方平台不可达

	// User/Auth error definition, 11000-11099
	ErrNicknameExists          = "11000"
	ErrUserNotFound            = "11001"
	ErrUserExists              = "11002"
	ErrPasswordWrong           = "11003"
	ErrUserBlacked             = "11004"
	ErrMobileUnmodified        = "11005"
	ErrMobileFormat            = "11006"
	ErrMobileNotSet            = "11007"
	ErrPasswordLength          = "11008"
	ErrPasswordFormat          = "11009"
	ErrNicknameStopWord        = "11010" // 包含违禁词
	ErrTokenNotFound           = "11011"
	ErrTokenFormat             = "11012"
	ErrTokenExpired            = "11013"
	ErrTokenRefreshExpired     = "11014"
	ErrRoleFormat              = "11015"
	ErrRoleWrong               = "11016"
	ErrNicknameFormat          = "11017"
	ErrUserBindWeChatForbidden = "11018"
	ErrUserBindMobileForbidden = "11019"

	//activity error definition,11200-11299
	ErrEnrollExists      = "11201"
	ErrConsultExists     = "11202"
	ErrActivityNotExists = "11203"
	ErrActivityEnrolled  = "11204"

	// Verification error definition, 11100-11119
	ErrSmsNotFound   = "11100"
	ErrWrongVeriCode = "11101"
	ErrSmsFrequent   = "11102"

	// diary and story, 11300-11399
	ErrStoryNotExist   = "11300"
	ErrLikedAlready    = "11301"
	ErrCommentYourself = "11302"

	// version control, 11400-11499
	ErrVersionNotFound        = "11400"
	ErrVersionDuplicate       = "11401"
	ErrVersionChannelNotFound = "11402"

	//statistic 11500-11599
	ErrStatisticSignIllegal = "11500"

	//config 11600-11699
	ErrConfigDuplicate     = "11600"
	ErrConfigScopeNotFound = "11601"
	ErrConfigNoScopeExist  = "11602"

	//scm 11700-11799
	ErrScmDuplicate = "11700"

	//confidant 11800-11899
	ErrConfidantActiveCodeNotExist    = "11800"
	ErrConfidantActiveCodeAlreadyUsed = "11801"
	ErrConfidantInviterExist          = "11802"
	ErrConfidantInviteDuplicate       = "11803"
	ErrConfidantInviteMax             = "11804"
	ErrConfidantInviteForbidden       = "11805"
	ErrConfidantAcceptForbidden       = "11806"
	ErrConfidantInviteeNotExist       = "11807"

	// hospital 11900-11999
	ErrHospitalDuplicate    = "11900"
	ErrHospitalNotFound     = "11901"
	ErrHospotalStaffNotBind = "11902"

	//goods 12000-12099
	ErrCollectedAlready      = "12000"
	ErrGoodsNotExist         = "12001"
	ErrGoodsStockEmpty       = "12002"
	ErrGoodsStatusNotMatched = "12003"

	// media 12100-12199
	ErrFileUploadFail = "12100"

	// order 12200-12299
	ErrOrderNotExist               = "12200"
	ErrOrderSourceNotMatched       = "12201"
	ErrOrderOwnerNotMatched        = "12202"
	ErrOrderStatusNotMatched       = "12203"
	ErrOrderHasAddedOn             = "12204"
	ErrAmountNotCorrect            = "12205"
	ErrOrderClosed                 = "12206"
	ErrOrderComplated              = "12207"
	ErrOrderSettleStatusNotMatched = "12208"

	// account 12300-12399
	ErrWithdrawTimesExceed   = "12300"
	ErrWithdrawAmountExceed  = "12301"
	ErrWithdrawTimeYetToCome = "12302"
	ErrAccountMissAlipay     = "12303"
	ErrAccountMissCash       = "12304"
	ErrInsufficientBalance   = "12305"

	// payment 12400-12499
	ErrCurrencyNotSupported = "12400"
	ErrUnknownHookType      = "12401"

	//easemob 12500-12599
	ErrEasemobUserTokenNone = "12501"
	ErrEasemobRegisterUser  = "12502"

	// coupon 12600-12699
	ErrCouponSoldOut  = "12600"
	ErrCouponUserSent = "12601"

	// leaflet 12700-12799
	ErrLeafletCounselorDuplicate = "12700"

	// proposal 12800-12899
	ErrDesireNotEnd       = "12800"
	ErrDesireScopeIllegal = "12801"

	// wechat 12900-12999
	ErrWeChatTemplateParams     = "12900"
	ErrWeChatNoSuchTemplate     = "12901"
	ErrWeChatUnAuth             = "12902"
	ErrWeChatUnionIdNotExist    = "12903"
	ErrWeChatQRCodeDupDuplicate = "12904"
)

var ErrAlerts = map[string]string{
	// Generic error definition, 10000-10999
	ErrInternal:       "发生未知错误，请稍后再试～",
	ErrInMaintain:     "服务器在维护中，请稍后再试～",
	ErrRequestFormat:  "网络请求错误",
	ErrApiDeprecated:  "当前版本过低，需要升级咯～",
	ErrSessionExpired: "用户已从其他设备登录",
	ErrInBlacklist:    "当前用户已无法访问服务",
	ErrTempBlock:      "您已暂时被禁止访问，请稍后再试",
	ErrUnimplemented:  "未实现的接口",

	// User/Auth error definition, 11000-11099
	ErrNicknameExists:          "昵称已存在",
	ErrUserNotFound:            "用户不存在",
	ErrUserExists:              "用户已存在",
	ErrPasswordWrong:           "密码错误",
	ErrUserBlacked:             "当前用户已进入黑名单",
	ErrMobileUnmodified:        "手机号码无法被更改",
	ErrMobileFormat:            "手机号码格式错误",
	ErrMobileNotSet:            "未绑定手机号码",
	ErrPasswordLength:          "密码长度错误",
	ErrPasswordFormat:          "密码格式错误",
	ErrNicknameStopWord:        "您的昵称包含敏感词",
	ErrUserBindWeChatForbidden: "当前微信号已被绑定",
	ErrUserBindMobileForbidden: "当前手机号已被绑定",

	ErrTokenNotFound:       "网络请求错误",
	ErrTokenFormat:         "网络请求错误",
	ErrTokenExpired:        "网络请求错误",
	ErrTokenRefreshExpired: "TOKEN过期,请重新登录",

	ErrRoleFormat:     "请求权限错误",
	ErrRoleWrong:      "请求权限错误",
	ErrNicknameFormat: "昵称不合法",

	//activity error definition,11200-11299
	ErrEnrollExists:      "无法重复报名", // the user had enrolled this activity
	ErrActivityNotExists: "该活动不存在", // the activity does not exist
	ErrActivityEnrolled:  "用户已报名",

	// Verification error definition, 11100-11119
	ErrSmsNotFound:   "验证码未发送或已失效",
	ErrWrongVeriCode: "验证码错误",
	ErrSmsFrequent:   "验证码发送过于频繁",

	// diary and story, 11300-11399
	ErrStoryNotExist: "请求的美途故事不存在",
	ErrLikedAlready:  "请不要重复点赞",

	// version control, 11400-11499
	ErrVersionNotFound:        "版本未找到",
	ErrVersionDuplicate:       "版本重复",
	ErrVersionChannelNotFound: "版本所在渠道未发现",

	//statistic 11500-11599
	ErrStatisticSignIllegal: "签名错误",

	//config 11600-11699
	ErrConfigDuplicate:     "配置信息重复",
	ErrConfigScopeNotFound: "配置信息未发现",

	//scm 11700-11799
	ErrScmDuplicate: "供应链已存在",

	//confidant 11800-11899
	ErrConfidantActiveCodeNotExist:    "邀请码不存在",
	ErrConfidantActiveCodeAlreadyUsed: "邀请码已被使用",
	ErrConfidantInviterExist:          "您邀请的用户已经有邀请人啦",
	ErrConfidantInviteDuplicate:       "您已经邀请过该用户",
	ErrConfidantInviteMax:             "您已达最大好友邀请上限",
	ErrConfidantInviteForbidden:       "不能添加自己的推荐人为密友",
	ErrConfidantAcceptForbidden:       "您已有邀请人",
	ErrConfidantInviteeNotExist:       "推荐人不存在",

	ErrHospotalStaffNotBind: "登录用户未绑定医院",

	// goods 12000-12099
	ErrCollectedAlready:      "请不要重复收藏",
	ErrGoodsNotExist:         "商品不存在",
	ErrGoodsStockEmpty:       "库存不足",
	ErrGoodsStatusNotMatched: "商品状态不匹配，不能执行此操作",

	// media 12100-12199
	ErrFileUploadFail: "文件上传失败",

	// order 12200-12299
	ErrOrderNotExist:               "订单不存在",
	ErrOrderSourceNotMatched:       "订单来源不匹配",
	ErrOrderOwnerNotMatched:        "订单不存在",
	ErrOrderStatusNotMatched:       "订单状态不适合当前操作",
	ErrOrderHasAddedOn:             "订单已经追加过，不能重复追加",
	ErrAmountNotCorrect:            "付款金额不正确",
	ErrOrderClosed:                 "该订单超时未支付，请重新下单",
	ErrOrderComplated:              "订单已完成",
	ErrOrderSettleStatusNotMatched: "清算状态不匹配",
}

type Error struct {
	Code    string `json:"code"`              // 应答码
	Message string `json:"message,omitempty"` // 错误消息，如果请求成功，不返回这个字段
	Alert   string `json:"alert,omitempty"`   // 错误提示，用于客户端提醒用户
}

func (r *Error) Error() string {
	return `code: ` + r.Code + `, message:` + r.Message
}

func (r *Error) String() string {
	return r.Error()
}

func NewError(code, format string, a ...interface{}) error {
	return &Error{Code: code, Message: fmt.Sprintf(format, a...), Alert: ErrAlerts[code]}
}

func NewRpcError(code, format string, a ...interface{}) error {
	c, err := strconv.Atoi(code)
	if err != nil {
		c = 10000 // if code is not number, set to ErrInternal
	}

	return grpc.Errorf(codes.Code(c), format, a...)
}

func FromRpcError(err error) error {
	if err == nil {
		return nil
	}

	code := grpc.Code(err)
	if code < 10000 {
		// Not a 17mei error code
		return NewError(ErrInternal, err.Error())
	}

	return NewError(fmt.Sprintf("%05d", code), grpc.ErrorDesc(err))
}

// Wrap convert Error to grpc error, most of them are Error
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case *Error:
		return NewRpcError(e.Code, e.Message)
	default:
		if grpc.Code(err) == codes.Unknown {
			return NewRpcError(ErrInternal, e.Error())
		} else {
			return err
		}
	}
}
