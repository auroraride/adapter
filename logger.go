// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-12
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "go.uber.org/zap"

type Logger interface {
    Trace(args ...interface{})
    Debug(args ...interface{})
    Info(args ...any)
    Warn(args ...interface{})
    Error(args ...any)
    Fatal(args ...interface{})

    Tracef(format string, args ...interface{})
    Debugf(format string, args ...interface{})
    Infof(format string, args ...any)
    Warnf(format string, args ...interface{})
    Errorf(format string, args ...any)
    Fatalf(format string, args ...interface{})
}

type ZapLogger interface {
    Fatal(msg string, fields ...zap.Field)
    Panic(msg string, fields ...zap.Field)
    DPanic(msg string, fields ...zap.Field)
    Error(msg string, fields ...zap.Field)
    Warn(msg string, fields ...zap.Field)
    Info(msg string, fields ...zap.Field)
    Debug(msg string, fields ...zap.Field)

    Infof(format string, args ...interface{})
    Errorf(format string, args ...interface{})

    GetLogger() *zap.Logger
}
