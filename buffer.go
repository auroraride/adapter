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

func CheckSum(data []byte) uint32 {
    var (
        sum    uint32
        index  int
        length = len(data)
    )
    // 以每16位为单位进行求和，直到所有的字节全部求完或者只剩下一个8位字节（如果剩余一个8位字节说明字节数为奇数个）
    for length > 1 {
        sum += uint32(data[index])<<8 + uint32(data[index+1])
        index += 2
        length -= 2
    }
    // 如果字节数为奇数个，要加上最后剩下的那个8位字节
    if length > 0 {
        sum += uint32(data[index])
    }
    // 加上高16位进位的部分
    sum += sum >> 16
    // 别忘了返回的时候先求反
    return ^sum
}

// CheckSum2 校验和
func CheckSum2(data []byte) uint32 {
    var (
        sum    uint32
        length = len(data)
    )

    for i := 0; i < length; i++ {
        sum += uint32(data[i])
        if sum > 0xff {
            sum = ^sum
            sum += 1
        }
    }
    sum &= 0xff
    return sum
}
