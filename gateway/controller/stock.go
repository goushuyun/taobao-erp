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

// update the goods info
func UpdateGoodsInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Goods{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "UpdateGoodsInfo", req, "goods_id", "remark")
}

// update the goods info
func GoodsBatchUpload(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadModel{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GoodsBatchUpload", req)
}

// add a batch upload record
func SaveGoodsBatchUploadRecord(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "SaveGoodsBatchUploadRecord", req, "origin_file", "origin_filename")
}

// get the batch upload record list
func GetGoodsBatchUploadRecords(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsBatchUploadRecords", req)
}

// get the pending goods check list
func GetGoodsPendingCheckList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsPendingCheck{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsPendingCheckList", req)
}

// deal with the goods check
func DealWithGoodsPendingCheckList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsPendingCheck{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "DealWithGoodsPendingCheckList", req, "id")
}

// get the location stock
func GetLocationStock(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Location{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetLocationStock", req)
}

// get the location stock
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Location{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "UpdateLocation", req, "location_id", "warehouse", "shelf", "floor")
}

// get goods pending gathered
func GetGoodsPendingGatherData(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsPendingGatherData", req)
}

// get goods pending gathered
func GetGoodsShiftRecord(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsShiftRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsShiftRecord", req)
}
