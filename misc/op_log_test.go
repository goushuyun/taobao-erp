/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/26 18:32
 */

package misc

import (
	"fmt"
	"testing"
)

func TestNewLog(t *testing.T) {
	l1 := NewOpLog(OP_INSERT, "elvizlai", nil)
	fmt.Println(string(l1))

	type x struct {
		Name string
		Age  int64
	}

	xx := x{Name: "sdrlzyz", Age: 25}

	l2 := NewOpLog(OP_DELETE, "elvizlai", &xx)

	fmt.Println(string(l2))
}
