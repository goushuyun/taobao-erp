package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/wothing/log"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/tealeg/xlsx"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//提交订单
func OrderSubmit(w http.ResponseWriter, r *http.Request) {
	req := &pb.OrderSubmitModel{}
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
	misc.CallWithResp(w, r, "bc_order", "OrderSubmit", req, "mobile", "name", "address", "school_id")
}

//模拟支付成功
func PaySuccess(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	misc.CallWithResp(w, r, "bc_order", "PaySuccess", req)
}

//app端订单列表
func OrderListApp(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自用户
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
	misc.CallWithResp(w, r, "bc_order", "OrderList", req)
}

//seller端订单列表
func OrderListSeller(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	req.SearchType = 1
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "OrderList", req)
}

//打印订单
func PrintOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "PrintOrder", req)
}

//发货订单
func DeliverOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "DeliverOrder", req)
}

//配送订单
func DistributeOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "DistributeOrder", req)
}

//订单备注
func RemarkOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "RemarkOrder", req, "id")
}

//确认订单
func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
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
	misc.CallWithResp(w, r, "bc_order", "ConfirmOrder", req)
}

//确认订单
func OrderDetail(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
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
	misc.CallWithResp(w, r, "bc_order", "OrderDetail", req)
}

//确认订单
func OrderDetailSeller(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	//搜索类型 来自商家
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StoreId = c.StoreId
		req.SellerId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "OrderDetail", req)
}

//订单售后
func AfterSaleApply(w http.ResponseWriter, r *http.Request) {
	req := &pb.AfterSaleModel{}
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId

	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "AfterSaleApply", req, "order_id")
}

//订单售后
func CloseOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	c := token.Get(r)

	// get store_id
	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId

	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "CloseOrder", req, "id")
}

//订单售后
func UserCenterNecessaryOrderCount(w http.ResponseWriter, r *http.Request) {
	req := &pb.UserCenterOrderCount{}
	c := token.Get(r)

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
	misc.CallWithResp(w, r, "bc_order", "UserCenterNecessaryOrderCount", req)
}

//订单售后
func OrderShareOperation(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	c := token.Get(r)

	if c != nil && c.StoreId != "" && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "OrderShareOperation", req, "id")
}

//订单售后
func HandleAfterSaleOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.AfterSaleModel{}
	c := token.Get(r)

	if c != nil && c.StoreId != "" && c.SellerId != "" {
		req.StaffId = c.SellerId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "HandleAfterSaleOrder", req)
}

//订单售后
func AfterSaleOrderHandledResult(w http.ResponseWriter, r *http.Request) {
	req := &pb.AfterSaleModel{}
	c := token.Get(r)

	if c != nil && c.StoreId != "" {
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}
	misc.CallWithResp(w, r, "bc_order", "AfterSaleOrderHandledResult", req)
}

//重新计算订单数量
func RestatisticOrderNum(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	// c := token.Get(r)
	//
	// if c != nil && c.StoreId != "" {
	// } else {
	// 	misc.RespondMessage(w, r, map[string]interface{}{
	// 		"code":    errs.ErrTokenNotFound,
	// 		"message": "token not found",
	// 	})
	// 	return
	// }
	misc.CallWithResp(w, r, "bc_order", "RestatisticOrderNum", req)
}

//导出订单
func ExportDeliveryOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.Order{}
	body := r.FormValue("params")
	if err := misc.Bytes2Struct([]byte(body), req); err != nil {
		misc.RespondMessage(w, r, err)
		return
	}

	// Call RPC 请求订单详情
	resp, ctx := &pb.OrderListResp{}, misc.GenContext(r)
	err := misc.CallSVC(ctx, "bc_order", "ExportDeliveryOrderData", req, resp)
	if err != nil {
		misc.RespondMessage(w, r, err)
		return
	}
	// 开始写入Excel
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("发货单")
	if err != nil {
		log.Error(err)
		res := misc.NewErrResult(errs.ErrInternal, "Error on AddSheet")
		misc.RespondMessage(w, r, res)
		return
	}

	row = sheet.AddRow()
	head := []string{"订单编号", "收件人", "固话", "手机", "地址", "发货信息", "备注", "代收金额", "保价金额", "业务类型"}
	row.WriteSlice(&head, len(head))

	for _, detail := range resp.Data {
		row = sheet.AddRow()
		row.AddCell().SetString(detail.Order.Id)
		row.AddCell().SetString(detail.Order.Name)
		row.AddCell().SetString("")
		row.AddCell().SetString(detail.Order.Mobile)
		row.AddCell().SetString(detail.Order.Address)

		goodsInfo := ``
		for _, item := range detail.Items {
			itemInfo := item.BookIsbn
			itemInfo += " " + item.BookTitle
			if item.Type == 0 {
				// 新书
				itemInfo += "[新]" + ` `
			} else if item.Type == 1 {
				// 旧书
				itemInfo += "[旧]" + ` `
			}

			// 书本数量
			itemInfo += "x" + strconv.Itoa(int(item.Amount)) + `   `
			for _, location := range item.Locations {
				itemInfo += location.StorehouseName + "--" + location.ShelfName + "--" + location.FloorName + `   `
			}

			goodsInfo += itemInfo
		}

		row.AddCell().SetString(goodsInfo)
		row.AddCell().SetString(detail.Order.Remark)
		row.AddCell().SetString("")
		row.AddCell().SetString("")
		row.AddCell().SetString("")
	}

	filename := "发货单_" + time.Now().Format("20060102-15:04") + ".xlsx"
	w.Header().Set("Content-Disposition",
		`attachment; filename="`+filename+`"; filename*=utf-8''`+filename)

	w.Header().Set("Content-Type",
		`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8`)
	file.Write(w)
}

//导出订单
func ExportDistributeOrder(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("params")
	req := &pb.Order{}
	if err := misc.Bytes2Struct([]byte(body), req); err != nil {
		misc.RespondMessage(w, r, err)
		return
	}

	// Call RPC 请求订单详情
	resp, ctx := &pb.DistributeOrdersResp{}, misc.GenContext(r)
	err := misc.CallSVC(ctx, "bc_order", "ExportDistributeOrderData", req, resp)
	if err != nil {
		misc.RespondMessage(w, r, err)
		return
	}

	// 开始写入Excel
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("配货单")
	if err != nil {
		log.Error(err)
		res := misc.NewErrResult(errs.ErrInternal, "Error on AddSheet")
		misc.RespondMessage(w, r, res)
		return
	}

	row = sheet.AddRow()
	head := []string{"ISBN", "书名", "出版社", "数量", "类别", "库存位置"}
	row.WriteSlice(&head, len(head))

	for _, detail := range resp.Data {
		row = sheet.AddRow()
		row.AddCell().SetString(detail.Isbn)
		row.AddCell().SetString(detail.Title)
		row.AddCell().SetString(detail.Publisher)
		row.AddCell().SetInt(int(detail.Num))
		if detail.Type == 0 {
			// 新书
			row.AddCell().SetString("[新书]")

		} else if detail.Type == 1 {
			// 旧书
			row.AddCell().SetString("[旧书]")
		}

		row.AddCell().SetString(detail.Locations)

	}

	filename := "配货单_" + time.Now().Format("20060102-15:04") + ".xlsx"
	w.Header().Set("Content-Disposition",
		`attachment; filename="`+filename+`"; filename*=utf-8''`+filename)

	w.Header().Set("Content-Type",
		`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8`)
	file.Write(w)
}
