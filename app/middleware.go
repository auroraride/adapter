// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-24
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "go.uber.org/zap"
    "net/http"
)

func ContextMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return next(NewBaseContext(c))
        }
    }
}

func RecoverMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := Context(c)

            defer func() {
                if r := recover(); r != nil {
                    switch v := r.(type) {
                    case *ApiResponse:
                        _ = ctx.SendResponse(v.Code, v.Message, v.Data)
                    default:
                        err := fmt.Errorf("%v", r)
                        zap.L().Named("panic").WithOptions(zap.WithCaller(false)).Error("捕获未处理崩溃", zap.Error(err))
                        _ = Context(c).SendResponse(http.StatusInternalServerError, err)
                    }
                }
            }()
            return next(ctx)
        }
    }
}

func UserMiddleware(skipper middleware.Skipper) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := Context(c)
            header := c.Request().Header

            // 获取user信息
            id := header.Get(adapter.HeaderUserID)
            typ := header.Get(adapter.HeaderUserType)

            if !skipper(c) && (id == "" || typ == "") {
                Panic(http.StatusUnauthorized, adapter.ErrorUserRequired)
            }

            user := &adapter.User{
                ID:   id,
                Type: adapter.UserType(typ),
            }
            ctx.User = user
            return next(ctx)
        }
    }
}

func UserTypeManagerMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := Context(c)

            if ctx.User == nil || ctx.User.Type != adapter.UserTypeManager {
                Panic(http.StatusForbidden, adapter.ErrorManagerRequired)
            }

            return next(ctx)
        }
    }
}
