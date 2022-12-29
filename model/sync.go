// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package model

type DataType string

const (
    DataTypeCabinetSync DataType = "cabinet_sync"
)

type DataStuff interface {
    *CabinetSyncData
}

type CabinetSyncData struct {
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}

type Data[T DataStuff] struct {
    Type  DataType `json:"type"`
    Value T        `json:"value"`
}
