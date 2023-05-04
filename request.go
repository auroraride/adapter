// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-05
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type ResponseVerifiable interface {
	Verify() error
}

type AurResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func (r *AurResponse[T]) Verify() error {
	if r.Code == 0 {
		return nil
	}
	return errors.New(r.Message)
}

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func (r *Response[T]) Verify() error {
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

func Post[T any](url string, user *User, payload any, params ...any) (data T, err error) {
	var cb func(*resty.Response)

	for _, param := range params {
		switch v := param.(type) {
		case func(*resty.Response):
			cb = v
		}
	}

	var r *resty.Response
	res := new(Response[T])
	r, err = CreateRequest(user).SetBody(payload).SetResult(res).Post(url)
	if cb != nil {
		cb(r)
	}

	if err != nil {
		return
	}

	data = res.Data
	err = res.Verify()
	return
}

type RequestHeader struct {
	Key   string
	Value string
}

type RequestMethod string

const (
	RequestMethodGet RequestMethod = "Get"
	RquestMethodPost RequestMethod = "POST"
)

func (r RequestMethod) String() string {
	return string(r)
}

func FastRequest[T ResponseVerifiable](url string, method RequestMethod, params ...any) (data T, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	if method != RequestMethodGet {
		req.Header.SetMethod(method.String())
	}

	for _, param := range params {
		switch v := param.(type) {
		case []byte:
			req.SetBody(v)
		case *User:
			req.Header.Set(HeaderUserID, v.ID)
			req.Header.Set(HeaderUserType, v.Type.String())
		case *RequestHeader:
			req.Header.Set(v.Key, v.Value)
		case RequestHeader:
			req.Header.Set(v.Key, v.Value)
		}
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(resp.Body(), &data)

	if err != nil {
		return
	}

	return data, data.Verify()
}
