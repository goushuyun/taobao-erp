package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/goushuyun/weixin-golang/pb"
)

//AddGoods 增加商品
func AddGoods(w http.ResponseWriter, r *http.Request) {

	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{SellerId: c.SellerId, StoreId: c.StoreId}
	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "AddGoods", req, "book_id", "isbn", "location")
}

//UpdateGoods 更新商品
func UpdateGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{SellerId: c.SellerId, StoreId: c.StoreId}

	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "UpdateGoods", req)
}

//SearchGoods 搜索商品
func SearchGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{SellerId: c.SellerId, StoreId: c.StoreId}
	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "SearchGoods", req)
}

//GetGoodsTypeInfo 获取商品单类型基础类型
func GetGoodsTypeInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.TypeGoods{}
	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "GetGoodsTypeInfo", req)
}

//DelOrRemoveGoods  删除或者下架商品
func DelOrRemoveGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.DelGoodsReq{}

	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "DelOrRemoveGoods", req)
}

//GoodsLocationOperate 商品货架管理
func GoodsLocationOperate(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsLocation{}
	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "GoodsLocationOperate", req)
}

//AppSearchGoods 搜索图书 isbn 用于用户端搜索
func AppSearchGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{StoreId: c.StoreId}

	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "AppSearchGoods", req)
}

//AppSearchGoods 搜索图书 isbn 用于用户端搜索
func GoodsBactchUploadOperate(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadModel{StoreId: c.StoreId}

	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "GoodsBactchUploadOperate", req, "discount", "storehouse_id", "shelf_id", "floor_id", "origin_file", "origin_filename")
}

//GoodsBactchUploadList 获取批量上传数据
func GoodsBactchUploadList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadModel{StoreId: c.StoreId}

	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_goods", "GoodsBactchUploadList", req)
}
