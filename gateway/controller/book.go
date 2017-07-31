package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

//get book info
func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "GetBookInfo", req)
}

//get book info
func SaveBook(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "SaveBook", req, "isbn", "price", "title")
}

//get book info
func UpdateBookInfo(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}
	misc.CallWithResp(w, r, "book", "UpdateBookInfo", req, "id")
}
