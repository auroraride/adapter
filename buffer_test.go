// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/stretchr/testify/assert"
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
