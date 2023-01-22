// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-20
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "github.com/go-redis/redis/v9"
    "testing"
)

func TestNew(t *testing.T) {
    writer := NewRedisWriter(redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6379",
        DB:   0,
    }))
    New("test", writer, true)
    Info("test new message db 0")
}
