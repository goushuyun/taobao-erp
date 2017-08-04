package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/errs"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/misc/token"
	"github.com/goushuyun/taobao-erp/pb"
)

func ListGoodsAllLocations(w http.ResponseWriter, r *http.Request) {
	req := &pb.Goods{}

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

	misc.CallWithResp(w, r, "stock", "ListGoodsAllLocations", req, "goods_id")
}

func LocationFazzyQuery(w http.ResponseWriter, r *http.Request) {
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

	misc.CallWithResp(w, r, "stock", "LocationFazzyQuery", req)
}

func UpdateMapRow(w http.ResponseWriter, r *http.Request) {
	req := &pb.MapRowBatch{}

	misc.CallWithResp(w, r, "stock", "UpdateMapRow", req, "data")
}

func SaveGoods(w http.ResponseWriter, r *http.Request) {
	req := &pb.Goods{}

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

	misc.CallWithResp(w, r, "stock", "SaveGoods", req, "book_id")
}

func SaveMapRow(w http.ResponseWriter, r *http.Request) {
	req := &pb.MapRow{}

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

	misc.CallWithResp(w, r, "stock", "SaveMapRow", req, "stock", "goods_id", "location_id")
}

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

//get book info
func SearchGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.GoodsInfo{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "SearchGoods", req)
}
