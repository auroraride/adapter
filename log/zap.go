// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
    grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func New(cfg *Config, options ...Option) {
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

    grpczap.ReplaceGrpcLoggerV2(logger.WithOptions(zap.WithCaller(false), zap.IncreaseLevel(zapcore.WarnLevel)).Named("grpc"))

    for _, opt := range options {
        opt.apply(logger)
    }
}
