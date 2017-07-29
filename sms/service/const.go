package service

import "github.com/goushuyun/weixin-golang/pb"

var (
	templateId = map[pb.SMSType]string{
		pb.SMSType_CommonCheckCode:    "164093",
		pb.SMSType_Delivery:           "174196",
		pb.SMSType_AutoConfirmReceipt: "174203",
	}
)

const (
	accountsid   = "aaf98f894cfa16bc014d039ced990783"
	url          = "https://app.cloopen.com:8883/2013-12-26/Accounts/"
	appid        = "8a216da85afaadec015b1d5533390dbb"
	accounttoken = "52591dc2aedd49a8b74027f7cafcd6e6"
)

type SMSTemplate struct {
	To         string   `json:"to"`
	AppId      string   `json:"appId"`
	TemplateId string   `json:"templateId"`
	Datas      []string `json:"datas"`
}
