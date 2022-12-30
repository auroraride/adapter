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
    Start   model.VoidFunc
    Connect model.VoidFunc
    Close   model.VoidFunc
}

type Tcp struct {
    gnet.BuiltinEventEngine

    address  string
    codec    codec.Codec
    logger   logger.StdLogger
    receiver model.BytesCallback

    Hooks Hook

    closeCh chan bool
}

func NewTcp(addr string, l logger.StdLogger, c codec.Codec, receiver model.BytesCallback) *Tcp {
    return &Tcp{
        address:  addr,
        logger:   l,
        codec:    c,
        receiver: receiver,
    }
}

func (t *Tcp) OnBoot(gnet.Engine) (action gnet.Action) {
    t.logger.Infof("[ADAPTER] TCP启动: %s", t.address)

    if t.Hooks.Boot != nil {
        t.Hooks.Boot()
    }

    return gnet.None
}

func (t *Tcp) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    t.logger.Infof("[ADAPTER] 已断开连接: %v", c.RemoteAddr())
    if t.closeCh != nil {
        t.closeCh <- true
    }
    return
}

func (t *Tcp) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    t.logger.Infof("[ADAPTER] 已开始连接: %v", c.RemoteAddr())
    return
}

func (t *Tcp) OnTraffic(c gnet.Conn) (action gnet.Action) {
    var (
        b   []byte
        err error
    )

    for {
        b, err = t.codec.Decode(c)
        if err == codec.IncompletePacket {
            break
        }
        if err != nil {
            t.logger.Errorf("[ADAPTER] 消息读取失败, err: %v", err)
            return
        }

        t.receiver(b)
    }

    return gnet.None
}
