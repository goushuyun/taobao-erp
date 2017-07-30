package controller

import (
	"net/http"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/pb"
)

func SendIdentifyingCode(w http.ResponseWriter, r *http.Request) {
	req := &pb.SMSReq{}

	misc.CallWithResp(w, r, "sms", "SendIdentifyingCode", req, "mobile")
}
