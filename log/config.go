// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap/zapcore"
    "io"
    "strings"
    "time"
)

type Config struct {
    Json        bool   `json:"json"`
    Stdout      bool   `json:"stdout"`
    Application string `json:"application"`
}

func (cfg *Config) ToZapCoreEncoderConfig() zapcore.EncoderConfig {
    var prefix, suffix string
    if !cfg.Json {
        prefix = "["
        suffix = "]"
    }

    return zapcore.EncoderConfig{
        CallerKey:     "caller",
        LevelKey:      "level",
        MessageKey:    "message",
        TimeKey:       "@timestamp",
        StacktraceKey: "stacktrace",
        NameKey:       "namespace",
        LineEnding:    zapcore.DefaultLineEnding,
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + t.Format("2006-01-02 15:04:05.000") + suffix)
        },
        EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + level.CapitalString() + suffix)
        },
        EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + caller.TrimmedPath() + suffix)
        },
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeName: func(s string, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + strings.ToUpper(s) + suffix)
        },
        NewReflectedEncoder: func(w io.Writer) zapcore.ReflectedEncoder {
            enc := jsoniter.NewEncoder(w)
            enc.SetEscapeHTML(false)
            return enc
        },
        // ConsoleSeparator: " ",
    }
}

func (cfg *Config) Encoder() zapcore.Encoder {
    var (
        config = cfg.ToZapCoreEncoderConfig()
        enc    zapcore.Encoder
    )
    if cfg.Json {
        enc = zapcore.NewJSONEncoder(config)
    } else {
        enc = zapcore.NewConsoleEncoder(config)
    }
    return WrapEncoder(cfg, enc)
}
