package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

func SendSms(w http.ResponseWriter, r *http.Request) {
	req := &pb.User{}

	misc.CallWithResp(w, r, "sms", "SendSMS", req, "type", "mobile", "message")
}
