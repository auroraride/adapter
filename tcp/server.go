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
    address  string
    receiver model.BytesCallback

    *Tcp
    gnet.BuiltinEventEngine
}

func NewServer(addr string, l logger.StdLogger, c codec.Codec, r model.BytesCallback) *Server {
    return &Server{
        address:  addr,
        receiver: r,
        Tcp: &Tcp{
            logger: l,
            codec:  c,
        },
    }
}

func (s *Server) Run() {
    s.logger.Fatal(gnet.Run(
        s,
        s.address,
        gnet.WithMulticore(true),
        gnet.WithReuseAddr(true),
        gnet.WithLogger(s.logger),
    ))
}

func (s *Server) OnBoot(_ gnet.Engine) (action gnet.Action) {
    s.logger.Infof("[ADAPTER] TCP服务器已启动 %s", s.address)
    return gnet.None
}

func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {
    var (
        b   []byte
        err error
    )

    for {
        b, err = s.codec.Decode(c)
        if err == codec.IncompletePacket {
            break
        }
        if err != nil {
            s.logger.Errorf("[ADAPTER] 消息读取失败, err: %v", err)
            return
        }

        s.receiver(b)
    }

    return gnet.None
}
