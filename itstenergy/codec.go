// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package itstenergy

import (
    "encoding/binary"
    "fmt"
    "github.com/auroraride/adapter"
)

type Code byte

const (
    CodeStop          Code = 0xD
    CodeStart         Code = 0x3A
    CodeLoginRecv     Code = 0xA1
    CodeLoginSent     Code = 0xA2
    CodeHeartbeatRecv Code = 0xB1
    CodeControlRecv   Code = 0xC1
    CodeControlSent   Code = 0xC2
)

func (c Code) Byte() byte {
    return byte(c)
}

func (c Code) ByteEqual(b byte) bool {
    return c.Byte() == b
}

func Pack() {
}

func Unpack(b []byte) (data []byte, err error) {
    // 校验起始符和结束符
    if !CodeStart.ByteEqual(b[0]) || !CodeStop.ByteEqual(b[len(b)-1]) {
        err = adapter.ErrorData
        return
    }

    // 获取IMEI
    imei := b[2:17]
    fmt.Println("IMEI:", string(imei))

    // 获取数据长度
    datalen := binary.BigEndian.Uint16(b[17:19]) + 19
    data = b[19:datalen]

    // 校验和

    sn := data[0:16]
    fmt.Println("电池包编码:", string(sn))

    sver := data[16:18]
    fmt.Println("BMS软件版本", binary.BigEndian.Uint16(sver))

    hver := data[18:20]
    fmt.Println("BMS硬件版本", binary.BigEndian.Uint16(hver))

    sver4g := data[20:22]
    fmt.Println("4G软件版本", binary.BigEndian.Uint16(sver4g))

    hver4g := data[22:24]
    fmt.Println("4G硬件版本", binary.BigEndian.Uint16(hver4g))

    sn4g := uint(binary.BigEndian.Uint32(data[24:28])) + 787986650000
    fmt.Println("4G板SN:", sn4g)

    iccd := data[28:48]
    fmt.Println("SIM卡ICCID:", string(iccd))

    soc := float64(binary.BigEndian.Uint16(data[48:50])) * 0.01
    fmt.Printf("电池设计容量: %.2fAH\n", soc)

    num := binary.BigEndian.Uint16(data[50:52])
    fmt.Println("电池包串数:", num)
    // 判断功能码
    fc := b[1]
    switch Code(fc) {
    case CodeLoginRecv:

    }

    return
}
