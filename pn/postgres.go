// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package pn

import (
    "github.com/auroraride/adapter"
    "github.com/goccy/go-json"
    "golang.org/x/exp/slices"
)

type Action string

func (a Action) String() string {
    return string(a)
}

const (
    Update Action = "UPDATE"
    Delete Action = "DELETE"
    Insert Action = "INSERT"
)

const (
    ChannelCabinet Channel = "cabinet"
    ChannelBin     Channel = "bin"
)

var (
    Channels = []Channel{
        ChannelCabinet,
        ChannelBin,
    }
)

type Channel string

func (c Channel) String() string {
    return string(c)
}

func (c Channel) Validate() error {
    if slices.Contains(Channels, c) {
        return nil
    }
    return adapter.ErrorData
}

type Message[T any] struct {
    Table  Channel `json:"table"`
    Action Action  `json:"action"`
    Data   T       `json:"data"`
}

func ParseData[T any](b []byte) (data T, err error) {
    var v Message[T]
    err = json.Unmarshal(b, &v)
    if err != nil {
        return
    }

    data = v.Data
    return
}
