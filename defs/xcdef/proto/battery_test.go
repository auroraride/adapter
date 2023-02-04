// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-03
// Based on adapter by liasica, magicrolan@qq.com.

package proto

import (
    "context"
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/go-redis/redis/v9"
    grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
    grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "io"
    stdLog "log"
    "net"
    "testing"
    "time"
)

type testServer struct {
    UnimplementedBatteryServiceServer
}

func (s *testServer) GetBatterySample(ctx context.Context, req *BatteryBatchQueryRequest) (res *BatterySampleResponse, err error) {
    // res = &BatterySampleResponse{Items: make([]*BatterySampleResponse_Battery, 0)}
    return
}

func TestBatteryService(t *testing.T) {
    var zaplogger *zap.Logger
    log.New(&log.Config{
        FormatJson:  true,
        Stdout:      true,
        Application: "testbattery",
        Writers: []io.Writer{
            log.NewRedisWriter(redis.NewClient(&redis.Options{})),
        },
    }, log.WithReplacer(func(logger *zap.Logger) {
        zaplogger = logger.WithOptions(zap.WithCaller(false), zap.IncreaseLevel(zapcore.WarnLevel)).Named("grpc-battery")
        grpczap.ReplaceGrpcLoggerV2(zaplogger)
    }))

    lis, err := net.Listen("tcp", ":8972")
    if err != nil {
        stdLog.Fatal(err)
    }
    s := grpc.NewServer(
        grpcmiddleware.WithUnaryServerChain(
            grpcctxtags.UnaryServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
            grpczap.UnaryServerInterceptor(zaplogger, adapter.GrpcZapDefaultOptions...),
        ),
    )
    RegisterBatteryServiceServer(s, &testServer{})
    stdLog.Fatal(s.Serve(lis))
}

func TestNewBatteryServiceClient(t *testing.T) {
    var zaplogger *zap.Logger
    log.New(&log.Config{
        FormatJson:  true,
        Stdout:      true,
        Application: "testbattery",
        Writers: []io.Writer{
            log.NewRedisWriter(redis.NewClient(&redis.Options{})),
        },
    }, log.WithReplacer(func(logger *zap.Logger) {
        zaplogger = logger.WithOptions(zap.WithCaller(false), zap.IncreaseLevel(zapcore.WarnLevel)).Named("grpc-battery")
        grpczap.ReplaceGrpcLoggerV2(zaplogger)
    }))

    conn, err := grpc.Dial("127.0.0.1:5031", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        stdLog.Fatal(err)
    }
    defer func(conn *grpc.ClientConn) {
        _ = conn.Close()
    }(conn)

    c := NewBatteryServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, _ := c.GetBatteryDetail(ctx, &BatteryQueryRequest{Id: 210})
    fmt.Println(res)
}
