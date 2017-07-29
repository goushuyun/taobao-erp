package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//AddSchool 增加学校
func AddSchool(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_school", "AddSchool", req, "name", "tel", "lat", "lng")
}

//UpdateSchool 更改学校基本信息
func UpdateSchool(w http.ResponseWriter, r *http.Request) {

	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_school", "UpdateSchool", req, "id", "name", "tel", "lat", "lng")
}

//UpdateSchool 更改学校基本信息
func UpdateExpressFee(w http.ResponseWriter, r *http.Request) {
	//检测token
	c := token.Get(r)
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId, Seller: &pb.SellerInfo{Id: c.SellerId, Mobile: c.Mobile}}
	misc.CallWithResp(w, r, "bc_school", "UpdateExpressFee", req, "id")
}

//StoreSchools 店铺下的所有学校
func StoreSchools(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_school", "StoreSchools", req)
}

//StoreSchools 店铺下的所有学校
func StoreSchoolsApp(w http.ResponseWriter, r *http.Request) {
	req := &pb.School{}
	misc.CallWithResp(w, r, "bc_school", "StoreSchools", req)
}

//GetSchoolById 根据学校id获取学校信息
func GetSchoolById(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_school", "GetSchoolById", req)
}

//GetSchoolById 根据学校id获取学校信息
func UpdateSchoolRecylingState(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_school", "UpdateSchoolRecylingState", req, "school_ids")
}

//DelSchool 删除学校
func DelSchool(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.School{StoreId: c.StoreId, DelStaffId: c.SellerId}
	misc.CallWithResp(w, r, "bc_school", "DelSchool", req, "id")
}
