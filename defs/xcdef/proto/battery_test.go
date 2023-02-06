// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-03
// Based on adapter by liasica, magicrolan@qq.com.

package proto

import (
    "context"
    "fmt"
    "github.com/auroraride/adapter/defs/xcdef/proto/entpb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    stdLog "log"
    "net"
    "testing"
    "time"
)

type testServer struct {
    UnimplementedBatteryServer
}

func (s *testServer) GetBatterySample(ctx context.Context, req *BatteryBatchQueryRequest) (res *BatterySampleResponse, err error) {
    res = &BatterySampleResponse{Items: make([]*entpb.Heartbeat, 0)}
    return
}

func (s *testServer) GetBatteryDetail(context.Context, *BatteryQueryRequest) (bat *entpb.Battery, err error) {
    // res = &BatterySampleResponse{Items: make([]*BatterySampleResponse_Battery, 0)}
    bat = &entpb.Battery{Id: 1}
    return
}

func TestBattery(t *testing.T) {
    lis, err := net.Listen("tcp", ":8972")
    if err != nil {
        stdLog.Fatal(err)
    }
    s := grpc.NewServer()
    RegisterBatteryServer(s, &testServer{})
    stdLog.Fatal(s.Serve(lis))
}

func TestNewBatteryClient(t *testing.T) {
    conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        stdLog.Fatal(err)
    }
    defer func(conn *grpc.ClientConn) {
        _ = conn.Close()
    }(conn)

    c := NewBatteryClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    n := 0
    for {
        if n >= 10 {
            break
        }
        n += 1
        info, _ := c.GetBatterySample(ctx, &BatteryBatchQueryRequest{Sn: []string{"B10D787986981494"}})
        fmt.Println("info", info)
        time.Sleep(1 * time.Second)
    }

    detail, _ := c.GetBatteryDetail(ctx, &BatteryQueryRequest{Sn: "B10D787986981494"})
    fmt.Println("detail", detail)
}
