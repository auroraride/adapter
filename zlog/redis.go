// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-21
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import "github.com/go-redis/redis"

type RedisWriter struct {
    cli *redis.Client
    key string
}

func NewRedisWriter(cli *redis.Client) *RedisWriter {
    return &RedisWriter{
        cli: cli,
        key: "application-log",
    }
}

func (w *RedisWriter) Write(p []byte) (int, error) {
    n, err := w.cli.RPush(w.key, p).Result()
    return int(n), err
}
