// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package middleware

import (
    "bufio"
    "github.com/labstack/echo/v4"
    ew "github.com/labstack/echo/v4/middleware"
    "io"
    "net"
    "net/http"
)

type DumpConfig struct {
    Skipper               ew.Skipper
    RequestHeaderSkipper  ew.Skipper
    ResponseHeaderSkipper ew.Skipper
}

type DumpResponseWriter struct {
    io.Writer
    http.ResponseWriter
}

func (w *DumpResponseWriter) WriteHeader(code int) {
    w.ResponseWriter.WriteHeader(code)
}

func (w *DumpResponseWriter) Write(b []byte) (int, error) {
    return w.Writer.Write(b)
}

func (w *DumpResponseWriter) Flush() {
    w.ResponseWriter.(http.Flusher).Flush()
}

func (w *DumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
    return w.ResponseWriter.(http.Hijacker).Hijack()
}

func Dump() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return next(c)
        }
    }
}
