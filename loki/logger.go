// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/async"
)

type Logger struct {
	job          string
	url          string
	reportCaller bool
	sender       chan []byte

	Caller         *runtime.Frame
	CallerSplitter func(string) string
	Formatter      Formatter
	WaitGroup      *sync.WaitGroup
}

func New() *Logger {
	logger := &Logger{
		job:          "varjob",
		url:          "http://localhost:3100/loki/api/v1/push",
		Formatter:    &DefaultFormatter{},
		reportCaller: true,
		WaitGroup:    &sync.WaitGroup{},
		sender:       make(chan []byte),
	}

	go logger.run()

	return logger
}

func (logger *Logger) run() {
	for {
		select {
		case b := <-logger.sender:
			go async.WithTask(func() {
				logger.send(b)
			})
		}
	}
}

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

func (logger *Logger) Log(job string, level Level, args ...any) {
	if job == "" {
		return
	}

	if logger.reportCaller {
		logger.Caller = adapter.GetCaller(1)
	}

	str := fmt.Sprint(args...)

	msg := logger.Formatter.Format(level, str, logger)

	if len(msg) > 0 {
		logger.WaitGroup.Add(1)

		entry := NewEntry()
		defer ReleaseEntry(entry)

		entry.Streams = []EntryStream{
			{
				Stream: EntryJob{Job: job},
				Values: [][]string{
					{strconv.FormatInt(time.Now().UnixNano(), 10), string(msg)},
				},
			},
		}

		logger.sender <- entry.Bytes()
	}

	if level == FatalLevel {
		os.Exit(1)
	}
}

func (logger *Logger) Logf(job string, level Level, format string, args ...any) {
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
