// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "encoding/binary"
    "github.com/goccy/go-json"
)

const (
    typeSize = 4
)

type DataType uint32

const (
    TypeCabinet DataType = 1 + iota
    TypeBattery
    TypeExchangeStep
)

func Unpack(b []byte) (t DataType, data any, err error) {
    if len(b) < typeSize {
        err = ErrorIncompletePacket
        return
    }

    t = DataType(binary.BigEndian.Uint32(b[:typeSize]))
    switch t {
    case TypeCabinet:
        data = new(CabinetMessage)
    case TypeBattery:
        data = new(BatteryMessage)
    case TypeExchangeStep:
        data = new(ExchangeStepMessage)
    }

    err = json.Unmarshal(b[typeSize:], data)
    return
}

func Pack(m Messenger) (b []byte, err error) {
    var t DataType
    switch m.(type) {
    default:
        err = ErrorData
        return
    case *CabinetMessage:
        t = TypeCabinet
    case *BatteryMessage:
        t = TypeBattery
    case *ExchangeStepMessage:
        t = TypeExchangeStep
    }

    var message []byte
    message, _ = m.MarshalBinary()
    msgLen := typeSize + len(message)

    b = make([]byte, msgLen)
    binary.BigEndian.PutUint32(b[:typeSize], uint32(t))

    copy(b[typeSize:msgLen], message)
    return
}
