package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
)

//CartAdd 增加购物车
func CartAdd(w http.ResponseWriter, r *http.Request) {
	req := &pb.Cart{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.StoreId = c.StoreId
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_cart", "CartAdd", req, "goods_id")

}

//CartList 购物车列表
func CartList(w http.ResponseWriter, r *http.Request) {
	req := &pb.Cart{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.StoreId = c.StoreId
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_cart", "CartList", req)

}

//CartUpdate 更改购物车
func CartUpdate(w http.ResponseWriter, r *http.Request) {
	req := &pb.CartUpdateReq{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.StoreId = c.StoreId
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_cart", "CartUpdate", req)

}

//CartDel 删除购物车
func CartDel(w http.ResponseWriter, r *http.Request) {
	req := &pb.Cart{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.StoreId = c.StoreId
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_cart", "CartDel", req)

}
