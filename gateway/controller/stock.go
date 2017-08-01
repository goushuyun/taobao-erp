package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/errs"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/misc/token"
	"github.com/goushuyun/taobao-erp/pb"
)

func GetLocationId(w http.ResponseWriter, r *http.Request) {
	req := &pb.Location{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "GetLocationId", req, "warehouse", "shelf", "floor")
}
