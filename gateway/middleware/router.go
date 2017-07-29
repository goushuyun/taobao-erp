// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// wrap http.HandlerFunc to negroni.HandlerFunc
func Wrap(handler http.HandlerFunc) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler(rw, r)
		next(rw, r)
	}
}

type Router struct {
	prefix string
	*httprouter.Router
}

func NewWithPrefix(prefix string) *Router {
	return &Router{prefix: prefix, Router: httprouter.New()}
}

func (r *Router) Register(uri string, handlers ...negroni.HandlerFunc) {
	n := negroni.New()
	for _, h := range handlers {
		n.UseFunc(h)
	}
	r.Handler("POST", r.prefix+uri, n)
}

func (r *Router) RegisterGET(uri string, handlers ...negroni.HandlerFunc) {
	n := negroni.New()
	for _, h := range handlers {
		n.UseFunc(h)
	}
	r.Handler("GET", r.prefix+uri, n)
}
