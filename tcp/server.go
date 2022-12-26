// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
    "github.com/auroraride/adapter/model"
    "github.com/panjf2000/gnet/v2"
)

type Server struct {
    *Tcp
}

func NewServer(addr string, l logger.StdLogger, c codec.Codec, r model.BytesCallback) *Server {
    s := &Server{
        Tcp: NewTcp(addr, l, c, r),
    }
    return s
}

func (s *Server) Run() error {
    return gnet.Run(
        s,
        s.address,
        gnet.WithMulticore(true),
        gnet.WithReuseAddr(true),
        gnet.WithLogger(s.logger),
    )
}
