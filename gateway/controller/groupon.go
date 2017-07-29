package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//SharedMajorBatchSave 通用专业批量增加
func SharedMajorBatchSave(w http.ResponseWriter, r *http.Request) {
	req := &pb.SharedMajor{}
	// call RPC to handle request
	misc.CallWithResp(w, r, "bc_groupon", "SharedMajorBatchSave", req)
}

//SharedMajorList 获取专业列表（筛选获取）
func SharedMajorList(w http.ResponseWriter, r *http.Request) {
	req := &pb.SharedMajor{}
	misc.CallWithResp(w, r, "bc_groupon", "SharedMajorList", req)
}

//创建学校的学院
func SaveSchoolInstitute(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.SchoolInstitute{}
	misc.CallWithResp(w, r, "bc_groupon", "SaveSchoolInstitute", req, "school_id", "name")
}

//创建学院专业
func SaveInstituteMajor(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.InstituteMajor{}
	misc.CallWithResp(w, r, "bc_groupon", "SaveInstituteMajor", req, "institute_id", "name")
}

//获取学校学院专业列表
func GetSchoolMajorInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.SchoolMajorInfoReq{StoreId: c.StoreId, UserType: 1}
	misc.CallWithResp(w, r, "bc_groupon", "GetSchoolMajorInfo", req)
}

//获取学校学院专业列表
func GetSchoolMajorInfoApp(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.SchoolMajorInfoReq{StoreId: c.StoreId, UserType: 0}
	misc.CallWithResp(w, r, "bc_groupon", "GetSchoolMajorInfo", req)
}

//创建班级购
func SaveGroupon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{StoreId: c.StoreId, FounderId: c.SellerId, FounderType: 2}
	misc.CallWithResp(w, r, "bc_groupon", "SaveGroupon", req, "term", "school_id", "institute_id", "institute_major_id", "founder_id", "class", "founder_name", "founder_mobile", "expire_at", "items")
}

//创建班级购
func SaveGrouponApp(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{FounderId: c.UserId, FounderType: 1, StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "SaveGroupon", req, "term", "school_id", "institute_id", "institute_major_id", "founder_id", "class", "founder_name", "founder_mobile", "expire_at", "items")
}

//班级购列表
func GrouponList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "GrouponList", req)
}

//班级购列表
func GrouponListApp(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "GrouponList", req)
}

//我的班级购
func MyGroupon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{FounderId: c.SellerId, FounderType: 2}
	misc.CallWithResp(w, r, "bc_groupon", "MyGroupon", req)
}

//我的班级购
func MyGrouponApp(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{FounderId: c.UserId, FounderType: 1, StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "MyGroupon", req)
}

//新增班级购项
func GetGrouponItems(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{}
	misc.CallWithResp(w, r, "bc_groupon", "GetGrouponItems", req, "id")
}

//获取班级购参与人信息
func GetGrouponPurchaseUsers(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{}
	misc.CallWithResp(w, r, "bc_groupon", "GetGrouponPurchaseUsers", req, "id")
}

//获取班级购操作日志
func GetGrouponOperateLog(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{}
	misc.CallWithResp(w, r, "bc_groupon", "GetGrouponOperateLog", req, "id")
}

//修改班级购
func UpdateGruopon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{}
	misc.CallWithResp(w, r, "bc_groupon", "UpdateGruopon", req, "id")
}

//批量班级购日期
func BatchUpdateGrouponExpireAt(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Groupon{StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "BatchUpdateGrouponExpireAt", req, "expire_at", "update_ids")
}

func StarGroupon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GrouponOperateLog{FounderId: c.UserId, FounderType: 1, OperateType: "star"}
	misc.CallWithResp(w, r, "bc_groupon", "StarGroupon", req, "groupon_id", "founder_name")
}

func ShareGroupon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GrouponOperateLog{FounderId: c.UserId, FounderType: 1, OperateType: "share"}
	misc.CallWithResp(w, r, "bc_groupon", "ShareGroupon", req, "groupon_id", "founder_name")
}

//下单
func GrouponSubmit(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.GrouponSubmitModel{UserId: c.UserId, StoreId: c.StoreId}

	misc.CallWithResp(w, r, "bc_groupon", "GrouponSubmit", req, "mobile", "name", "address", "groupon_id")
}

//保存学生学籍信息
func SaveUserSchoolStatus(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.UserSchoolStatus{UserId: c.UserId, StoreId: c.StoreId}

	misc.CallWithResp(w, r, "bc_groupon", "SaveUserSchoolStatus", req, "school_id", "institute_id", "institute_major_id")
}

//更新学生学籍信息
func UpdateUserSchoolStatus(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.UserSchoolStatus{UserId: c.UserId, StoreId: c.StoreId}

	misc.CallWithResp(w, r, "bc_groupon", "UpdateUserSchoolStatus", req, "id")
}

//获取学生学籍
func GetUserSchoolStatus(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.UserSchoolStatus{UserId: c.UserId, StoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_groupon", "GetUserSchoolStatus", req)
}

//删除专业
func DelInstituMajor(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.InstituteMajor{}
	misc.CallWithResp(w, r, "bc_groupon", "DelInstituMajor", req, "id")
}

//修改学校专业名称
func UpdateInstituteMajor(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.InstituteMajor{}
	misc.CallWithResp(w, r, "bc_groupon", "UpdateInstituteMajor", req, "id", "name")
}

//删除学院
func DelSchoolInstitute(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.SchoolInstitute{}
	misc.CallWithResp(w, r, "bc_groupon", "DelSchoolInstitute", req, "id")
}

//修改学校学院名称
func UpdateSchoolInstitute(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.SellerId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.SchoolInstitute{}
	misc.CallWithResp(w, r, "bc_groupon", "UpdateSchoolInstitute", req, "id", "name")
}

//用户点赞记录
func HasStarGroupon(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)

	if c == nil || c.StoreId == "" || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GrouponOperateLog{FounderId: c.UserId}
	misc.CallWithResp(w, r, "bc_groupon", "HasStarGroupon", req, "groupon_id")
}
