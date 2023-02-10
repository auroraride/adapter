// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
    "go.uber.org/zap/zapcore"
)

func NewCore(cfg *Config, ws zapcore.WriteSyncer, enab zapcore.LevelEnabler) zapcore.Core {
    return WrapCore(cfg, zapcore.NewCore(cfg.Encoder(), ws, enab))
}

func WrapCore(cfg *Config, c zapcore.Core) zapcore.Core {
    return &core{Core: c, config: cfg}
}

type core struct {
    zapcore.Core

    config *Config
}

func (c core) With(fields []zapcore.Field) zapcore.Core {
    return &core{Core: c.Core.With(fields), config: c.config}
}

func (c *core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
    if ent.LoggerName == "" {
        ent.LoggerName = c.config.Application
    } else {
        ent.LoggerName = ent.LoggerName + "-" + c.config.Application
    }

    if c.Enabled(ent.Level) {
        return ce.AddCore(ent, c)
    }
    return ce
}

func (c core) Write(ent zapcore.Entry, fields []zapcore.Field) error {
    return c.Core.Write(ent, fields)
}
