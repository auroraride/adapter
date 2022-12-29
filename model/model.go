// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package model

const (
    BinIDKey     = "bin-%d"
    CabinetIDKey = "cabinet-%d"
)

type Bool bool

const (
    True  Bool = true
    False Bool = false
)

func (b Bool) String() string {
    switch b {
    case True:
        return "是"
    default:
        return "否"
    }
}

type (
    VoidFunc      func()
    BytesCallback func(b []byte)
)

type CabinetSyncRequest struct {
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}
