/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/05 17:35
 */

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/urfave/negroni"
	"github.com/wothing/log"

	"github.com/goushuyun/weixin-golang/db"
	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/token"
)

var whiteList = map[string]bool{
	"/v1/user/refresh_token": true,
}

// JWTJWTMiddleware check if Authorization not empty
func JWTMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			authHeader = r.URL.Query().Get("token")
		}

		if authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			switch len(authHeaderParts) {
			case 1:
				next(rw, r)
				return
			case 2:
				if strings.ToLower(authHeaderParts[0]) != "bearer" || authHeaderParts[1] == "" {
					misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenFormat, "token format error"))
					return
				}
			default:
				misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenFormat, "token length error"))
				return
			}

			c, err := token.Check(authHeaderParts[1])
			if err != nil {
				log.Warn("authHeader:", authHeader, "err:", err)
				misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenFormat, "token illegal"))
				return
			}

			// token version check
			if !c.VerifyVersion() {
				misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenRefreshExpired, "token version error, need relogin"))
				return
			}

			if c.VerifyIsExpired() {
				if c.VerifyCanRefresh() {
					if !whiteList[r.RequestURI] {
						rw.Header().Add("X-JWT-Token", token.Refresh(c))
					}
				} else {
					misc.RespondMessage(rw, r, errs.NewError(errs.ErrTokenRefreshExpired, "token is overdue, need relogin"))
					return
				}
			}

			// claims is ptr
			r = r.WithContext(context.WithValue(r.Context(), "claims", c))
		}

		next(rw, r)
	}
}

func SessionMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if c := token.Get(r); c != nil {
			if c.Session == "H5" {
				// do nothing if session is H5
			} else {
				rc := db.GetRedisConn()
				s, err := redis.String(rc.Do("get", "s:"+c.UserId))
				rc.Close()
				if err == nil && s != c.Session {
					misc.RespondMessage(rw, r, errs.NewError(errs.ErrSessionExpired, "please re-signin"))
					return
				}
			}
		}
		next(rw, r)
	}
}
