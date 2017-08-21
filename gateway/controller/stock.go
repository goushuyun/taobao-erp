package controller

import (
	"net/http"
	"time"

	log "github.com/wothing/log"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/tealeg/xlsx"

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

// get goods pending gathered
func GetShiftRocordExportDate(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.User{Id: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetShiftRocordExportDate", req)
}

// get goods pending gathered
func ExportGoodsShiftRecord(w http.ResponseWriter, r *http.Request) {

	req := &pb.GoodsShiftRecord{}
	body := r.FormValue("params")
	if err := misc.Bytes2Struct([]byte(body), req, "user_id", "start_at", "end_at"); err != nil {
		misc.RespondMessage(w, r, err)
		return
	}

	// Call RPC 请求订单详情
	resp, ctx := &pb.GoodsShiftRecordListResp{}, misc.GenContext(r)
	err := misc.CallSVC(ctx, "stock", "GetGoodsShiftRecord", req, resp)
	if err != nil {
		misc.RespondMessage(w, r, err)
		return
	}
	// 开始写入Excel
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("出库单")
	if err != nil {
		log.Error(err)
		res := misc.NewErrResult(errs.ErrInternal, "Error on AddSheet")
		misc.RespondMessage(w, r, res)
		return
	}

	row = sheet.AddRow()
	head := []string{"isbn", "书名", "库位", "架位", "层数", "数量", "日期", "业务类型"}
	row.WriteSlice(&head, len(head))

	for _, item := range resp.Data {
		row = sheet.AddRow()
		var book_no string
		if item.BookNo != "" {
			book_no = "_" + item.BookNo
		}
		row.AddCell().SetString(item.Isbn + book_no)
		row.AddCell().SetString(item.BookTitle)
		row.AddCell().SetString(item.Warehouse)
		row.AddCell().SetString(item.Shelf)
		row.AddCell().SetString(item.Floor)

		row.AddCell().SetInt64(item.Stock)
		row.AddCell().SetString(time.Unix(item.CreateAt, 0).Format("2006-01-02 15:04:05"))
		if item.OperateType == "load" {
			row.AddCell().SetString("入库")
		} else {
			row.AddCell().SetString("出库")
		}
	}

	user := &pb.User{Id: req.UserId, ExportStartAt: req.StartAt, ExportEndAt: req.EndAt}
	misc.CallSVC(ctx, "stock", "UpdateShiftRocordExportDate", user, &pb.NormalResp{})

	filename := "出库单_" + time.Now().Format("2006年01月02日") + ".xlsx"
	w.Header().Set("Content-Disposition",
		`attachment; filename="`+filename+`"; filename*=utf-8''`+filename)

	w.Header().Set("Content-Type",
		`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8`)
	file.Write(w)

}
