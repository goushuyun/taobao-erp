package controller

import (
	"encoding/json"
	"net/http"

	"github.com/goushuyun/weixin-golang/errs"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
	"github.com/wothing/log"
)

type payload struct {
	Data Data `json:"data"`
}
type Data struct {
	Object *pb.PaySuccessCallbackPayload `json:"object"`
}

func RefundSuccessNotify(w http.ResponseWriter, r *http.Request) {
	log.Debugf("The req is : %s\v", r.Context().Value("body"))

	// callback string
	callback, ok := r.Context().Value("body").([]byte)
	if !ok {
		log.Error("interface to string error")
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "json unmarshal error",
		})
		return
	}

	// callback struct
	p := &pb.PaySuccessCallbackPayload{}
	data := Data{Object: p}
	obj := payload{
		Data: data,
	}

	// unmarshal
	err := json.Unmarshal(callback, &obj)
	if err != nil {
		log.Error(err)
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "json unmarshal error",
		})
		return
	}

	log.JSONIndent(p)

	// 组织数据，提送退款成功消息
	refundSuccessReq := &pb.AfterSaleModel{
		RefundFee:     p.Amount,
		RefundTradeNo: p.TransactionNo, //退款单号
		IsSuccess:     p.Succeed,
		TradeNo:       p.Charge,
	}

	misc.CallWithResp(w, r, "bc_order", "AfterSaleOrderHandledResult", refundSuccessReq)
}

func PaySuccessNotify(w http.ResponseWriter, r *http.Request) {
	log.Debugf("The response is : %s\n", r.Context().Value("body"))

	// callback string
	callback, ok := r.Context().Value("body").([]byte)
	if !ok {
		log.Error("interface to string error")
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "json unmarshal error",
		})
		return
	}

	// callback struct
	p := &pb.PaySuccessCallbackPayload{}
	data := Data{Object: p}
	obj := payload{
		Data: data,
	}

	// unmarshal
	err := json.Unmarshal(callback, &obj)
	if err != nil {
		log.Error(err)
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "json unmarshal error",
		})
		return
	}

	log.Debugf("The callback obj is %+v\n", obj)

	if p.Metadata["event"] == "recharge" {
		// recharge success
		recharge_model := &pb.RechargeModel{
			Id:          p.OrderNo,
			RechargeFee: p.Amount,
			CompleteAt:  p.TimePaid,
			PayWay:      "alipay_pc_direct",
			TradeNo:     p.TransactionNo,
			ChargeId:    p.Id,
		}

		log.JSON(recharge_model)

		_, err = misc.CallRPC(misc.GenContext(r), "bc_store", "RechargeHandler", recharge_model)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.Ok,
			"message": "ok",
		})

	} else if p.Metadata["event"] == "buy_goods" {

		// 封装支付成功请求对象
		order := &pb.Order{
			Id:         p.OrderNo,
			TradeNo:    p.Id,
			PayChannel: p.Channel,
		}

		log.Debugf("The order obj is %+v\n", order)
		_, err = misc.CallRPC(misc.GenContext(r), "bc_order", "PaySuccess", order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.Ok,
			"message": "ok",
		})

	}

}

func GetCharge(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	req := &pb.GetChargeReq{Ip: ip}
	misc.CallWithResp(w, r, "bc_payment", "GetCharge", req, "channel", "order_no", "amount", "subject", "body")
}
