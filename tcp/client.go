// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "errors"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
    "github.com/panjf2000/gnet/v2"
    "time"
)

type Client struct {
    *Tcp

    Conn   *Conn
    Sender chan any
}

func NewClient(addr string, l logger.StdLogger, c codec.Codec) *Client {
    cli := &Client{
        Tcp:    NewTcp(addr, l, c, nil),
        Sender: make(chan any),
    }
    cli.Tcp.closeCh = make(chan bool)
    return cli
}

func (c *Client) Run() {
    for {
        err := c.dial()
        c.logger.Errorf("[ADAPTER] TCP (%s) 连接失败: %v, 5s后重试连接...", c.address, err)
        time.Sleep(5 * time.Second)
    }
}

func (c *Client) dial() (err error) {
    if c.Hooks.Start != nil {
        c.Hooks.Start()
    }

    var (
        cli  *gnet.Client
        conn gnet.Conn
    )

    cli, err = gnet.NewClient(
        c,
        gnet.WithLogger(c.logger),
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

    conn, err = cli.Dial("tcp", c.address)
    if err != nil {
        return
    }

    c.Conn = &Conn{
        Conn:  conn,
        codec: c.Tcp.codec,
    }

    if c.Hooks.Connect != nil {
        go c.Hooks.Connect()
    }

    for {
        select {
        case data := <-c.Sender:
            err = c.Conn.Send(data)
            if err != nil {
                c.logger.Errorf("[ADAPTER] 消息发送失败 (%#v): %v", data, err)
            }
        case <-c.closeCh:
            if err == nil {
                err = errors.New("未知原因断开连接")
            }
            if c.Hooks.Close != nil {
                go c.Hooks.Close()
            }
            return
        default:
            _, err = c.codec.Decode(c.Conn)
            if err != nil && err != codec.IncompletePacket {
                c.logger.Errorf("[ADAPTER] 消息读取失败: %v", err)
                c.closeCh <- true
            }
        }
    }
}
