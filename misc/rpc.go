/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/06/13 09:42
 */

package misc

import (
	"context"
	"net/http"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/wothing/log"
	"github.com/wothing/worc"
	"google.golang.org/grpc/metadata"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/pb"
)

var serviceList = map[string]interface{}{
	"bc_seller":     pb.NewSellerServiceClient,
	"bc_sms":        pb.NewSMSServiceClient,
	"bc_store":      pb.NewStoreServiceClient,
	"bc_mediastore": pb.NewMediastoreClient,
	"bc_school":     pb.NewSchoolServiceClient,
	"bc_location":   pb.NewLocationServiceClient,
	"bc_books":      pb.NewBooksServiceClient,
	"bc_goods":      pb.NewGoodsServiceClient,
	"bc_topic":      pb.NewTopicServiceClient,
	"bc_weixin":     pb.NewWeixinServiceClient,
	"bc_user":       pb.NewUserServiceClient,
	"bc_circular":   pb.NewCircularServiceClient,
	"bc_cart":       pb.NewCartServiceClient,
	"bc_order":      pb.NewOrderServiceClient,
	"bc_account":    pb.NewAccountServiceClient,
	"bc_address":    pb.NewAddressServiceClient,
	"bc_payment":    pb.NewPaymentServiceClient,
	"bc_retail":     pb.NewRetailServiceClient,
	"bc_statistic":  pb.NewStatisticServiceClient,
	"bc_master":     pb.NewMasterServiceClient,
	"bc_groupon":    pb.NewGrouponServiceClient,
}

func StartServiceConns(address string, serviceNames []string) {
	worc.StartServiceConns(address, serviceNames)
}

func CloseServiceConns() {
	worc.CloseServiceConns()
}

// CallRPC
// before using this method, you should register all services by worc.StartServiceConns
// and close all conns at end of application by using worc.CloseServiceConns
// Ex:
// worc.StartServiceConns(addressOfEtcd,serviceList)
// defer worc.CloseServiceConns()
func CallRPC(ctx context.Context, serviceName string, method string, req interface{}) (interface{}, error) {
	log.CtxDebugf(ctx, "%s.%s %#v", serviceName, method, req)
	rp, err := worc.CallRPC(ctx, serviceList[serviceName], serviceName, method, req)
	if err != nil {
		log.CtxErrorf(ctx, "CallRPC '%s' '%s' error: %s", serviceName, method, err.Error())
		return nil, errs.FromRpcError(err)
	}

	rpv := reflect.Indirect(reflect.ValueOf(rp))
	// if Code field exist, it is previous version
	if rpv.FieldByName("Code").IsValid() {
		code, _ := rpv.FieldByName("Code").Interface().(string)
		message, _ := rpv.FieldByName("Message").Interface().(string)
		if code != errs.Ok {
			log.CtxErrorf(ctx, "CallRPC '%s' '%s' error: %v", serviceName, method, rpv)
			return nil, errs.NewError(code, message)
		}
	}
	return rp, nil
}

func CallSVC(ctx context.Context, serviceName string, method string, req interface{}, resp interface{}) error {
	rp, err := CallRPC(ctx, serviceName, method, req)
	if err != nil {
		return err
	}

	// set val
	respv := reflect.Indirect(reflect.ValueOf(resp))
	if respv.CanSet() {
		respv.Set(reflect.Indirect(reflect.ValueOf(rp)))
	} else {
		log.Terrorf(GetTidFromContext(ctx), "CallSVC '%s' '%s' error: %s", serviceName, method, "resp is not addressable")
		return errs.NewError(errs.ErrInternal, "resp is not addressable")
	}
	return nil
}

func CallRPCWithNewCtx(ctx context.Context, serviceName string, method string, req interface{}) (interface{}, error) {
	md, ok := metadata.FromContext(ctx)
	if ok {
		ctx = metadata.NewContext(context.Background(), md)
	} else {
		ctx = context.Background()
	}
	return CallRPC(ctx, serviceName, method, req)
}

func CallSVCWithNewCtx(ctx context.Context, serviceName string, method string, req interface{}, resp interface{}) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		ctx = metadata.NewContext(context.Background(), md)
	} else {
		ctx = context.Background()
	}
	return CallSVC(ctx, serviceName, method, req, resp)
}

func CallWithResp(rw http.ResponseWriter, r *http.Request, serviceName string, method string, req proto.Message, constraints ...string) {
	err := Request2Struct(r, req, constraints...)
	if err != nil {
		RespondMessage(rw, r, err)
		return
	}

	resp, err := CallRPC(GenContext(r), serviceName, method, req)
	if err != nil {
		RespondMessage(rw, r, err)
		return
	}

	RespondMessage(rw, r, resp)
}
