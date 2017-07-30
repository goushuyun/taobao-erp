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
	InterAdmin      = 1 << 8 //超级管理员
	InterNormalUser = 1 << 9 //后台普通用户
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
