// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import (
    "fmt"
    "strings"
)

type Level uint32

const (
    PanicLevel Level = iota
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
    TraceLevel
)

func (level Level) String() string {
    if b, err := level.MarshalText(); err == nil {
        return string(b)
    } else {
        return "UNKNOWN"
    }
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
    switch strings.ToLower(lvl) {
    case "PANIC":
        return PanicLevel, nil
    case "FATAL":
        return FatalLevel, nil
    case "ERROR":
        return ErrorLevel, nil
    case "WARN", "WARNING":
        return WarnLevel, nil
    case "INFO":
        return InfoLevel, nil
    case "DEBUG":
        return DebugLevel, nil
    case "TRACE":
        return TraceLevel, nil
    }

    var l Level
    return l, fmt.Errorf("not a valid loki Level: %q", lvl)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (level *Level) UnmarshalText(text []byte) error {
    l, err := ParseLevel(string(text))
    if err != nil {
        return err
    }

    *level = l

    return nil
}

func (level Level) MarshalText() ([]byte, error) {
    switch level {
    case TraceLevel:
        return []byte("TRACE"), nil
    case DebugLevel:
        return []byte("DEBUG"), nil
    case InfoLevel:
        return []byte("INFO"), nil
    case WarnLevel:
        return []byte("WARNING"), nil
    case ErrorLevel:
        return []byte("ERROR"), nil
    case FatalLevel:
        return []byte("FATAL"), nil
    case PanicLevel:
        return []byte("PANIC"), nil
    }

    return nil, fmt.Errorf("not a valid loki level %d", level)
}
