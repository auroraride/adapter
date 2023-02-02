// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-17
// Based on adapter by liasica, magicrolan@qq.com.

package exhook

import (
    "bytes"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/go-redis/redis/v9"
    "go.uber.org/zap"
    "io"
    "testing"
)

var (
    splitter = []byte("/")
    A1       = []byte("A1")
    A2       = []byte("A2")
    B1       = []byte("B1")
    B2       = []byte("B2")
    C1       = []byte("C1")
    C2       = []byte("C2")
)

func TestRun(t *testing.T) {
    log.Initialize(&log.Config{
        FormatJson:  true,
        Stdout:      true,
        Application: "exhook",
        Writers: []io.Writer{
            log.NewRedisWriter(redis.NewClient(&redis.Options{})),
        },
    })

    s := NewServer(HookMessagePublish, HookMessageDelivered)
    s.OnMessageReceived = func(in *MessagePublishRequest) (reply *Message) {
        topic := adapter.ConvertString2Bytes(in.Message.Topic)
        if len(topic) != 22 {
            zap.L().Error("topic长度应为22", zap.Error(adapter.ErrorData))
        }

        reply = in.Message
        code := topic[20:]

        switch {
        case bytes.Equal(code, A1):
            // 直接发送IMEI
            reply.Payload = topic[4:19]
            // 发布A2订阅
            topic[20], topic[21] = A2[0], A2[1]
            reply.Topic = adapter.ConvertBytes2String(topic)
        }

        return
    }

    s.Run(":9801")
}
