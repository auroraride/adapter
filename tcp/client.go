// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
    jsoniter "github.com/json-iterator/go"
    "net"
    "time"
)

type Client struct {
    net.Conn
    *Tcp

    Sender    chan any
    reconnect chan bool
    readStop  chan bool
    writeStop chan bool
}

func NewClient(addr string, l logger.StdLogger, c codec.Codec) *Client {
    return &Client{
        Sender:    make(chan any),
        reconnect: make(chan bool),
        readStop:  make(chan bool),
        writeStop: make(chan bool),
        Tcp:       NewTcp(addr, l, c, nil),
    }
}

func (c *Client) Send(data any) (err error) {
    b, _ := jsoniter.Marshal(data)
    _, err = c.Write(c.codec.Encode(b))
    return
}

func (c *Client) Run() {

    for {
        select {
        case <-c.reconnect:
            c.writeStop <- true
            c.readStop <- true
            c.logger.Errorf("[ADAPTER] TCP (%s) 连接失败, 3s后重试连接...", c.address)
            time.Sleep(3 * time.Second)
            c.dial()
        }
    }
}

func (c *Client) dial() {
    conn, err := net.Dial("tcp", c.address)
    if err != nil {
        c.logger.Errorf("[ADAPTER] TCP (%s) 连接错误: %v", c.address, err)
        return
    }

    defer conn.Close()

    if c.Hooks.Connect != nil {
        c.Hooks.Connect()
    }

    c.Conn = conn

    go c.readBump()
    go c.writeBump()
}

func (c *Client) readBump() {
    for {

    }
}

func (c *Client) writeBump() {
    for {
        select {
        case <-c.writeStop:
            return
        case data := <-c.Sender:
            err := c.Send(data)
            if err != nil {
                c.logger.Errorf("[ADAPTER] 消息发送失败 (%#v): %v", data, err)
            }
        }
    }
}
