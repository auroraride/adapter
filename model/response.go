// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package model

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Data    any    `json:"data,omitempty"`
}
