// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/maintain"
)

type EchoConfig struct {
	AuthSkipper middleware.Skipper
	Maintain    maintain.Config
	DumpSkipper middleware.Skipper
}

func NewEcho(cfg *EchoConfig) (e *echo.Echo) {
	e = echo.New()

	// 默认json序列化工具
	e.JSONSerializer = &adapter.DefaultJSONSerializer{}

	// http请求错误
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		ctx := Context(c)
		message := err
		code := http.StatusInternalServerError
		var data any
		switch err.(type) {
		case *echo.HTTPError:
			target := err.(*echo.HTTPError)
			message = fmt.Errorf("%v", target.Message)
		}
		_ = ctx.SendResponse(code, message, data)
	}

	// 未找到
	echo.NotFoundHandler = func(c echo.Context) error {
		return Context(c).SendResponse(http.StatusNotFound, adapter.ErrorNotFound)
	}

	// 错误的请求方式
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		routerAllowMethods, ok := c.Get(echo.ContextKeyHeaderAllow).(string)
		if ok && routerAllowMethods != "" {
			c.Response().Header().Set(echo.HeaderAllow, routerAllowMethods)
		}
		return Context(c).SendResponse(http.StatusBadRequest, fmt.Errorf("%v", echo.ErrMethodNotAllowed.Message))
	}

	// 绑定校验器
	e.Validator = NewValidator()

	// 获取远程IP
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// middlewares
	e.Use(
		ContextMiddleware(),
		RecoverMiddleware(),
		UserMiddleware(cfg.AuthSkipper),
		NewDumpLoggerMiddleware().WithDefaultConfig(cfg.DumpSkipper),
	)

	// 运维接口
	e.GET("/maintain/update/:token", maintain.NewController(cfg.Maintain, Quit).UpdateApi)
	return
}
