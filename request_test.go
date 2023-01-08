// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-08
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/go-resty/resty/v2"
    "github.com/labstack/echo/v4"
    "net/http"
    "testing"
)

func startEchoTestServer() {
    e := echo.New()
    e.Any("/test", func(c echo.Context) error {
        return c.JSON(http.StatusOK, &Response[string]{Code: http.StatusOK, Data: "ok"})
    })
    _ = e.Start(":8833")
}

func TestPost(t *testing.T) {
    go startEchoTestServer()

    res, err := Post[string]("http://localhost:8833/test", &User{
        Type: UserTypeUnknown,
        ID:   "userID",
    }, nil, func(r *resty.Response) {
        t.Log(string(r.Body()))
    })
    if err != nil {
        t.Fail()
    }
    t.Log(res)
}
