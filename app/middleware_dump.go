// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	ew "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/auroraride/adapter"
)

var (
	dumpReqHeader  = []byte("Request Header")
	dumpReqBody    = []byte("Request Body")
	dumpResHeader  = []byte("Response Header")
	dumpResBody    = []byte("Response Body")
	dumpEqual      = append(adapter.Space, append(adapter.Equal, adapter.Space...)...)
	dumpLeftSplit  = append(bytes.Repeat(adapter.Hyphen, 4), adapter.LeftSquareBracket...)
	dumpRightSplit = append(adapter.RightSquareBracket, append(bytes.Repeat(adapter.Hyphen, 4), adapter.Newline...)...)
)

type DumpHandler func(echo.Context, []byte, []byte)

type HeaderSkipper func(string) bool

type DumpConfig struct {
	Skipper ew.Skipper

	RequestHeader        bool
	RequestHeaderSkipper HeaderSkipper

	ResponseHeader        bool
	ResponseHeaderSkipper HeaderSkipper

	ResponseBodySkipper ew.Skipper

	Extra func(echo.Context) []byte
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

func dumpBuffer(cfg *DumpConfig, c echo.Context, reqBody, resBody []byte) []byte {
	// if skip dump
	if cfg.Skipper != nil && cfg.Skipper(c) {
		return nil
	}

	var buffer bytes.Buffer

	// log time
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05.00000"))

	// log [METHOD]
	buffer.Write(adapter.Space)
	buffer.Write(adapter.LeftSquareBracket)
	buffer.WriteString(c.Request().Method)
	buffer.Write(adapter.RightSquareBracket)
	buffer.Write(adapter.Space)

	// log uri \n
	buffer.WriteString(c.Request().RequestURI)
	buffer.Write(adapter.Newline)

	// log request header
	if cfg.RequestHeader {
		// ----[Request Header]----
		buffer.Write(dumpLeftSplit)
		buffer.Write(dumpReqHeader)
		buffer.Write(dumpRightSplit)

		// TODO c.Request().Header.Write
		// k = v
		for _, s := range getHeaders(c.Request().Header, cfg.RequestHeaderSkipper) {
			buffer.WriteString(s)
			buffer.Write(adapter.Newline)
		}
	}

	// log request body
	if len(reqBody) > 0 {
		// ----[Request Body]----
		buffer.Write(dumpLeftSplit)
		buffer.Write(dumpReqBody)
		buffer.Write(dumpRightSplit)
		buffer.Write(reqBody)
		buffer.Write(adapter.Newline)
	}

	// log response header
	if cfg.ResponseHeader {
		// ----[Response Header]----
		buffer.Write(dumpLeftSplit)
		buffer.Write(dumpResHeader)
		buffer.Write(dumpRightSplit)

		// k = v

		for _, s := range getHeaders(c.Response().Header(), cfg.ResponseHeaderSkipper) {
			buffer.WriteString(s)
			buffer.Write(adapter.Newline)
		}
	}

	// log response body
	if len(resBody) > 0 {
		// ----[Response Body]----
		buffer.Write(dumpLeftSplit)
		buffer.Write(dumpResBody)
		buffer.Write(dumpRightSplit)
		buffer.Write(resBody)
		buffer.Write(adapter.Newline)
	}

	buffer.Write(adapter.Newline)

	return buffer.Bytes()
}

func dump(handler DumpHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// Request
			var reqBody []byte
			if c.Request().Body != nil { // Read
				reqBody, _ = io.ReadAll(c.Request().Body)
			}
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			// Response
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &DumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			err = next(c)

			// if err != nil {
			//     c.Error(err)
			// }

			// Callback
			handler(c, reqBody, resBody.Bytes())

			return
		}
	}
}

type DumpFileConfig struct {
	Path string
}

type DumpFileMiddleware struct {
	ch   chan []byte
	path string
	day  int
}

func NewDumpFile(params ...any) *DumpFileMiddleware {
	d := "runtime/dump"
	for _, param := range params {
		switch v := param.(type) {
		case *DumpFileConfig:
			d = v.Path
		}
	}

	// create api log path
	if err := adapter.CreateDirectoryIfNotExist(d); err != nil {
		panic(err)
	}

	mw := &DumpFileMiddleware{
		ch:   make(chan []byte),
		path: d,
		day:  time.Now().Day(),
	}

	go func() {
		for {
			select {
			case b := <-mw.ch:
				if len(b) > 0 {
					mw.write(b)
				}
			}
		}
	}()

	return mw
}

func (mw *DumpFileMiddleware) write(b []byte) {
	p := filepath.Join(mw.path, time.Now().Format("2006-01-02")+".log")
	f, err := os.OpenFile(p, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return
	}

	defer func() {
		_ = f.Close()
	}()

	_, _ = f.Write(b)
}

func (mw *DumpFileMiddleware) WithDefaultConfig() echo.MiddlewareFunc {
	return mw.WithConfig(&DumpConfig{
		RequestHeader:  true,
		ResponseHeader: false,
	})
}

func (mw *DumpFileMiddleware) WithConfig(cfg *DumpConfig) echo.MiddlewareFunc {
	return dump(func(c echo.Context, reqBody []byte, resBody []byte) {
		mw.ch <- dumpBuffer(cfg, c, reqBody, resBody)
	})
}

type DumpZapLoggerMiddleware struct {
}

func NewDumpLoggerMiddleware() *DumpZapLoggerMiddleware {
	return &DumpZapLoggerMiddleware{}
}

func getHeaders(headers http.Header, skipper HeaderSkipper) (strs []string) {
	for k := range headers {
		if skipper != nil && skipper(k) {
			continue
		}
		strs = append(strs, k+" = "+headers.Get(k))
	}
	return
}

func (mw *DumpZapLoggerMiddleware) WithConfig(cfg *DumpConfig) echo.MiddlewareFunc {
	return dump(func(c echo.Context, reqBody []byte, resBody []byte) {
		if cfg.Skipper != nil && cfg.Skipper(c) {
			return
		}

		fields := []zap.Field{
			zap.String("remote_addr", c.Request().RemoteAddr),
		}

		// log request header
		if cfg.RequestHeader {
			fields = append(fields, zap.Strings("request_header", getHeaders(c.Request().Header, cfg.RequestHeaderSkipper)))
		}

		// log request body
		if len(reqBody) > 0 {
			fields = append(fields, zap.ByteString("request_body", reqBody))
		}

		// log response header
		if cfg.ResponseHeader {
			fields = append(fields, zap.Strings("response_header", getHeaders(c.Response().Header(), cfg.ResponseHeaderSkipper)))
		}

		if cfg.ResponseBodySkipper == nil {
			cfg.ResponseBodySkipper = func(c echo.Context) bool {
				return false
			}
		}

		// log response body
		if len(resBody) > 0 && !cfg.ResponseBodySkipper(c) {
			fields = append(fields, zap.ByteString("response_body", resBody))
		}

		if cfg.Extra != nil {
			extraData := cfg.Extra(c)
			if extraData != nil {
				fields = append(fields, zap.ByteString("extra", extraData))
			}
		}

		// x := adapter.GetCaller(0)
		go zap.L().Named("dump").WithOptions(zap.WithCaller(false)).Info(
			"["+c.Request().Method+"] "+c.Request().RequestURI,
			fields...,
		)
	})
}

func (mw *DumpZapLoggerMiddleware) WithDefaultConfig(skipper ew.Skipper) echo.MiddlewareFunc {
	return mw.WithConfig(&DumpConfig{
		RequestHeader:  true,
		ResponseHeader: false,
		Skipper:        skipper,
	})
}
