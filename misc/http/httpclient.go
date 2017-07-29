/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/08/02 11:05
 */

package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/goushuyun/weixin-golang/errs"
	"github.com/wothing/log"
)

var client = &http.Client{}

func GET(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errs.NewRpcError(errs.ErrUnreachable, "status code %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
	}

	log.Debugf("The response data is : %s", data)

	return data, nil
}

func GETWithUnmarshal(url string, resp interface{}) error {
	data, err := GET(url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, resp); err != nil {
		return errs.NewRpcError(errs.ErrInternal, err.Error())
	}
	return nil
}

func POST(url string, data []byte) ([]byte, error) {
	resp, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errs.NewRpcError(errs.ErrUnreachable, "status code %d", resp.StatusCode)
	}

	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
	}

	return data, nil
}

func POSTWithUnmarshal(url string, req interface{}, resp interface{}) error {
	var data []byte
	var err error

	switch req.(type) {
	case []byte:
		data = req.([]byte)
	default:
		data, err = json.Marshal(req)
		if err != nil {
			return errs.NewRpcError(errs.ErrInternal, err.Error())
		}
	}

	data, err = POST(url, data)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return errs.NewRpcError(errs.ErrInternal, err.Error())
	}
	return nil
}
