// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "errors"
    "net/http"
)

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Data    any    `json:"data,omitempty"`
}

type ResponseStuff[T any] struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Data    T      `json:"data,omitempty"`
}

func (r *ResponseStuff[T]) VerifyResponse() error {
    if r.Code == http.StatusOK {
        return nil
    }
    return errors.New(r.Message)
}
