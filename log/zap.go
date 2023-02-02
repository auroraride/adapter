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

func Initialize(cfg *Config) {
    var opts []zapcore.WriteSyncer
    if cfg.Stdout {
        opts = append(opts, zapcore.AddSync(os.Stdout))
    }
    for _, w := range cfg.Writers {
        opts = append(opts, zapcore.AddSync(w))
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

    // SetStandardLogger(logger)
    zap.ReplaceGlobals(logger)
}
