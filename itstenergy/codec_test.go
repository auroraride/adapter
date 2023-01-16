// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package itstenergy

import (
    "encoding/base64"
    "encoding/binary"
    "encoding/hex"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPack(t *testing.T) {
    b, err := hex.DecodeString(`3A1670584342303836323032323131303037301E0A0002006400000C8C0C8600390030000000060000010000300031003000310C880C860C870C860C870C890C890C890C880C870C870C870C860C870C870C890C880C870C860C870C890C890C880C8C0BFF05DC5630001800010E420D350AAB2ADB0D`)
    assert.Nil(t, err)
    data := b[3 : len(b)-3]
    alert := data[36:40]
    num := binary.BigEndian.Uint32(alert)

    alertIndexes := []int{0, 1, 2, 3, 6, 7, 8, 11, 12, 13, 14, 15, 16}
    alerts := map[int]string{
        0:  "总压低",
        1:  "总压高",
        2:  "单体低",
        3:  "单体高",
        6:  "放电过流",
        7:  "充电过流",
        8:  "SOC低",
        11: "充电高温",
        12: "充电低温",
        13: "放电高温",
        14: "放电低温",
        15: "短路",
        16: "MOS高温",
    }
    for _, i := range alertIndexes {
        fmt.Printf("index: %d, %s -> %t\n", i, alerts[i], num>>i&1 == 1)
    }

    // for i := num; i != 0; i >>= 1 {
    //     bit := i & 1
    //     fmt.Println(i, bit)
    // }
}

func TestUnpack(t *testing.T) {
    b, err := base64.StdEncoding.DecodeString("OqE4Njk1MjMwNTQyMTM4OTgANFhDQjExMTExMTExMTExMTEACgADAAYAAQAE4qI4OTg2MDQ5NDI2MjFEMDExOTExNAu4ABgN3w0=")
    assert.Nil(t, err)

    var data []byte
    data, err = Unpack(b)

    t.Logf("%s", string(data))
    assert.NotZero(t, data)
}
