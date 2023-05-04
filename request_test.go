// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-08
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func startEchoTestServer() {
	e := echo.New()
	e.Any("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Response[string]{Code: http.StatusOK, Data: "ok"})
	})
	_ = e.Start(":8833")
}

func TestPost(t *testing.T) {
	go startEchoTestServer()

	res, err := Post[string]("http://localhost:8833/test", &User{
		Type: UserTypeUnknown,
		ID:   "userID",
	}, nil, func(r *resty.Response) {
		t.Log(string(r.Body()))
	})
	if err != nil {
		t.Fail()
	}
	t.Log(res)
}

func TestFastRequest(t *testing.T) {
	type data struct {
		Name string `json:"name"`
	}
	res, err := FastRequest[*AurResponse[data]]("http://localhost:5533/kit/cabinet/name/CH7208KXHD220408016", RequestMethodGet)

	require.Nil(t, err)
	require.Equal(t, data{Name: "斜口"}, res.Data)
}
