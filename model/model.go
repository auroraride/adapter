// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package model

type (
    VoidFunc      func()
    BytesCallback func(b []byte)
)

type CabinetSyncRequest struct {
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}
