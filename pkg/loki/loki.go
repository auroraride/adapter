// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-11
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import "runtime"

var (
    std = New()
)

type Loki struct {
    job string
    api string
}

func New() *Loki {
    runtime.Caller(0)
    return &Loki{
        job: "job",
        api: "http://localhost:3100/loki/api/v1/push",
    }
}

func StandardLogger() *Loki {
    return std
}

func SetApi(api string) *Loki {
    std.api = api
    return std
}

func SetJob(job string) *Loki {
    std.job = job
    return std
}
