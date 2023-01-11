// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/auroraride/adapter/pkg/loki"
    "testing"
)

func TestLoki(t *testing.T) {
    rf := loki.GetCaller()
    t.Log(rf.File)
}
