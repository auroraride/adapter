// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package snag

import (
    "github.com/auroraride/adapter"
    "runtime/debug"
)

func WithRecover(cb func(), logger adapter.Logger) {

    defer func() {
        if v := recover(); v != nil {
            logger.Errorf("捕获未处理崩溃: %v\n%s", v, debug.Stack())
        }
    }()

    cb()
}
