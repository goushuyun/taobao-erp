package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/goushuyun/weixin-golang/pb"

	"github.com/goushuyun/weixin-golang/misc"
)

//今日销售额统计
func StatisticToday(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsSalesStatisticModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_statistic", "StatisticToday", req)
}

//历史总销售额统计
func StatisticTotal(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsSalesStatisticModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_statistic", "StatisticTotal", req)
}

//历史日统计
func StatisticDaliy(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsSalesStatisticModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_statistic", "StatisticDaliy", req)
}

//历史月份统计
func StatisticMonth(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsSalesStatisticModel{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_statistic", "StatisticMonth", req)
}
