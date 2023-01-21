// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-17
// Based on adapter by liasica, magicrolan@qq.com.

package exhook

import (
    "bytes"
    log "github.com/sirupsen/logrus"
    "testing"
)

func TestRun(t *testing.T) {
    var (
        splitter = []byte("/")
        a1       = []byte("A1")
        a2       = []byte("A2")
    )

    s := NewServer(log.StandardLogger(), HookMessagePublish, HookMessageDelivered)
    s.OnMessageReceived = func(in *MessagePublishRequest) (reply *Message) {
        topic := bytes.Split([]byte(in.Message.Topic), splitter)
        reply = in.Message
        if bytes.Equal(topic[2], a1) {
            topic[2] = a2

            reply.Payload = topic[1]
            reply.Topic = string(bytes.Join(topic, splitter))
        }

        return
    }

    s.Run(":9801")
}