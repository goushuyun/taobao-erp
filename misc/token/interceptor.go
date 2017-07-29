package token

import (
	"github.com/goushuyun/weixin-golang/errs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Jwt(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token := getTokenFromContext(ctx)
	if token != "" {
		c, err := Check(token)
		if err != nil {
			return nil, errs.NewError(errs.ErrTokenFormat, "token illegal")
		}

		if c.VerifyIsExpired() {
			if c.VerifyCanRefresh() {
				grpc.SendHeader(ctx, metadata.Pairs("token", Refresh(c)))
			} else {
				return nil, errs.NewError(errs.ErrTokenRefreshExpired, "need relogin")
			}
		}

		ctx = context.WithValue(ctx, claimsKey{}, c)
	}
	return handler(ctx, req)
}

// GetTokenFromContext extract token from context
func getTokenFromContext(ctx context.Context) string {
	if md, ok := metadata.FromContext(ctx); ok {
		if len(md["token"]) > 0 {
			return md["token"][0]
		}
	}
	return ""
}

type claimsKey struct{}

func GetClaims(ctx context.Context) (Claims, bool) {
	c, ok := ctx.Value(claimsKey{}).(Claims)
	return c, ok
}
