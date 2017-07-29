// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package misc

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/goushuyun/weixin-golang/misc/token"
)

// GetTidFromContext 从 gPRC context 中取到 tid
func GetTidFromContext(ctx context.Context) string {
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			return md["tid"][0]
		}
	}
	return ""
}

// TODO remove
// ATTENTION none can not be delete
func GetUserIdFromToken(r *http.Request) string {
	if c := token.Get(r); c != nil {
		return c.UserId
	} else {
		return "none"
	}
}

// TODO remove
func GetHospitalIdFromToken(r *http.Request) string {
	if c := token.Get(r); c != nil {
		return c.HospId
	} else {
		return ""
	}
}

func GenContext(r *http.Request) context.Context {
	tid := r.Context().Value("tid").(string)
	return metadata.NewContext(context.Background(), metadata.Pairs("tid", tid))
}
