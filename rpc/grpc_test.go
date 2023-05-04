// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-05
// Based on adapter by liasica, magicrolan@qq.com.

package rpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/auroraride/adapter/rpc/pb"
)

var (
	testAddress = "127.0.0.1:60010"
)

type testServer struct {
	pb.UnimplementedEchoServer
}

func (*testServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: time.Now().Format("2006-01-02 15:04:05.999") + " -> " + req.Message}, nil
}

func (*testServer) BidirectionalStreamingEcho(srv pb.Echo_BidirectionalStreamingEchoServer) error {
	for {
		recv, err := srv.Recv()
		if err != nil {
			return err
		}
		fmt.Println(recv.Message)
		if err = srv.Send(&pb.EchoResponse{Message: "[S] " + recv.Message}); err != nil {
			return err
		}
	}
}

func TestServer(t *testing.T) {
	err := NewServer(testAddress, func(s *grpc.Server) {
		pb.RegisterEchoServer(s, &testServer{})
	})
	require.NoError(t, err)
}

func TestClient(t *testing.T) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	logger, _ := cfg.Build()
	grpczap.ReplaceGrpcLoggerV2(logger)

	c, err := NewClient[pb.EchoClient](testAddress, pb.NewEchoClient)
	require.NoError(t, err)

	for {
		t.Log("performing unary request at:", time.Now().Format("2006-01-02 15:04:05.999"))

		var res *pb.EchoResponse
		res, err = c.UnaryEcho(context.Background(), &pb.EchoRequest{Message: "keepalive demo"})

		t.Log("request done at:", time.Now().Format("2006-01-02 15:04:05.999"))
		require.NoError(t, err)

		t.Log("RPC response:", res)
		time.Sleep(3 * time.Second)
	}
}
