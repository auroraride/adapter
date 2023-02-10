// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-30
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/go-redis/redis/v9"
    "go.uber.org/zap"
)

type Stream string

const (
    StreamCabinet  Stream = "SYNC:CABINET"
    StreamExchange Stream = "SYNC:EXCHANGE"
)

func (s Stream) String() string {
    return string(s)
}

type Receiver[T any] func(*T)

type Sync[T any] struct {
    client *redis.Client

    receiver  Receiver[T]
    stream    string
    key       string
    namespace string
}

func New[T any](client *redis.Client, e adapter.Environment, stream Stream, reader Receiver[T]) *Sync[T] {
    return &Sync[T]{
        client:    client,
        stream:    e.UpperString() + ":" + stream.String(),
        receiver:  reader,
        namespace: "SYNC",
        key:       "__DATA__",
    }
}

func (s *Sync[T]) Run() {
    ctx := context.Background()
    xReadArgs := &redis.XReadArgs{
        Streams: []string{s.stream, "0-0"},
        Count:   1,
        Block:   0,
    }
    for {
        results, err := s.client.XRead(ctx, xReadArgs).Result()
        if err != nil {
            zap.L().Named(s.namespace).WithOptions(zap.WithCaller(false)).Error("同步消息读取失败", zap.Error(err))
            continue
        }
        if len(results) > 0 {
            for _, result := range results {
                for _, message := range result.Messages {
                    id := message.ID
                    go func() {
                        zap.L().Named(s.namespace).WithOptions(zap.WithCaller(false)).Info("收到同步消息", zap.String("values", message.Values[s.key].(string)))

                        var data *T
                        data, err = Unmarshal[T](s.key, message.Values)
                        if err != nil {
                            zap.L().Named(s.namespace).WithOptions(zap.WithCaller(false)).Error("同步消息解析失败", zap.Error(err))
                            return
                        }
                        s.receiver(data)
                    }()
                    s.client.XDel(ctx, s.stream, id)
                }
            }
        }
    }
}

func (s *Sync[T]) Push(data any) {
    m, err := Marshal(s.key, data)
    if err != nil {
        zap.L().Named(s.namespace).WithOptions(zap.WithCaller(false)).Error("同步消息格式化失败", zap.Error(err))
    }

    err = s.client.XAdd(context.Background(), &redis.XAddArgs{
        Stream: s.stream,
        ID:     "*",
        Values: m,
    }).Err()
    if err != nil {
        zap.L().Named(s.namespace).WithOptions(zap.WithCaller(false)).Error("同步消息发送失败", zap.Error(err))
    }
}