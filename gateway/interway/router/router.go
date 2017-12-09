package router

import (
	c "github.com/goushuyun/taobao-erp/gateway/controller"
	m "github.com/goushuyun/taobao-erp/gateway/middleware"
)

//SetRouterV1 设置seller的router
func SetRouterV1() *m.Router {
	v1 := m.NewWithPrefix("/v1")

	// mediastock
	v1.Register("/mediastore/get_up_token", m.Wrap(c.GetUpToken))

	// users
	v1.Register("/users/register", m.Wrap(c.Register))
	v1.Register("/users/check_user_exist", m.Wrap(c.CheckUserExist))
	v1.Register("/users/login", m.Wrap(c.Login))
	v1.Register("/users/change_passwod", m.Wrap(c.ChangePwd))

	// sms
	v1.Register("/sms/send_identifying_code", m.Wrap(c.SendIdentifyingCode))

	// book
	v1.Register("/book/get_book_info", m.Wrap(c.GetBookInfo))
	v1.Register("/book/get_local_book_info", m.Wrap(c.GetLocalBookInfo))
	v1.Register("/book/save_book_info", m.Wrap(c.SaveBook))
	v1.Register("/book/update_book_info", m.Wrap(c.UpdateBookInfo))
	v1.Register("/book/submit_audit", m.Wrap(c.SubmitBookAudit))
	v1.Register("/book/get_audit_list", m.Wrap(c.GetBookAuditRecord))
	v1.Register("/book/get_organized_audit_list", m.Wrap(c.GetOrganizedBookAuditList))
	v1.Register("/book/handle_book_audit_list", m.Wrap(c.HandleBookAudit))

	// stock
	v1.Register("/stock/get_location_id", m.Wrap(c.GetLocationId))
	v1.Register("/stock/save_map_row", m.Wrap(c.SaveMapRow))
	v1.Register("/stock/save_goods", m.Wrap(c.SaveGoods))
	v1.Register("/stock/update_map_row", m.Wrap(c.UpdateMapRow))
	v1.Register("/stock/location_fazzy_query", m.Wrap(c.LocationFazzyQuery))
	v1.Register("/stock/search_goods", m.Wrap(c.SearchGoods))
	v1.Register("/stock/update_goods_info", m.Wrap(c.UpdateGoodsInfo))
	v1.Register("/stock/list_goods_all_locations", m.Wrap(c.ListGoodsAllLocations))
	v1.Register("/stock/upload_goods_batch_data", m.Wrap(c.GoodsBatchUpload))
	v1.Register("/stock/save_goods_batch_upload_record", m.Wrap(c.SaveGoodsBatchUploadRecord))
	v1.Register("/stock/get_goods_batch_upload_records", m.Wrap(c.GetGoodsBatchUploadRecords))
	v1.Register("/stock/get_goods_pending_check_list", m.Wrap(c.GetGoodsPendingCheckList))
	v1.Register("/stock/deal_goods_pending_check", m.Wrap(c.DealWithGoodsPendingCheckList))
	v1.Register("/stock/get_location_stock", m.Wrap(c.GetLocationStock))
	v1.Register("/stock/update_location", m.Wrap(c.UpdateLocation))
	v1.Register("/stock/get_pending_gatherd_goods", m.Wrap(c.GetGoodsPendingGatherData))
	v1.Register("/stock/get_goods_shift_record", m.Wrap(c.GetGoodsShiftRecord))
	v1.Register("/stock/get_shift_record_latest_export_date", m.Wrap(c.GetShiftRocordExportDate))
	v1.RegisterGET("/stock/export_csv", m.Wrap(c.ExportCsv))
	v1.RegisterGET("/stock/export_goods_shift_record", m.Wrap(c.ExportGoodsShiftRecord))
	return v1
}
