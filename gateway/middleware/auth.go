/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/06 14:43
 */

package middleware

import (
	"net/http"

	"github.com/urfave/negroni"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/goushuyun/weixin-golang/seller/role"
)

// CheckAuth checks
// 1. if 'claims' exist
// 2. if scope match
// 3. if roles not empty, check if the role matches at least one, maybe using c.HasOneRole is better?
func CheckAuth(scope token.TokenType, roles ...role.Role) negroni.HandlerFunc {
	var rs role.Role = 0
	for _, r := range roles {
		rs = rs | r
	}

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		c := token.Get(r)
		if c == nil {
			misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenNotFound, "token not found"))
			return
		}

		if c.Scope != scope {
			misc.RespondMessage(rw, r, errs.NewError(errs.ErrAuth, "scope not match"))
			return
		}

		if rs != 0 && rs&c.Role == 0 {
			misc.RespondMessage(rw, r, errs.NewError(errs.ErrAuth, "unauthorized"))
			return
		}

		next(rw, r)
	}
}
