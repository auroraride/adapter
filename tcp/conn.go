// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/message"
    "github.com/panjf2000/gnet/v2"
)

type Conn struct {
    gnet.Conn

    codec codec.Codec
}

func (c *Conn) Send(data message.Messenger) (err error) {
    var b []byte
    b, err = message.Pack(data)
    if err != nil {
        return
    }

    _, err = c.Write(c.codec.Encode(b))
    return
}
