// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "errors"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/message"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "time"
)

type Client struct {
    *Tcp

    conn   gnet.Conn
    Sender chan message.Messenger
}

func NewClient(addr string, l adapter.ZapLogger, c codec.Codec) *Client {
    cli := &Client{
        Tcp:    NewTcp(addr, l, c, nil),
        Sender: make(chan message.Messenger),
    }
    cli.Tcp.closeCh = make(chan bool)
    return cli
}

func (c *Client) Run() {
    for {
        err := c.dial()
        c.logger.Info(
            "连接失败, 5s后重试连接...",
            c.logserv,
            zap.Error(err),
        )
        time.Sleep(5 * time.Second)
    }
}

func (c *Client) dial() (err error) {
    if c.Hooks.Start != nil {
        c.Hooks.Start()
    }

    var (
        cli *gnet.Client
    )

    cli, err = gnet.NewClient(
        c,
        gnet.WithReuseAddr(true),
    )
    if err != nil {
        return
    }
    err = cli.Start()
    if err != nil {
        return
    }

    defer cli.Stop()

    c.conn, err = cli.Dial("tcp", c.address)
    if err != nil {
        return
    }

    if c.Hooks.Connect != nil {
        go c.Hooks.Connect()
    }

    for {
        select {
        case data := <-c.Sender:
            go c.Send(data)
            // var b []byte
            // b, err = message.Pack(data)
            // if err != nil {
            //     return
            // }
            //
            // encoded := c.codec.Encode(b)
            // _, err = c.Conn.Write(encoded)
            // // encoded, err = c.Conn.Send(data)
            // if err != nil {
            //     c.logger.Info(
            //         "消息发送失败",
            //         c.logserv,
            //         zap.Error(err),
            //         zap.Binary("encoded", encoded),
            //     )
            // }
        case <-c.closeCh:
            if err == nil {
                err = errors.New("未知原因断开连接")
            }
            if c.Hooks.Close != nil {
                go c.Hooks.Close()
            }
            return
        default:
            _, err = c.codec.Decode(c.conn)
            if err != nil && err != adapter.ErrorIncompletePacket {
                c.logger.Info(
                    "消息读取失败",
                    c.logserv,
                    zap.Error(err),
                )
                c.closeCh <- true
            }
        }
    }
}

func (c *Client) Send(data message.Messenger) {
    var (
        packed  []byte
        encoded []byte
        err     error
    )
    defer func() {
        if err != nil {
            c.logger.Info(
                "消息发送失败",
                c.logserv,
                zap.Error(err),
                zap.Binary("encoded", encoded),
                zap.Binary("packed", packed),
            )
        }
    }()

    packed, err = message.Pack(data)
    if err != nil {
        return
    }

    encoded = c.codec.Encode(packed)
    _, err = c.conn.Write(encoded)
    return
}
