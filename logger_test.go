// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-12
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/auroraride/adapter/loki"
    "testing"
)

func TestLogger(t *testing.T) {
    loki.SetJob("testjob")

    for i := 0; i < 3; i += 1 {
        loki.Infof("test loki: %v", i)
    }

    loki.Wait()
}
