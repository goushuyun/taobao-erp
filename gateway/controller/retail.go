package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

// 增加零售
func RetailSubmit(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.RetailSubmitModel{SellerId: c.SellerId, StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_retail", "RetailSubmit", req, "school_id", "total_fee")
}

//RetailList 零售检索
func RetailList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Retail{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_retail", "RetailList", req)
}
