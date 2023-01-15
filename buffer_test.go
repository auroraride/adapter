// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/stretchr/testify/assert"
    "hash/crc32"
    "testing"
)

func TestBuffer(t *testing.T) {
    buf := NewBuffer()
    defer ReleaseBuffer(buf)

    assert.Zero(t, buf.Len())
    buf.Write([]byte("1"))
    assert.NotZero(t, buf.Len())
    ReleaseBuffer(buf)

    assert.Zero(t, NewBuffer().Len())
}

func TestCheckSum(t *testing.T) {
    data := []byte("hello world")
    s1 := CheckSum(data)
    s2 := CheckSum2(data)
    s3 := crc32.ChecksumIEEE(data)
    t.Logf("%x, %x, %x", s1, s2, s3)
}
