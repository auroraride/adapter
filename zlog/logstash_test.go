// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-20
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "go.uber.org/zap"
    "testing"
)

func TestNewLogstash(t *testing.T) {
    writer := getLogstash("127.0.0.1:50000")
    // n, err := writer.Write([]byte(`{"service":"test","log.level":"info","@timestamp":"2023-01-21T08:20:27.613+0800","log.origin":{"file.name":"zlog/zap.go","file.line":25},"message":"some logging info","count":17,"error":{"message":"boom"},"ecs.version":"1.6.0"}`))
    // t.Log(n, err)
    New("test", writer, true)
    Named("XT").Info("test named message")
    Info("this is test message", zap.Int("int", 1))
    Named("YT").Info("test named message")
    Info("this is test message", zap.Int("int", 1))
}
