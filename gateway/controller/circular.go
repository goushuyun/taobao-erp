package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/errs"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/goushuyun/weixin-golang/pb"
)

//AddCircular 增加轮播图
func AddCircular(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_circular", "AddCircular", req, "image")
}

//DelCircular 删除轮播图
func DelCircular(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_circular", "DelCircular", req, "id")
}

//UpdateCircular 修改轮播图信息
func UpdateCircular(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_circular", "UpdateCircular", req, "id")
}

//CircularList 轮播图list
func CircularList(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_circular", "CircularList", req)
}

//CircularList 轮播图list
func CircularListApp(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	misc.CallWithResp(w, r, "bc_circular", "CircularList", req)
}

//CircularList 初始化轮播图
func CircularInit(w http.ResponseWriter, r *http.Request) {
	req := &pb.Circular{}
	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_circular", "CircularInit", req)
}
