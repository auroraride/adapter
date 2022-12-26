// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "errors"
    "fmt"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
    "github.com/auroraride/adapter/model"
    "github.com/panjf2000/gnet/v2"
    "time"
)

type Client struct {
    *Tcp

    Conn   *Conn
    Sender chan *model.CabinetSyncRequest
    stop   chan bool
}

func NewClient(addr string, l logger.StdLogger, c codec.Codec) *Client {
    cli := &Client{
        Tcp:    NewTcp(addr, l, c, nil),
        Sender: make(chan *model.CabinetSyncRequest),
        stop:   make(chan bool),
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
        c.Hooks.Connect()
    }

    go c.readPump()

    for {
        select {
        case data := <-c.Sender:
            err = c.Conn.Send(data)
            if err != nil {
                c.logger.Errorf("[ADAPTER] 消息发送失败 (%#v): %v", data, err)
            }
        case <-c.closeCh:
            c.stop <- true
            err = errors.New("未知原因断开连接")
            return
        }
    }
}

func (c *Client) readPump() {
    for {
        select {
        case <-c.stop:
            return
        default:
            _, err := c.codec.Decode(c.Conn)
            if err != nil && err != codec.IncompletePacket {
                fmt.Println(err)
            }
        }
    }
}
