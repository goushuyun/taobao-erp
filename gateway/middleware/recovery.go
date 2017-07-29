package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/urfave/negroni"
	"github.com/wothing/log"
)

const (
	printStack bool = false
	stackAll   bool = false
	stackSize  int  = 1024 * 8
)

func RecoveryMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)

				stack := make([]byte, stackSize)
				stack = stack[:runtime.Stack(stack, stackAll)]

				log.Errorf("panic http request: %s, err=%s, stack:\n%s", r.RequestURI, err, string(stack))

				if printStack {
					fmt.Fprintf(rw, "PANIC: %s\n%s", err, stack)
				} else {
					fmt.Fprint(rw, "internal error")
				}
			}
		}()

		next(rw, r)
	}
}
