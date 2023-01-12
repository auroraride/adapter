// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-12
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/auroraride/adapter/pkg/loki"
    "testing"
)

func TestLogger(t *testing.T) {
    loki.SetJob("testjob")
    loki.Info("wait test go")

    loki.Wait()
}
