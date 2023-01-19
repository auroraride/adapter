// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-15
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "bytes"
    "sync"
)

var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func NewBuffer() *bytes.Buffer {
    return bufferPool.Get().(*bytes.Buffer)
}

func ReleaseBuffer(buf *bytes.Buffer) {
    // See https://golang.org/issue/23199
    const maxSize = 1 << 16
    if buf != nil && buf.Cap() < maxSize {
        buf.Reset()
        bufferPool.Put(buf)
    }
}
