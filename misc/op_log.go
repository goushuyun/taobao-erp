/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/26 18:24
 */

package misc

import (
	"encoding/json"
	"time"
)

// used for produce logs for data change

type opType int

const (
	_ opType = iota
	OP_INSERT
	OP_UPDATE
	OP_DELETE
)

type opLog struct {
	OpType    opType      `json:"op,omitempty"`
	Operator  string      `json:"operator,omitempty"`
	Snapshot  interface{} `json:"snapshot,omitempty"`
	OperateAt int64       `json:"operate_at,omitempty"`
}

//INSERT INTO users (logs) VALUES ('{"snapshot":null,"operator":"elvizlai","operate_at":1464259313}'::JSONB||'[]')

func NewOpLog(op opType, operator string, snapshot interface{}) []byte {
	data, _ := json.Marshal(&opLog{OpType: op, Snapshot: snapshot, Operator: operator, OperateAt: time.Now().Unix()})
	return data
}
