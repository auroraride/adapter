// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-20
// Based on adapter by liasica, magicrolan@qq.com.

package zlog

import (
    "errors"
    "io"
    "net"
    "syscall"
)

var logstash *Logstash

type Logstash struct {
    conn    *net.TCPConn
    address string
    times   int
}

func (l *Logstash) Write(b []byte) (n int, err error) {
    n, err = l.conn.Write(b)
    switch {
    case
        errors.Is(err, net.ErrClosed),
        errors.Is(err, io.EOF),
        errors.Is(err, syscall.EPIPE):
        if l.times > 3 {
            panic("重试次数过多")
        }
        l.times += 1
        l.connect()
        return l.Write(b)
    }
    return
}

func (l *Logstash) connect() {
    l.times += 1

    addr, err := net.ResolveTCPAddr("tcp", l.address)
    if err != nil {
        panic(err)
    }

    l.conn, err = net.DialTCP("tcp", nil, addr)
    if err != nil {
        panic(err)
    }
    return
}

func getLogstash(address string) *Logstash {
    if logstash == nil {
        logstash = &Logstash{address: address}
        logstash.connect()
    }

    return logstash
}
