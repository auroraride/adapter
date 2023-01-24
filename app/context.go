// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package app

import (
    "github.com/auroraride/adapter"
    "github.com/labstack/echo/v4"
    "net/http"
)

type BaseContext struct {
    echo.Context
    User *adapter.User
}

func NewBaseContext(c echo.Context) *BaseContext {
    return &BaseContext{Context: c}
}

func Context(c echo.Context) *BaseContext {
    ctx, ok := c.(*BaseContext)
    if ok {
        return ctx
    }
    return NewBaseContext(c)
}

func (c *BaseContext) BindValidate(ptr any) {
    err := c.Bind(ptr)
    if err != nil {
        Panic(http.StatusBadRequest)
    }
    err = c.Validate(ptr)
    if err != nil {
        Panic(http.StatusBadRequest)
    }
}

func ContextAndBinding[T any](c echo.Context) (ctx *BaseContext, req *T) {
    ctx = Context(c)
    req = new(T)
    ctx.BindValidate(req)
    return
}
