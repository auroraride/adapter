// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package snag

import (
    "fmt"
    "github.com/auroraride/adapter"
    "go.uber.org/zap"
)

func WithRecover(cb func(), logger adapter.ZapLogger) {

    defer func() {
        if v := recover(); v != nil {
            // logger.Errorf("捕获未处理崩溃: %v\n%s", v, debug.Stack())
            logger.Panic(
                "捕获未处理崩溃",
                zap.Stack("stack"),
                zap.Error(fmt.Errorf("%v", v)),
            )
        }
    }()

    cb()
}
