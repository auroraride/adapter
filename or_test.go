// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "testing"

func TestOr(t *testing.T) {
	if 1 != Or(true, 1, 2) {
		t.Fail()
	}
	if "A" != Or(false, "B", "A") {
		t.Fail()
	}
}
