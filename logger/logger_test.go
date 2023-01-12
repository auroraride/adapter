// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package logger

import (
    log "github.com/sirupsen/logrus"
    "testing"
)

func TestLogger(t *testing.T) {

    // 日志
    LoadWithConfig(Config{
        Color:  true,
        Level:  "info",
        Age:    8192,
        Caller: true,
        Path:   "/tmp/logrus",
    })

    log.Info("test")
}
