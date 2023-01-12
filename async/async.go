// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-07
// Based on adapter by liasica, magicrolan@qq.com.

package async

import (
    "github.com/google/uuid"
    "sync"
)

var tasks *sync.Map

func init() {
    if tasks == nil {
        tasks = &sync.Map{}
    }
}

func WithTask(cb func()) {
    uid := uuid.New().String()
    // 添加异步任务
    tasks.Store(uid, 1)
    // 退出移除异步任务
    defer tasks.Delete(uid)

    cb()
}

func WithTaskReturn[T any](cb func() T) T {
    uid := uuid.New().String()
    // 添加异步任务
    tasks.Store(uid, 1)
    // 退出移除异步任务
    defer func() {
        tasks.Delete(uid)
    }()

    return cb()
}

func IsDone() bool {
    var n int
    if tasks == nil {
        return true
    }

    tasks.Range(func(_, _ any) bool {
        n += 1
        return true
    })

    return n == 0
}
