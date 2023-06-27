// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-24
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/auroraride/adapter"
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
						zap.L().WithOptions(zap.WithCaller(false)).Error("捕获未处理崩溃", zap.Error(err), zap.Stack("stack"))
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

func UserTypeMiddleware(typ adapter.UserType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := Context(c)

			if ctx.User == nil || ctx.User.Type != typ {
				Panic(http.StatusForbidden, adapter.ErrorManagerRequired)
			}

			return next(ctx)
		}
	}
}

func UserTypeManagerMiddleware() echo.MiddlewareFunc {
	return UserTypeMiddleware(adapter.UserTypeManager)
}

func UserTypeAgentMiddleware() echo.MiddlewareFunc {
	return UserTypeMiddleware(adapter.UserTypeAgent)
}
