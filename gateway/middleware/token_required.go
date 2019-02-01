package middleware

import (
	"net/http"
	"strings"

	"github.com/goushuyun/taobao-erp/errs"

	"github.com/urfave/negroni"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/misc/token"
	"github.com/goushuyun/log"
)

func TokenRequiredMiddle() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		log.Debugf("The URL path is %s\n", r.URL.Path)

		// check whether token is exist, if not found, return it
		if strings.HasPrefix(r.URL.Path, "/v1/store") {
			if c := token.Get(r); c == nil {
				misc.RespondMessage(w, r, map[string]interface{}{
					"code":    errs.ErrTokenNotFound,
					"message": "need token but not found",
				})
			}
		}

		next(w, r)
	}
}
