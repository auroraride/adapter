// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package itstenergy

import (
    "encoding/base64"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUnpack(t *testing.T) {
    b, err := base64.StdEncoding.DecodeString("OqE4Njk1MjMwNTQyMTM4OTgANFhDQjExMTExMTExMTExMTEACgADAAYAAQAE4qI4OTg2MDQ5NDI2MjFEMDExOTExNAu4ABgN3w0=")
    assert.Nil(t, err)

    var data []byte
    data, err = Unpack(b)

    t.Logf("%s", string(data))
    assert.NotZero(t, data)
}
