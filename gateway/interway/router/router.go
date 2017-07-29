package router

import (
	c "github.com/goushuyun/weixin-golang/gateway/controller"
	m "github.com/goushuyun/weixin-golang/gateway/middleware"
)

//SetRouterV1 设置seller的router
func SetRouterV1() *m.Router {
	v1 := m.NewWithPrefix("/v1")
	// payment
	v1.Register("/payment/get_charge", m.Wrap(c.GetCharge))

	// weixin
	v1.Register("/weixin/get_auth_url", m.Wrap(c.GetAuthURL))
	v1.Register("/weixin/get_api_query_auth", m.Wrap(c.GetApiQueryAuth))
	v1.Register("/weixin/get_office_account_info", m.Wrap(c.GetOfficeAccountInfo))

	// books
	v1.Register("/books/get_book_info_by_isbn", m.Wrap(c.GetBookInfoByISBN))
	v1.Register("/books/modify_book_info", m.Wrap(c.ModifyBookInfo))
	v1.Register("/books/insert_book", m.Wrap(c.InsertBook))

	//seller 开始
	v1.Register("/seller/login", m.Wrap(c.SellerLogin))
	v1.Register("/seller/register", m.Wrap(c.SellerRegister))
	v1.Register("/seller/check_mobile", m.Wrap(c.CheckMobileExist))
	v1.Register("/seller/get_sms", m.Wrap(c.GetTelCode))
	v1.Register("/seller/self_stores", m.Wrap(c.SelfStores))
	v1.Register("/seller/get_update_sms", m.Wrap(c.GetUpdateTelCode))
	v1.Register("/seller/update_password", m.Wrap(c.UpdatePasswordAndLogin))
	//store 开始
	v1.Register("/store/add", m.Wrap(c.AddStore))
	v1.Register("/store/update", m.Wrap(c.UpdateStore))
	v1.Register("/store/add_real_store", m.Wrap(c.AddRealStore))
	v1.Register("/store/update_real_store", m.Wrap(c.UpdateRealStore))
	v1.Register("/store/store_info", m.Wrap(c.StoreInfo))
	v1.Register("/store/enter_store", m.Wrap(c.EnterStore))
	v1.Register("/store/change_logo", m.Wrap(c.ChangeStoreLogo))
	v1.Register("/store/real_stores", m.Wrap(c.RealStores))
	v1.Register("/store/check_code", m.Wrap(c.CheckCode))
	v1.Register("/store/transfer_store", m.Wrap(c.TransferStore))
	v1.Register("/store/del_real_store", m.Wrap(c.DelRealStore))
	v1.Register("/store/get_card_ope_sms", m.Wrap(c.GetCardOperSmsCode))
	v1.Register("/store/save_card", m.Wrap(c.SaveStoreWithdrawCard))
	v1.Register("/store/update_card", m.Wrap(c.UpdateStoreWithdrawCard))
	v1.Register("/store/get_store_card", m.Wrap(c.GetWithdrawCardInfoByStore))
	v1.Register("/store/index_order_num_statistic", m.Wrap(c.StoreHistoryStateOrderNum))
	v1.Register("/store/withdraw_apply", m.Wrap(c.WithdrawApply))
	v1.Register("/store/recharge_apply", m.Wrap(c.RechargeApply))
	v1.Register("/store/recharge_handler", m.Wrap(c.RechargeHandler)) //上线需要去掉
	v1.Register("/store/save_or_update_shortcut_remark", m.Wrap(c.SaveOrUpdateOrderShortcutRemark))
	v1.Register("/store/get_shortcut_remark_list", m.Wrap(c.GetOrderShortcutRemark))
	v1.Register("/store/get_cecycling_qrcode", m.Wrap(c.GetRecyclingQrcode))
	// mediastore
	v1.Register("/mediastore/get_upload_token", m.Wrap(c.GetUplaodToken))
	v1.Register("/mediastore/refresh_urls", m.Wrap(c.RefreshUrls))

	//school
	v1.Register("/school/add", m.Wrap(c.AddSchool))
	v1.Register("/school/update", m.Wrap(c.UpdateSchool))
	v1.Register("/school/update_express_fee", m.Wrap(c.UpdateExpressFee))
	v1.Register("/school/store_schools", m.Wrap(c.StoreSchools))
	v1.Register("/school/del", m.Wrap(c.DelSchool))

	// location
	v1.Register("/location/add_location", m.Wrap(c.AddLocation))
	v1.Register("/location/update_location", m.Wrap(c.UpdateLocation))
	v1.Register("/location/list_location", m.Wrap(c.ListLocation))
	v1.Register("/location/list_children_location", m.Wrap(c.GetChildrenLocation))
	v1.Register("/location/del_location", m.Wrap(c.DelLocation))

	//goods
	v1.Register("/goods/add", m.Wrap(c.AddGoods))
	v1.Register("/goods/update", m.Wrap(c.UpdateGoods))
	v1.Register("/goods/search", m.Wrap(c.SearchGoods))
	v1.Register("/goods/del_or_remove_goods", m.Wrap(c.DelOrRemoveGoods))
	v1.Register("/goods/goods_location_operate", m.Wrap(c.GoodsLocationOperate))
	v1.Register("/goods/batch_upload", m.Wrap(c.GoodsBactchUploadOperate))
	v1.Register("/goods/batch_list", m.Wrap(c.GoodsBactchUploadList))

	//topic
	v1.Register("/topic/add", m.Wrap(c.AddTopic))
	v1.Register("/topic/del", m.Wrap(c.DelTopic))
	v1.Register("/topic/update", m.Wrap(c.UpdateTopic))
	v1.Register("/topic/add_item", m.Wrap(c.AddTopicItem))
	v1.Register("/topic/del_item", m.Wrap(c.DelTopicItem))
	v1.Register("/topic/search", m.Wrap(c.SearchTopics))
	v1.Register("/topic/topics_info", m.Wrap(c.TopicsInfo))

	//circular
	v1.Register("/circular/add", m.Wrap(c.AddCircular))
	v1.Register("/circular/update", m.Wrap(c.UpdateCircular))
	v1.Register("/circular/del", m.Wrap(c.DelCircular))
	v1.Register("/circular/list", m.Wrap(c.CircularList))
	v1.Register("/circular/init", m.Wrap(c.CircularInit))

	//order
	v1.Register("/order/print", m.Wrap(c.PrintOrder))
	v1.Register("/order/deliver", m.Wrap(c.DeliverOrder))
	v1.Register("/order/distribute", m.Wrap(c.DistributeOrder))
	v1.Register("/order/search", m.Wrap(c.OrderListSeller))
	v1.Register("/order/detail", m.Wrap(c.OrderDetailSeller))
	v1.Register("/order/handle_after_sale", m.Wrap(c.HandleAfterSaleOrder))
	v1.Register("/order/after_sale_result", m.Wrap(c.AfterSaleOrderHandledResult))
	v1.RegisterGET("/order/export_delivery_order", m.Wrap(c.ExportDeliveryOrder))
	v1.RegisterGET("/order/export_distribute_order", m.Wrap(c.ExportDistributeOrder))
	v1.Register("/order/restatistic_order_num", m.Wrap(c.RestatisticOrderNum))
	v1.Register("/order/order_remark", m.Wrap(c.RemarkOrder))

	//retail
	v1.Register("/retail/add", m.Wrap(c.RetailSubmit))
	v1.Register("/retail/find", m.Wrap(c.RetailList))

	//statistic
	v1.Register("/statistic/get_today_sales", m.Wrap(c.StatisticToday))
	v1.Register("/statistic/get_total_sales", m.Wrap(c.StatisticTotal))
	v1.Register("/statistic/get_daliy_sales", m.Wrap(c.StatisticDaliy))
	v1.Register("/statistic/get_month_sales", m.Wrap(c.StatisticMonth))

	//account
	v1.Register("/account/find_list", m.Wrap(c.FindAccountItems))
	v1.Register("/account/seller_account", m.Wrap(c.AccountStatistic))

	//recyling
	v1.Register("/recyling/store_recyling_info", m.Wrap(c.AccessStoreRecylingInfo))
	v1.Register("/recyling/update_store_recyling_info", m.Wrap(c.UpdateStoreRecylingInfo))
	v1.Register("/recyling/recyling_order_list", m.Wrap(c.GetStoreRecylingOrderList))
	v1.Register("/recyling/update_recyling_order", m.Wrap(c.UpdateRecylingOrder))
	v1.Register("/recyling/update_school_recyling_state", m.Wrap(c.UpdateSchoolRecylingState))

	//Groupon
	v1.Register("/groupon/shared_major_batch_save", m.Wrap(c.SharedMajorBatchSave))
	v1.Register("/groupon/search_shared_major", m.Wrap(c.SharedMajorList))
	v1.Register("/groupon/save_institute", m.Wrap(c.SaveSchoolInstitute))
	v1.Register("/groupon/save_institute_major", m.Wrap(c.SaveInstituteMajor))
	v1.Register("/groupon/get_school_majors", m.Wrap(c.GetSchoolMajorInfo))
	v1.Register("/groupon/save_groupon", m.Wrap(c.SaveGroupon))
	v1.Register("/groupon/find_groupon", m.Wrap(c.GrouponList))
	v1.Register("/groupon/my_groupon", m.Wrap(c.MyGroupon))
	v1.Register("/groupon/groupon_items", m.Wrap(c.GetGrouponItems))
	v1.Register("/groupon/groupon_related_user", m.Wrap(c.GetGrouponPurchaseUsers))
	v1.Register("/groupon/groupon_log", m.Wrap(c.GetGrouponOperateLog))
	v1.Register("/groupon/update_groupon", m.Wrap(c.UpdateGruopon))
	v1.Register("/groupon/reset_expire_at", m.Wrap(c.BatchUpdateGrouponExpireAt))
	v1.Register("/groupon/del_institute_major", m.Wrap(c.DelInstituMajor))
	v1.Register("/groupon/update_institute_major", m.Wrap(c.UpdateInstituteMajor))
	v1.Register("/groupon/del_school_institute", m.Wrap(c.DelSchoolInstitute))
	v1.Register("/groupon/update_school_institute", m.Wrap(c.UpdateSchoolInstitute))

	return v1
}
