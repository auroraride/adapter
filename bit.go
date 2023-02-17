// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

type PositiveNumber interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func GetTrueBits[T PositiveNumber](num T, max int) (bits []int) {
    bits = make([]int, 0)
    index := 0
    for i := num; i != 0; i >>= 1 {
        if max <= index {
            return
        }
        if i&1 == 1 {
            bits = append(bits, index)
        }
        index += 1
    }
    return
}
