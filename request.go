// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-05
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "errors"
    "github.com/go-resty/resty/v2"
    jsoniter "github.com/json-iterator/go"
    "net/http"
)

type Response[T any] struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Data    T      `json:"data,omitempty"`
}

func (r *Response[T]) VerifyResponse() error {
    if r.Code == http.StatusOK {
        return nil
    }
    return errors.New(r.Message)
}

func CreateRequest(user *User) *resty.Request {
    client := resty.New()
    json := jsoniter.ConfigCompatibleWithStandardLibrary
    client.JSONMarshal = json.Marshal
    client.JSONUnmarshal = json.Unmarshal

    return client.R().
        SetHeaders(map[string]string{
            HeaderUserID:   user.ID,
            HeaderUserType: user.Type.String(),
        })
}

func Post[T any](url string, user *User, playload any, params ...any) (data T, err error) {
    var cb func(*resty.Response)

    for _, param := range params {
        switch v := param.(type) {
        case func(*resty.Response):
            cb = v
        }
    }

    var r *resty.Response
    res := new(Response[T])
    r, err = CreateRequest(user).SetBody(playload).SetResult(res).Post(url)
    if cb != nil {
        cb(r)
    }

    if err != nil {
        return
    }

    data = res.Data
    err = res.VerifyResponse()
    return
}
