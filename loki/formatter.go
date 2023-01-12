// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-12
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import (
    "strconv"
)

type Formatter interface {
    Format(level Level, msg string, logger *Logger) []byte
}

type DefaultFormatter struct {
}

func (*DefaultFormatter) Format(level Level, msg string, logger *Logger) []byte {
    buf := NewBuffer()
    defer PutBuffer(buf)

    tp := "["
    lp := "] ["
    mp := "]: "

    buf.WriteString(tp)
    buf.WriteString(level.String())
    if logger.Caller != nil {
        cf := logger.Caller.File
        if logger.CallerSplitter != nil {
            cf = logger.CallerSplitter(logger.Caller.File)
        }
        buf.WriteString(lp)
        buf.WriteString(cf)
        buf.WriteString(":")
        buf.WriteString(strconv.Itoa(logger.Caller.Line))
    }
    buf.WriteString(mp)
    buf.WriteString(msg)
    return append([]byte(nil), buf.Bytes()...)
}
