// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-25
// Based on adapter by liasica, magicrolan@qq.com.

package maintain

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/async"
    "github.com/labstack/echo/v4"
    "golang.org/x/exp/slices"
    "net/http"
    "strings"
    "time"
)

type controller struct {
    cfg  Config
    quit chan bool
}

func NewController(cfg Config, quit chan bool) *controller {
    return &controller{
        cfg:  cfg,
        quit: quit,
    }
}

func (ctl *controller) UpdateApi(c echo.Context) (err error) {
    _ = Create()

    addr := c.Request().RemoteAddr
    n := strings.Index(addr, ":")
    host := addr[:n]
    fmt.Println("request update <<<", host)

    if !slices.Contains(ctl.cfg.IP, host) || c.Param("token") != ctl.cfg.Token {
        return adapter.ErrorForbidden
    }

    go ctl.doQuit()

    return c.JSON(http.StatusOK, map[string]any{
        "remote": c.Request().RemoteAddr,
    })
}

func (ctl *controller) doQuit() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for ; true; <-ticker.C {
        // 是否有进行中的异步业务
        if async.IsDone() {
            ctl.quit <- true
            return
        }
    }
}
