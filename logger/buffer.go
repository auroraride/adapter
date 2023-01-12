package logger

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

func PutBuffer(buf *bytes.Buffer) {
    // See https://golang.org/issue/23199
    const maxSize = 1 << 16
    if buf != nil && buf.Cap() < maxSize {
        buf.Reset()
        bufferPool.Put(buf)
    }
}
