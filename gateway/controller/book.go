package controller

import (
	"fmt"
	"net/http"

	"github.com/goushuyun/taobao-erp/errs"

	"github.com/goushuyun/taobao-erp/misc/token"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/users/role"
)

//get book info
func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "GetBookInfo", req)
}

//get book info from  local
func GetLocalBookInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "GetLocalBookInfo", req)
}

//get book info
func SaveBook(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	if c.Role != role.InterAdmin {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "权限不够",
		})
		return
	}
	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "SaveBook", req, "isbn", "price", "title")
}

//get book info
func UpdateBookInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	if c.Role != role.InterAdmin {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "权限不够",
		})
		return
	}
	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "UpdateBookInfo", req, "id")
}

//get book info
func SubmitBookAudit(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.BookAuditRecord{ApplyUserId: c.UserId}
	misc.CallWithResp(w, r, "book", "SubmitBookAudit", req)
}

//get book info
func GetBookAuditRecord(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.BookAuditRecord{}
	misc.CallWithResp(w, r, "book", "GetBookAuditRecord", req)
}

//get book info
func GetOrganizedBookAuditList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.BookAuditRecord{}
	misc.CallWithResp(w, r, "book", "GetOrganizedBookAuditList", req)
}

//handle the book audit
func HandleBookAudit(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	fmt.Println(c.Role)
	fmt.Println(role.InterAdmin)
	if c.Role != role.InterAdmin {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "权限不够",
		})
		return
	}
	req := &pb.BookAuditRecord{CheckUserId: c.UserId}
	misc.CallWithResp(w, r, "book", "HandleBookAudit", req)
}
