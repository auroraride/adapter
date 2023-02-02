// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/codec"
    "github.com/panjf2000/gnet/v2"
)

type Server struct {
    *Tcp
}

func NewServer(addr string, c codec.Codec, r adapter.BytesCallback) *Server {
    s := &Server{
        Tcp: NewTcp(addr, c, r),
    }
    return s
}

func (s *Server) Run() error {
    return gnet.Run(
        s,
        s.address,
        gnet.WithMulticore(true),
        gnet.WithReuseAddr(true),
    )
}
