// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package app

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/auroraride/adapter"
)

type ApiResponse = adapter.Response[any]

type StatusResponse struct {
	Status bool `json:"status"` // 是否成功
}

// CreateResponse 创建response结构体
// int: code
// error: message
// string: message
// 其他: data
func CreateResponse(params ...any) *ApiResponse {
	r := &ApiResponse{
		Code: http.StatusOK,
	}
	for _, param := range params {
		switch v := param.(type) {
		case int:
			r.Code = v
		case error:
			r.Message = v.Error()
		case string:
			r.Message = v
		default:
			r.Data = v
		}
	}
	return r
}

// SendResponse 发送响应
func (c *BaseContext) SendResponse(params ...any) error {
	buffer := &bytes.Buffer{}
	encoder := jsoniter.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(CreateResponse(params...))
	return c.JSONBlob(http.StatusOK, buffer.Bytes())
}
