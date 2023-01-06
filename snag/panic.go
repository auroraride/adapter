// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package snag

import (
    "github.com/sirupsen/logrus"
    "runtime/debug"
)

func WithPanic(cb func(), logger logrus.FieldLogger) {

    defer func() {
        if v := recover(); v != nil {
            logger.Errorf("捕获未处理崩溃: %v\n%s", v, debug.Stack())
        }
    }()

    cb()
}
