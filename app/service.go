// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
    "context"
    "github.com/auroraride/adapter"
    "net/http"
)

type Permission bool

const (
    PermissionRequired    Permission = true
    PermissionNotRequired Permission = false
)

type BaseService struct {
    User *adapter.User
    ctx  context.Context
}

func newService(params ...any) *BaseService {
    nq := PermissionRequired
    s := &BaseService{
        ctx: context.Background(),
    }
    for _, param := range params {
        switch v := param.(type) {
        case *adapter.User:
            s.User = v
        case Permission:
            nq = v
        case context.Context:
            s.ctx = v
        }
    }
    if s.User == nil && nq {
        Panic(http.StatusUnauthorized, adapter.ErrorUserRequired)
    }
    return s
}
