// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-03
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
    "go.uber.org/zap/zapcore"
    "google.golang.org/grpc/codes"
)

var (
    GrpcZapDefaultOptions = []grpczap.Option{
        grpczap.WithLevels(func(c codes.Code) zapcore.Level {
            if c == codes.Unauthenticated {
                // Make this a special case for tests, and an error.
                return zapcore.ErrorLevel
            }
            level := grpczap.DefaultClientCodeToLevel(c)
            return level
        }),
    }
)
