/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/06 16:29
 */

package role

type Role int64

const (
	Mask4App   int64 = 0xff
	Mask4Inter int64 = 0xfffe00 // not f because InterAdmin can not be changed
	Mask4Hosp  int64 = 0xf000000
)

// 0~7 App(8)
// 8~23 Inter(16)
// 24~31 Hosp(8)

// [WARNING] ADD OR RENAME ONLY, DO NOT EDIT VALUE!!! MAX is 1 << 53 BECAUSE JS MAXIMUM
const (
	AppNormalUser Role = 1 << 0 //普通用户
	AppConfidant       = 1 << 1 //密友权限

	InterAdmin                     = 1 << 8  //超级管理员
	InterNormalUser                = 1 << 9  //后台普通用户
	InterFinanceSupervisor         = 1 << 10 //财务主管
	InterFinanceOfficer            = 1 << 11 //财务专员
	InterCustomerServiceSupervisor = 1 << 12 //客服主管
	InterCustomerServiceOfficer    = 1 << 13 //客服专员
	InterOperationSupervisor       = 1 << 14 //运行主管
	InterOperationOfficer          = 1 << 15 //运营专员
	InterSupplySupervisor          = 1 << 16 //供应链主管
	InterRiskSupervisor            = 1 << 17 //风控主管

	HospAdmin      = 1 << 24
	HospNormalUser = 1 << 25
)

func (ro Role) HasOne(roles ...Role) bool {
	if len(roles) == 0 {
		return true
	}
	for _, r := range roles {
		if ro&r != 0 {
			return true
		}
	}
	return false
}

func (ro Role) HasAll(roles ...Role) bool {
	for _, r := range roles {
		if ro&r == 0 {
			return false
		}
	}
	return true
}
