/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/08/31 09:34
 */

package errs

import (
	"errors"
	"fmt"
	"testing"
)

func TestWoError(t *testing.T) {
	err1 := NewError(ErrAuth, "err auth")
	err2 := errors.New("self def")
	err3 := NewRpcError(ErrApiDeprecated, "rpc deprecated")
	fmt.Println(Wrap(err1))
	fmt.Println(Wrap(err2))
	fmt.Println(Wrap(err3))
}
