// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-17
// Based on adapter by liasica, magicrolan@qq.com.

package mq

type Message struct {
    Topic    string
    Payload  any
    Qos      byte
    Retained bool
}
