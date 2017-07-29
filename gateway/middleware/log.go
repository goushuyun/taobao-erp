package middleware

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"github.com/urfave/negroni"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/hack"
	"github.com/wothing/log"
)

func LogMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// do not read body of socket.io and no log for this request
		if r.URL.Path == "/v1/ws/" {
			next(rw, r)
			return
		}

		start := time.Now()
		tid := uuid.New()
		rw.Header().Set("X-Request-ID", tid)

		version := r.Header.Get("X-App-Version") // iOS, Android
		if version == "" {
			version = "web"
		}

		// get real ip from nginx header, if not, try to remove port
		if realIp := r.Header.Get("X-Real-IP"); realIp != "" {
			r.RemoteAddr = realIp
		} else {
			if i := strings.LastIndex(r.RemoteAddr, ":"); i > 0 {
				r.RemoteAddr = r.RemoteAddr[:i]
			}
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Terrorf(tid, "error on reading request body, from=%v, method=%v, remote=%v, agent=%v, version=%s", r.RequestURI, r.Method, r.RemoteAddr, r.UserAgent(), version)
			misc.RespondMessage(rw, r, misc.NewErrResult(errs.ErrRequestFormat, "error on reading request body"))
			return
		}
		r.Body.Close()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "tid", tid)
		ctx = context.WithValue(ctx, "body", body)
		//ctx = context.WithValue(ctx, "ctx", metadata.NewContext(context.Background(), metadata.Pairs("tid", tid)))
		r = r.WithContext(ctx)

		bodyFormat := replaceHttpReqPassword(hack.String(body))
		log.Tinfof(tid, "started handling request, from=%v, method=%v, remote=%v, agent=%v, version=%s, body=%v", r.RequestURI, r.Method, r.RemoteAddr, r.UserAgent(), version, bodyFormat)
		next(rw, r)
		log.Tinfof(tid, "completed handling request, status=%v, took=%v", rw.(negroni.ResponseWriter).Status(), time.Since(start))
	}
}

func replaceHttpReqPassword(s string) string {
	if len(s) > 10000 {
		s = s[:10000]
	}

	match := `"password":"` //len=12

	if i := strings.Index(s, match); i != -1 {
		if j := strings.Index(s[i+12:], `"`); j != -1 {
			return s[:i+12] + "****" + s[i+12+j:]
		}
	}

	return s
}
