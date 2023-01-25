// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/codec"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
)

type Hook struct {
    Boot    adapter.VoidFunc
    Start   adapter.VoidFunc
    Connect adapter.VoidFunc
    Close   adapter.VoidFunc
}

type Tcp struct {
    gnet.BuiltinEventEngine

    address   string
    codec     codec.Codec
    logger    *zap.Logger
    receiver  adapter.BytesCallback
    namespace string

    Hooks Hook

    closeCh chan bool
}

func NewTcp(addr string, l adapter.ZapLogger, c codec.Codec, receiver adapter.BytesCallback) *Tcp {
    return &Tcp{
        address:   addr,
        logger:    l.GetLogger().WithOptions(zap.AddCallerSkip(-2)),
        codec:     c,
        receiver:  receiver,
        namespace: "TCP",
    }
}

func (t *Tcp) OnBoot(gnet.Engine) (action gnet.Action) {
    t.logger.Named(t.namespace).Info("启动 -> " + t.address)

    if t.Hooks.Boot != nil {
        t.Hooks.Boot()
    }

    return gnet.None
}

func (t *Tcp) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    t.logger.Named(t.namespace).Info("已断开连接: " + c.RemoteAddr().String())
    if t.closeCh != nil {
        t.closeCh <- true
    }
    return
}

func (t *Tcp) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    t.logger.Named(t.namespace).Info("已开始连接: " + c.RemoteAddr().String())
    return
}

func (t *Tcp) OnTraffic(c gnet.Conn) (action gnet.Action) {
    var (
        b   []byte
        err error
    )

    for {
        b, err = t.codec.Decode(c)
        if err == adapter.ErrorIncompletePacket {
            break
        }
        if err != nil {
            t.logger.Named(t.namespace).Info(
                "消息读取失败",
                zap.Error(err),
            )
            return
        }

        go t.receiver(b)
    }

    return gnet.None
}
