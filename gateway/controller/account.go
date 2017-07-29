package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
)

//AddAddress 增加收货地址
func FindAccountItems(w http.ResponseWriter, r *http.Request) {
	req := &pb.FindAccountitemReq{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_account", "FindAccountItems", req)

}

//AddAddress 增加收货地址
func AccountStatistic(w http.ResponseWriter, r *http.Request) {
	req := &pb.Account{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_account", "AccountStatistic", req)

}
