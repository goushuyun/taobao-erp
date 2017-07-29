/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/08/08 07:31
 */

package hack

import (
	"testing"
)

func TestString(t *testing.T) {
	a := []byte("hello world")
	b := String(a)

	if b != "hello world" {
		t.Fatal("slice to string test failed", "b:", b)
	}

	a[0] = 'f'

	if b != "fello world" {
		t.Fatal("slice to string test failed", "b:", b)
	}

	// append wont change
	a = append(a, "abc"...)
	if b != "fello world" {
		t.Fatal("slice to string test failed", "b:", b)
	}
}

func TestSlice(t *testing.T) {

}
