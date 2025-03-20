// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	if 1 != Or(true, 1, 2) {
		t.Fail()
	}
	if "A" != Or(false, "B", "A") {
		t.Fail()
	}

	var a *string
	if time.Now().Second()%2 == 0 {
		a = nil
	} else {
		x := "hello"
		a = &x
	}
	v := OrFunc(
		func() bool {
			return a == nil
		},
		func() string {
			return "world"
		},
		func() string {
			return *a
		},
	)
	t.Logf("v: %s", v)
}
