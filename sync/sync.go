// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-30
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/auroraride/adapter"
)

type Stream string

const (
	StreamExchange    Stream = "SYNC:EXCHANGE"
	StreamBatteryFlow Stream = "SYNC:BATTERYFLOW"
)

func (s Stream) String() string {
	return string(s)
}

type Receiver[T any] func([]*T)

type Sync[T any] struct {
	client *redis.Client

	receiver Receiver[T]
	name     string
	key      string
	stream   Stream
}

func New[T any](client *redis.Client, e adapter.Environment, stream Stream, reader Receiver[T]) *Sync[T] {
	return &Sync[T]{
		client:   client,
		name:     e.UpperString() + ":" + stream.String(),
		receiver: reader,
		key:      "__DATA__",
		stream:   stream,
	}
}

func (s *Sync[T]) Run() {
	ctx := context.Background()
	id := "0"

	xReadArgs := &redis.XReadArgs{
		Streams: []string{s.name, id},
		Count:   100,
		Block:   0,
	}

	for {
		results, err := s.client.XRead(ctx, xReadArgs).Result()
		if err != nil {
			zap.L().WithOptions(zap.WithCaller(false)).Error("[SYNC] 同步消息读取失败", zap.Error(err))
			continue
		}
		if len(results) > 0 {
			var items []*T

			for _, result := range results {
				for _, message := range result.Messages {
					id = message.ID
					s.client.XDel(ctx, s.name, id)

					var item *T
					item, err = Unmarshal[T](s.key, message.Values)
					if err != nil {
						zap.L().WithOptions(zap.WithCaller(false)).Error("[SYNC] 同步消息解析失败", zap.Error(err))
						return
					}
					items = append(items, item)
				}
			}

			if len(items) > 0 {
				s.receiver(items)
			}
		}
	}
}

func (s *Sync[T]) Push(data any) {
	m, b, err := Marshal(s.key, data)
	// zap.L().WithOptions(zap.WithCaller(false)).Info("[SYNC] 发送同步消息", zap.ByteString("data", b))

	if err != nil {
		zap.L().WithOptions(zap.WithCaller(false)).Error("[SYNC] 同步消息格式化失败", zap.ByteString("data", b), zap.Error(err))
	}

	err = s.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: s.name,
		ID:     "*",
		Values: m,
	}).Err()

	if err != nil {
		zap.L().WithOptions(zap.WithCaller(false)).Error("[SYNC] 同步消息发送失败", zap.ByteString("data", b), zap.Error(err))
	}
}
