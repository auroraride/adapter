// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-05
// Based on adapter by liasica, magicrolan@qq.com.

package rpc

import (
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
)

var (
	// keepalive 默认配置
	kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // 如果在指定时间内收到 ping 次数大于一次，强制断开连接，默认 5min
		PermitWithoutStream: true,            // 没有活动的 stream 也允许 ping，默认关闭
	}

	// keepalive 服务端默认参数
	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     30 * time.Second, // 最大空闲连接时间，默认为无限制。超出这段时间后，serve 发送 GoWay，强制 client stream 断开
		MaxConnectionAge:      30 * time.Second, // 最大连接时间，默认为无限制。stream 连接超出这个值是发送一个 GoWay
		MaxConnectionAgeGrace: 30 * time.Second, // 超出 MaxConnectionAge 之后的宽限时长，默认无限制 (最小为 1s)
		Time:                  5 * time.Second,  // 如果一段时间客户端存活但没有 ping 请求，服务端发送一次 ping 请求，默认是 2hour
		Timeout:               5 * time.Second,  // 服务端发送 ping 请求超时的时间，默认20s，在发送ping包之后，Timeout 时间内没有收到 ack 则视为超时
	}

	// keepalive 客户端默认参数
	kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // ping 请求间隔时间，默认无限制，最小为 10s
		Timeout:             5 * time.Second,  // ping 超时时间，默认是 20s
		PermitWithoutStream: true,             // 没有活动的 stream 也允许 ping，默认关闭
	}

	// 重试默认参数
	rtcp = []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
		grpcretry.WithMax(4),
		grpcretry.WithBackoff(grpcretry.BackoffLinear(time.Millisecond * 500)),
	}
)
