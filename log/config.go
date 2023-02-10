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
    FormatJson bool
    Stdout     bool
    LoggerName string
    Writers    []io.Writer
}

func (cfg *Config) ToZapCoreEncoderConfig() zapcore.EncoderConfig {
    var prefix, suffix string
    if !cfg.FormatJson {
        prefix = "["
        suffix = "]"
    }

    return zapcore.EncoderConfig{
        CallerKey:     "caller",
        LevelKey:      "level",
        MessageKey:    "message",
        TimeKey:       "ts",
        StacktraceKey: "stacktrace",
        NameKey:       "logger",
        LineEnding:    zapcore.DefaultLineEnding,
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + t.Format("2006-01-02T15:04:05.000Z0700") + suffix)
        },
        EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + level.CapitalString() + suffix)
        },
        EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(prefix + caller.TrimmedPath() + suffix)
        },
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeName: func(s string, enc zapcore.PrimitiveArrayEncoder) {
            // if cfg.FormatJson {
            //     if e, ok := enc.(zapcore.ArrayEncoder); ok {
            //         _ = e.AppendObject(cfg.parseName(s))
            //     }
            //     return
            // }
            enc.AppendString(prefix + strings.ToLower(s) + suffix)
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
    if cfg.FormatJson {
        enc = zapcore.NewJSONEncoder(config)
    } else {
        enc = zapcore.NewConsoleEncoder(config)
    }
    return WrapEncoder(cfg, enc)
}
