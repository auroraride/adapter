// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
)

type encoder struct {
    zapcore.Encoder

    config *Config
}

func WrapEncoder(cfg *Config, enc zapcore.Encoder) zapcore.Encoder {
    return &encoder{
        Encoder: enc,
        config:  cfg,
    }
}

func (e *encoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (b *buffer.Buffer, err error) {
    b, err = e.Encoder.EncodeEntry(ent, fields)
    return
}

type loggerObject struct {
    namespace   string
    application string
}

func (o *loggerObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
    enc.AddString("application", o.application)
    enc.AddString("namespace", o.namespace)
    return nil
}

type object struct {
    key   string
    value string
}

func (o *object) MarshalLogObject(enc zapcore.ObjectEncoder) error {
    enc.AddString(o.key, o.value)
    return nil
}
