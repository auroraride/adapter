// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-20
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "fmt"
    "go.elastic.co/ecszap"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "io"
    "os"
)

type Logger struct {
    *zap.Logger
}

func New(application string, writer io.Writer, stdout bool) {
    var cores []zapcore.Core

    encoder := ecszap.NewDefaultEncoderConfig()
    encoder.EncodeCaller = ecszap.FullCallerEncoder
    encoder.EncodeCaller = func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
        if e, ok := enc.(zapcore.ArrayEncoder); ok {
            _ = e.AppendObject(&caller{c})
        }
    }

    level := zap.DebugLevel

    if stdout {
        cores = append(
            cores,
            ecszap.NewCore(encoder, os.Stdout, level),
        )
    }

    cores = append(cores, ecszap.NewCore(encoder, zapcore.AddSync(writer), level))

    core := zapcore.NewTee(cores...)
    std = &Logger{
        Logger: zap.New(core).
            WithOptions(zap.AddCaller(), zap.AddCallerSkip(2)).
            With(
                zap.String(Application, application),
            ),
    }
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
    fields = append(fields, zap.Stack("stack"))
    l.Logger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
    fields = append(fields, zap.Stack("stack"))
    l.Logger.Panic(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zap.Field) {
    fields = append(fields, zap.Stack("stack"))
    l.Logger.DPanic(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
    l.Logger.Error(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
    l.Logger.Warn(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
    l.Logger.Info(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
    l.Logger.Debug(msg, fields...)
}

func (l *Logger) parseParams(params ...any) (args []any, fields []zap.Field) {
    for _, arg := range params {
        switch v := arg.(type) {
        case zap.Field:
            fields = append(fields, v)
        default:
            args = append(args, v)
        }
    }
    return
}

func (l *Logger) Infof(format string, params ...any) {
    args, fields := l.parseParams(params...)
    l.Logger.Info(fmt.Sprintf(format, args...), fields...)
}

func (l *Logger) Errorf(format string, params ...any) {
    args, fields := l.parseParams(params...)
    l.Logger.Error(fmt.Sprintf(format, args...), fields...)
}

func (l *Logger) GetLogger() *zap.Logger {
    return l.Logger
}
