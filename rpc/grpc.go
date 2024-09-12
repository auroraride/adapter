// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-05
// Based on adapter by liasica, magicrolan@qq.com.

package rpc

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type ServerRegister func(s *grpc.Server)

func NewServer(address string, register ServerRegister, options ...grpc.ServerOption) (err error) {
	var lis net.Listener
	lis, err = net.Listen("tcp", address)
	if err != nil {
		return
	}

	options = append(
		options,
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
	)

	s := grpc.NewServer(options...)
	register(s)

	defer s.GracefulStop()

	return s.Serve(lis)
}

type ClientRegister func(conn *grpc.ClientConn)

func NewClient[T any](address string, register func(grpc.ClientConnInterface) T, options ...grpc.DialOption) (conn T, err error) {
	options = append(options,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		// grpc.WithBlock(),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(rtcp...)),
		// grpc.WithStreamInterceptor(grpcretry.StreamClientInterceptor(rtcp...)),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var c *grpc.ClientConn
	c, err = grpc.DialContext(ctx, address, options...)

	if err != nil {
		return
	}

	return register(c), nil
}

func NeedReconnect(err error) bool {
	s, ok := status.FromError(err)
	if !ok {
		return false
	}
	return s.Code() == codes.Unavailable
}
