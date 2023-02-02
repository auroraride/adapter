// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package snag

import (
    "fmt"
    "github.com/auroraride/adapter/zlog"
    "go.uber.org/zap"
)

func WithRecover(cb func()) {

    defer func() {
        if v := recover(); v != nil {
            zlog.Error(
                "捕获未处理崩溃",
                zap.Stack("stack"),
                zap.Error(fmt.Errorf("%v", v)),
            )
        }
    }()

    cb()
}
