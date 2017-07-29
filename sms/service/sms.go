package service

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"

	"github.com/wothing/log"
	"golang.org/x/net/context"
)

var (
	ErrSmsFrequent = errors.New("ErrFrequert")
)

type SMSServer struct{}

func (s *SMSServer) SendSMS(ctx context.Context, req *pb.SMSReq) (resp *pb.Void, err error) {

	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SendSMS", "%#v", req))

	if strings.HasPrefix(req.Mobile, "86-111") {
		return nil, errs.NewError(errs.ErrInternal, "111 开头测试帐号")
	}

	id := templateId[req.Type]
	send, err := s.GenSMSJson(id, req.Mobile, req.Message)
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, err.Error())
	}
	err = s.SendMessage(send)
	if err == ErrSmsFrequent {
		return nil, errs.NewError(errs.ErrInternal, err.Error())
	}
	if err != nil {
		log.Terrorf(tid, "err:%s", err.Error())
		return nil, errs.NewError(errs.ErrInternal, err.Error())
	}
	return &pb.Void{}, nil
}

func (*SMSServer) SendMessage(send string) error {
	ts := time.Now().Format("20060102150405")
	signature := accountsid + accounttoken + ts
	checksum := md5.Sum([]byte(signature))
	sig := strings.ToUpper(hex.EncodeToString(checksum[:]))
	url := url + accountsid + "/SMS/TemplateSMS?sig=" + sig
	auth := base64.StdEncoding.EncodeToString([]byte(accountsid + ":" + ts))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(send)))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Authorization", auth)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "160038") || strings.Contains(string(body), "160039") || strings.Contains(string(body), "160040") || strings.Contains(string(body), "160041") {
		return ErrSmsFrequent
	}
	if !strings.Contains(string(body), "000000") {
		return errors.New(string(body))
	}
	resp.Body.Close()
	return nil
}

func (*SMSServer) GenSMSJson(tid string, mobile string, message []string) (jsonString string, err error) {
	var ms []string
	for _, m := range strings.Split(mobile, ",") {
		str, err := misc.MobileFormat(m)
		if err != nil {
			return "", err
		}
		ms = append(ms, str)
	}
	if len(message) == 0 {
		message = []string{}
		// return "", errors.New("message is no content")
	}

	s := &SMSTemplate{To: strings.Join(ms, ","), AppId: appid, TemplateId: tid, Datas: message}
	send, _ := json.Marshal(s)
	log.Infof("send:%s", string(send))
	return string(send), nil
}
