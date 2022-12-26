// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
    "github.com/auroraride/adapter/model"
    "github.com/panjf2000/gnet/v2"
)

type Hook struct {
    Boot    model.VoidFunc
    Connect model.VoidFunc
}

type Tcp struct {
    gnet.BuiltinEventEngine

    address  string
    codec    codec.Codec
    logger   logger.StdLogger
    receiver model.BytesCallback

    Hooks Hook
}

func NewTcp(addr string, l logger.StdLogger, c codec.Codec, receiver model.BytesCallback) *Tcp {
    return &Tcp{
        address:  addr,
        logger:   l,
        codec:    c,
        receiver: receiver,
    }
}
