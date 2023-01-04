// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package message

import (
    "encoding/binary"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    jsoniter "github.com/json-iterator/go"
)

const (
    typeSize = 4
)

type DataType uint32

const (
    TypeCabkitSync         DataType = 1 + iota // 电柜同步消息
    TypeCabkitBattery                          // 电柜电池同步消息
    TypeCabkitExchangeStep                     // 电柜换电任务消息
)

func Unpack(b []byte) (t DataType, data any, err error) {
    if len(b) < typeSize {
        err = adapter.ErrorIncompletePacket
        return
    }

    t = DataType(binary.BigEndian.Uint32(b[:typeSize]))
    switch t {
    case TypeCabkitSync:
        data = new(cabdef.CabinetMessage)
    case TypeCabkitBattery:
        data = new(cabdef.BatteryMessage)
    case TypeCabkitExchangeStep:
        data = new(cabdef.ExchangeStepMessage)
    }

    err = jsoniter.Unmarshal(b[typeSize:], data)
    return
}

func Pack(m Messenger) (b []byte, err error) {
    var t DataType
    switch m.(type) {
    default:
        err = adapter.ErrorData
        return
    case *cabdef.CabinetMessage:
        t = TypeCabkitSync
    case *cabdef.BatteryMessage:
        t = TypeCabkitBattery
    case *cabdef.ExchangeStepMessage:
        t = TypeCabkitExchangeStep
    }

    var message []byte
    message, _ = m.MarshalBinary()
    msgLen := typeSize + len(message)

    b = make([]byte, msgLen)
    binary.BigEndian.PutUint32(b[:typeSize], uint32(t))

    copy(b[typeSize:msgLen], message)
    return
}
