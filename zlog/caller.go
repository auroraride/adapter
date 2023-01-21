// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-21
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "go.uber.org/zap/zapcore"
    "strconv"
)

type caller struct {
    zapcore.EntryCaller
}

func (c *caller) MarshalLogObject(enc zapcore.ObjectEncoder) error {
    enc.AddString("caller", c.File+":"+strconv.Itoa(c.Line))
    return nil
}
