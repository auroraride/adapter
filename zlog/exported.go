// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-20
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "go.uber.org/zap"
)

const Application = "application"

var (
    std *Logger
)

func StandardLogger() *Logger {
    return std
}

func Sync() {
    _ = std.Logger.Sync()
}

func Fatal(msg string, fields ...zap.Field) {
    std.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
    std.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
    std.DPanic(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    std.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    std.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
    std.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
    std.Debug(msg, fields...)
}

func Infof(format string, args ...any) {
    std.Infof(format, args...)
}

func Errorf(format string, args ...any) {
    std.Errorf(format, args...)
}

func (*Logger) clone() *Logger {
    cp := *(std.Logger)
    return &Logger{Logger: &cp}
}

func Named(name string) *Logger {
    // cp := std.clone()
    // cp.Logger = &Logger{Logger: }

    return &Logger{Logger: std.Logger.Named(name).WithOptions(zap.AddCallerSkip(-1))}
}
