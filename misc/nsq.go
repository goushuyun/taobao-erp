/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/26 12:40
 */

package misc

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/wothing/log"
	"golang.org/x/net/context"

	"github.com/goushuyun/weixin-golang/db"
	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc/hack"
)

type NSQCarrier struct {
	rawMsg  []byte
	Tid     string      `json:"tid"`
	Payload interface{} `json:"payload"`
}

func (carrier *NSQCarrier) encode() []byte {
	data, err := json.Marshal(carrier)
	if err != nil {
		panic(err)
	}
	return data
}

func (carrier *NSQCarrier) decode(data []byte) error {
	return json.Unmarshal(data, carrier)
}

func (carrier *NSQCarrier) NSQDecode() error {
	return carrier.decode(carrier.rawMsg)
}

// Send should init using NewProducer
func NSQPublish(topic string, ctx context.Context, v interface{}) error {
	if err := db.NSQPub.Publish(topic, (&NSQCarrier{Tid: GetTidFromContext(ctx), Payload: v}).encode()); err != nil {
		return errs.NewError(errs.ErrInternal, err.Error())
	} else {
		return nil
	}
}

func NSQDeferredPublish(topic string, delay time.Duration, ctx context.Context, v interface{}) error {
	if err := db.NSQPub.DeferredPublish(topic, delay, (&NSQCarrier{Tid: GetTidFromContext(ctx), Payload: v}).encode()); err != nil {
		return errs.NewError(errs.ErrInternal, err.Error())
	} else {
		return nil
	}
}

type HandleFunc func(c *NSQCarrier) error

const _MAX_MSG_ATTEMPTS = 12
const _BACKOFF_TIME = time.Second * 15

func NSQConsume(topic string, channel string, hf HandleFunc) {
	conn, err := nsq.NewConsumer(topic, channel, db.NSQSub.GetConfig())
	if err != nil {
		log.Fatal("NSQConsume", err)
	}

	conn.SetLogger(db.NSQLogger, nsq.LogLevelError)

	msgHandlerFunc := func() nsq.HandlerFunc {
		return func(msg *nsq.Message) error {
			c := &NSQCarrier{rawMsg: msg.Body}
			if err = hf(c); err != nil {
				if msg.Attempts > _MAX_MSG_ATTEMPTS {
					msg.Finish()
					log.Terrorf(c.Tid, "max attempted reached, give up, topic=%s, channel=%s, req=%s", topic, channel, hack.String(msg.Body))
				} else {
					msg.RequeueWithoutBackoff(_BACKOFF_TIME)
					log.Debug("handle error:", err)
				}
			} else {
				msg.Finish()
				log.Tinfof(c.Tid, "topic=%s, channel=%s, req=%s", topic, channel, SuperPrint(c.Payload))
			}

			return nil
		}
	}

	conn.AddConcurrentHandlers(msgHandlerFunc(), 4)

	err = conn.ConnectToNSQLookupd(db.NSQSub.GetAddr())
	if err != nil {
		log.Error("connect to nsqlookupd error:", err)
	}

	db.NSQSub.AddConn(conn)
}

// Consume using reflect
// handler first param MUST be a ptr
// handler MUST return err to indicate this msg can be finished
// it will be PANIC if handler not as expected
func NSQConsumeR(topic string, channel string, handler interface{}) {
	conn, err := nsq.NewConsumer(topic, channel, db.NSQSub.GetConfig())
	if err != nil {
		log.Fatal("NSQConsume", err)
	}

	conn.SetLogger(db.NSQLogger, nsq.LogLevelError)

	ht := reflect.TypeOf(handler)
	at := ht.In(0).Elem()
	hv := reflect.ValueOf(handler)

	h := func(msg *nsq.Message) error {
		v := reflect.New(at)
		c := &NSQCarrier{Payload: v.Interface()}
		err = c.decode(msg.Body)
		if err != nil {
			// this may never happened
			log.Errorf("decode error, topic=%s, channel=%s, msg body=%s", topic, channel, hack.String(msg.Body))
		}

		resp := hv.Call([]reflect.Value{v})

		if err := resp[0]; !err.IsNil() {
			if msg.Attempts > _MAX_MSG_ATTEMPTS {
				msg.Finish()
				log.Terrorf(c.Tid, "max attempted reached, give up, topic=%s, channel=%s, req=%s", topic, channel, hack.String(msg.Body))
			} else {
				msg.RequeueWithoutBackoff(_BACKOFF_TIME)
				log.Debug("handle error:", err)
			}
		} else {
			msg.Finish()
			log.Tinfof(c.Tid, "topic=%s, channel=%s, req=%s", topic, channel, SuperPrint(c.Payload))
		}

		return nil
	}

	conn.AddConcurrentHandlers(nsq.HandlerFunc(h), 4)

	err = conn.ConnectToNSQLookupd(db.NSQSub.GetAddr())
	if err != nil {
		log.Error("connect to nsqlookupd error:", err)
	}

	db.NSQSub.AddConn(conn)
}
