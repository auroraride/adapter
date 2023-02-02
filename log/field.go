// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import "go.uber.org/zap"

func Binary(b []byte) zap.Field {
    return zap.Binary("binary", b)
}

func Payload(payload any) zap.Field {
    return zap.Reflect("payload", payload)
}
