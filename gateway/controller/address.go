package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
)

//AddAddress 增加收货地址
func AddAddress(w http.ResponseWriter, r *http.Request) {
	req := &pb.AddressReq{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_address", "AddAddress", req)

}

//AddAddress 增加收货地址
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	req := &pb.AddressInfo{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_address", "UpdateAddress", req, "id")

}

//AddAddress 增加收货地址
func MyAddresses(w http.ResponseWriter, r *http.Request) {
	req := &pb.AddressInfo{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_address", "MyAddresses", req)

}

//AddAddress 增加收货地址
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	req := &pb.AddressReq{}
	c := token.Get(r)
	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_address", "DeleteAddress", req)

}
