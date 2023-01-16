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
    numb := fmt.Sprintf("%032b", num)
    t.Log(b, data, alert, num, numb)
}

func TestUnpack(t *testing.T) {
    b, err := base64.StdEncoding.DecodeString("OqE4Njk1MjMwNTQyMTM4OTgANFhDQjExMTExMTExMTExMTEACgADAAYAAQAE4qI4OTg2MDQ5NDI2MjFEMDExOTExNAu4ABgN3w0=")
    assert.Nil(t, err)

    var data []byte
    data, err = Unpack(b)

    t.Logf("%s", string(data))
    assert.NotZero(t, data)
}
