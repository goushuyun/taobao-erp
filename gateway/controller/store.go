package controller

import (
	"goushuyun/errs"
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
	"github.com/wothing/log"
)

func GetRecyclingQrcode(w http.ResponseWriter, r *http.Request) {
	req := &pb.Store{}

	// get store_id
	if c := token.Get(r); c != nil {
		req.Id = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	log.Debugf("The store_id is : %s", req.Id)

	misc.CallWithResp(w, r, "bc_store", "GetStoreRecyclingQrcode", req)
}

//AddStore 增加店铺接口
func AddStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)
	req := &pb.Store{Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "AddStore", req, "name")
}

//UpdateStore 修改
func UpdateStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	storeid := c.StoreId
	req := &pb.Store{Id: storeid, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "UpdateStore", req, "name")
}

//AddRealStore 增加实体店
func AddRealStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.RealStore{StoreId: c.StoreId, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "AddRealStore", req, "name", "province_code", "city_code", "scope_code", "address", "images")
}

//UpdateRealStore 修改实体店信息
func UpdateRealStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.RealStore{Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "UpdateRealStore", req, "name", "province_code", "city_code", "scope_code", "address", "images")
}

//StoreInfo 获取云店铺信息
func StoreInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Store{Id: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "StoreInfo", req)
}

//StoreInfo 获取云店铺信息
func StoreInfoApp(w http.ResponseWriter, r *http.Request) {
	req := &pb.Store{}
	misc.CallWithResp(w, r, "bc_store", "StoreInfo", req)
}

//EnterStore 获取云店铺信息
func EnterStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)
	req := &pb.Store{Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "EnterStore", req, "id")
}

//ChangeStoreLogo 更改店铺logo
func ChangeStoreLogo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Store{Id: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "ChangeStoreLogo", req, "logo")
}

//RealStores 获取实体店列表
func RealStores(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Store{Id: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "RealStores", req)
}

//检查code
func CheckCode(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.RegisterModel{}
	misc.CallWithResp(w, r, "bc_store", "CheckCode", req, "mobile", "message_code")
}

//TransferStore 转让店铺
func TransferStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.TransferStoreReq{Store: &pb.Store{Id: c.StoreId}}
	misc.CallWithResp(w, r, "bc_store", "TransferStore", req, "mobile", "message_code")
}

//删除实体店
func DelRealStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	log.Debug(c)

	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.RealStore{StoreId: c.StoreId, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_store", "DelRealStore", req, "id")

}

//删除实体店
func GetCardOperSmsCode(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.SmsCardSubmitModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "GetCardOperSmsCode", req)

}

//保存提现账号
func SaveStoreWithdrawCard(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.StoreWithdrawCard{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "SaveStoreWithdrawCard", req, "card_no", "card_name", "username", "code")

}

//保存提现账号
func UpdateStoreWithdrawCard(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.StoreWithdrawCard{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "UpdateStoreWithdrawCard", req, "id")

}

//保存提现账号
func GetWithdrawCardInfoByStore(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.StoreWithdrawCard{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "GetWithdrawCardInfoByStore", req)

}

//店铺首页历史订单各个状态统计
func StoreHistoryStateOrderNum(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.StoreHistoryStateOrderNumModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "StoreHistoryStateOrderNum", req)

}

//提现申请
func WithdrawApply(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.StoreWithdrawalsModel{StoreId: c.StoreId, StaffId: c.SellerId}
	misc.CallWithResp(w, r, "bc_store", "WithdrawApply", req, "withdraw_card_id", "withdraw_fee")

}

//充值申请
func RechargeApply(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.RechargeModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "RechargeApply", req, "recharge_fee")
}

//充值完成处理
func RechargeHandler(w http.ResponseWriter, r *http.Request) {

	req := &pb.RechargeModel{}
	misc.CallWithResp(w, r, "bc_store", "RechargeHandler", req, "recharge_fee")
}

//检索店铺额外信息
func FindStoreExtraInfo(w http.ResponseWriter, r *http.Request) {

	req := &pb.StoreExtraInfo{}
	misc.CallWithResp(w, r, "bc_store", "FindStoreExtraInfo", req)
}

//同步店铺信息和店铺额外信息
func SyncStoreExtraInfo(w http.ResponseWriter, r *http.Request) {

	req := &pb.StoreExtraInfo{}
	misc.CallWithResp(w, r, "bc_store", "SyncStoreExtraInfo", req)
}

//修改店铺增加信息
func UpdateStoreExtraInfo(w http.ResponseWriter, r *http.Request) {
	req := &pb.StoreExtraInfo{}
	misc.CallWithResp(w, r, "bc_store", "UpdateStoreExtraInfo", req, "id")
}

//获取云店回收信息
func AccessStoreRecylingInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	var storeId string
	if c != nil && c.StoreId != "" {
		storeId = c.StoreId
	}
	req := &pb.Recyling{StoreId: storeId}
	misc.CallWithResp(w, r, "bc_store", "AccessStoreRecylingInfo", req)
}

//提交预约订单接口
func UserSubmitRecylingOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.RecylingOrder{}
	misc.CallWithResp(w, r, "bc_store", "UserSubmitRecylingOrder", req, "store_id", "school_id", "user_id", "mobile", "addr", "appoint_start_at", "appoint_end_at")
}

//查看预约中的回收订单接口
func UserAccessPendingRecylingOrder(w http.ResponseWriter, r *http.Request) {
	req := &pb.RecylingOrder{}
	misc.CallWithResp(w, r, "bc_store", "UserAccessPendingRecylingOrder", req, "user_id")
}

//设置云店回收信息
func UpdateStoreRecylingInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Recyling{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "UpdateStoreRecylingInfo", req, "id")
}

//设置云店回收信息
func GetStoreRecylingOrderList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.RecylingOrder{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "GetStoreRecylingOrderList", req)
}

//设置云店回收信息
func UpdateRecylingOrder(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.RecylingOrder{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "UpdateRecylingOrder", req, "id")
}

//保存或者新增订单快捷备注
func SaveOrUpdateOrderShortcutRemark(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.StoreExtraInfo{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "SaveOrUpdateOrderShortcutRemark", req)
}

//获取订单快捷备注列表
func GetOrderShortcutRemark(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Store{Id: c.StoreId}
	misc.CallWithResp(w, r, "bc_store", "GetOrderShortcutRemark", req)
}
