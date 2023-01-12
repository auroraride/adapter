// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import (
    "fmt"
    "github.com/auroraride/adapter/async"
    "github.com/valyala/fasthttp"
    "os"
    "runtime"
    "strconv"
    "sync"
    "time"
)

type Logger struct {
    job          []byte
    url          string
    reportCaller bool

    Caller         *runtime.Frame
    CallerSplitter func(string) string
    Formatter      Formatter
    WaitGroup      *sync.WaitGroup
}

func New() *Logger {
    logger := &Logger{
        job:          []byte("varjob"),
        url:          "http://localhost:3100/loki/api/v1/push",
        Formatter:    &DefaultFormatter{},
        reportCaller: true,
        WaitGroup:    &sync.WaitGroup{},
    }

    return logger
}

var (
    bodyLeft     = []byte(`{"streams": [{ "stream": { "job": "`)
    bodyMid      = []byte(`" }, "values": [ [ "`)
    bodyMidSplit = []byte(`", "`)
    bodyRight    = []byte(`" ] ] }]}`)
)

func (logger *Logger) send(body []byte) {
    defer logger.WaitGroup.Done()

    req := fasthttp.AcquireRequest()
    defer fasthttp.ReleaseRequest(req)

    req.SetRequestURI(logger.url)
    req.Header.SetContentType("application/json")
    req.Header.SetMethod("POST")
    req.SetBody(body)

    resp := fasthttp.AcquireResponse()
    defer fasthttp.ReleaseResponse(resp)

    _ = fasthttp.Do(req, resp)
}

func (logger *Logger) Log(job []byte, level Level, args ...any) {
    if len(job) == 0 {
        return
    }

    if logger.reportCaller {
        logger.Caller = getCaller()
    }

    str := fmt.Sprint(args...)

    msg := logger.Formatter.Format(level, str, logger)

    if len(msg) > 0 {
        logger.WaitGroup.Add(1)
        go async.WithTask(func() {
            buf := NewBuffer()
            defer PutBuffer(buf)

            buf.Write(bodyLeft)
            buf.Write(job)
            buf.Write(bodyMid)
            buf.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
            buf.Write(bodyMidSplit)
            buf.Write(msg)
            buf.Write(bodyRight)
            logger.send(append([]byte(nil), buf.Bytes()...))
        })
    }

    if level == FatalLevel {
        os.Exit(1)
    }
}

func (logger *Logger) Logf(job []byte, level Level, format string, args ...any) {
    logger.Log(job, level, fmt.Sprintf(format, args...))
}

func (logger *Logger) Trace(args ...any) {
    logger.Log(logger.job, TraceLevel, args...)
}

func (logger *Logger) Debug(args ...any) {
    logger.Log(logger.job, DebugLevel, args...)
}

func (logger *Logger) Info(args ...any) {
    logger.Log(logger.job, InfoLevel, args...)
}

func (logger *Logger) Warn(args ...any) {
    logger.Log(logger.job, WarnLevel, args...)
}

func (logger *Logger) Error(args ...any) {
    logger.Log(logger.job, ErrorLevel, args...)
}

func (logger *Logger) Fatal(args ...any) {
    logger.Log(logger.job, FatalLevel, args...)
}

func (logger *Logger) Tracef(format string, args ...any) {
    logger.Logf(logger.job, TraceLevel, format, args...)
}

func (logger *Logger) Debugf(format string, args ...any) {
    logger.Logf(logger.job, DebugLevel, format, args...)
}

func (logger *Logger) Warnf(format string, args ...any) {
    logger.Logf(logger.job, WarnLevel, format, args...)
}

func (logger *Logger) Infof(format string, args ...any) {
    logger.Logf(logger.job, InfoLevel, format, args...)
}

func (logger *Logger) Errorf(format string, args ...any) {
    logger.Logf(logger.job, ErrorLevel, format, args...)
}

func (logger *Logger) Fatalf(format string, args ...any) {
    logger.Logf(logger.job, FatalLevel, format, args...)
}
