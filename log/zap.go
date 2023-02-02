// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func Test() {
    zap.L().Named("Test").Info("Test", zap.String("str", "v"))
}

func Initialize(cfg *Config) {
    var opts []zapcore.WriteSyncer
    if cfg.Stdout {
        opts = append(opts, zapcore.AddSync(os.Stdout))
    }
    if cfg.Application == "" {
        panic("application必填")
    }

    syncWriter := zapcore.NewMultiWriteSyncer(opts...)
    c := NewCore(
        cfg,
        syncWriter,
        zap.NewAtomicLevelAt(zap.DebugLevel),
    )
    logger := zap.New(c, zap.AddCaller())

    zap.ReplaceGlobals(logger)
}
