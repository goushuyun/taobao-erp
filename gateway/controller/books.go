package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/errs"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/misc/token"
	"github.com/goushuyun/weixin-golang/pb"
)

// insert book
func InsertBook(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}

	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_books", "InsertBook", req, "isbn", "title", "store_id", "price")
}

func ModifyBookInfo(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}
	misc.CallWithResp(w, r, "bc_books", "ModifyBookInfo", req, "id")
}

func GetBookInfoByISBN(w http.ResponseWriter, r *http.Request) {
	req := &pb.Book{}

	// get store_id
	if c := token.Get(r); c != nil {
		req.StoreId = c.StoreId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "bc_books", "GetBookInfoByISBN", req, "isbn")
}
