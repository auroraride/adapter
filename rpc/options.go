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
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	// keepalive 服务端默认参数
	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	// keepalive 客户端默认参数
	kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	// 重试默认参数
	rtcp = []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Canceled, codes.DataLoss, codes.Unavailable),
		grpcretry.WithMax(4),
		grpcretry.WithBackoff(grpcretry.BackoffLinear(time.Millisecond * 500)),
	}
)
