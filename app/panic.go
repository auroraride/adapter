// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package app

import (
    "github.com/auroraride/adapter"
    "net/http"
)

func Panic(params ...any) {
    r := CreateResponse(params...)
    if r.Code == http.StatusOK {
        r.Code = http.StatusInternalServerError
    }

    if r.Message == "" {
        switch r.Code {
        case http.StatusBadRequest:
            r.Message = adapter.ErrorBadRequest.Error()
        default:
            r.Message = adapter.ErrorInternalServer.Error()
        }
    }

    panic(r)
}
