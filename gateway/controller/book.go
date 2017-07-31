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
