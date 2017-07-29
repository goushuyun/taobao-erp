// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package misc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/wothing/log"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc/hack"
	"github.com/goushuyun/weixin-golang/misc/jsonpb"
)

var jbm = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}

func respondBytes(rw http.ResponseWriter, r *http.Request, data []byte) {
	tid := r.Context().Value("tid").(string)
	l := len(data)
	if l > 10000 {
		// If data > 10k, only log out 10k chars
		l = 10000
	}
	if bytes.HasPrefix(data, []byte(`{"code":"00000"`)) {
		log.Tinfof(tid, "responding, response=%s", hack.String(data[:l]))
	} else {
		log.Terrorf(tid, "responding, response=%s", hack.String(data[:l]))
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

func respondObject(rw http.ResponseWriter, r *http.Request, obj interface{}) {
	message, err := json.Marshal(obj)
	if err != nil {
		err = errs.NewError(errs.ErrInternal, err.Error())
		respondObject(rw, r, err)
		return
	}
	respondBytes(rw, r, message)
}

func RespondMessage(rw http.ResponseWriter, r *http.Request, message interface{}) {
	if message == nil {
		respondObject(rw, r, map[string]interface{}{"code": errs.Ok})
		return
	}

	switch m := message.(type) {
	case proto.Message:
		s, err := jbm.MarshalToString(m)
		if err != nil {
			err = errs.NewError(errs.ErrInternal, err.Error())
			respondObject(rw, r, err)
			return
		}
		// check code exist
		// if exist it previous version, do nothing
		// if not exist
		// check if Data exist
		rpv := reflect.Indirect(reflect.ValueOf(m))
		if !rpv.FieldByName("Code").IsValid() {
			if rpv.FieldByName("Data").IsValid() {
				s = `{"code":"00000",` + s[1:]
			} else {
				s = `{"code":"00000","data":` + s + `}`
			}
		}
		respondBytes(rw, r, hack.Slice(s))
	case string:
		respondBytes(rw, r, hack.Slice(m))
	default:
		respondObject(rw, r, message)
	}
}
