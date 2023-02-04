// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-03
// Based on adapter by liasica, magicrolan@qq.com.

package proto

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative battery.proto
//go:generate go generate ./entpb
