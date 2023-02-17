// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-17
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "bytes"
    "github.com/stretchr/testify/require"
    "reflect"
    "strings"
    "testing"
    "unsafe"
)

var L = 1024 * 1024
var str = strings.Repeat("a", L)
var s = bytes.Repeat([]byte{'a'}, L)

func BenchmarkConvertString2Bytes(b *testing.B) {
    for i := 0; i < b.N; i++ {
        bb := *(*[]byte)(unsafe.Pointer(&str))
        if len(bb) != L {
            b.Fatal()
        }
    }
}

func BenchmarkConvertString2Bytes2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        /* #nosec G103 */
        var bb []byte
        bh := (*reflect.SliceHeader)(unsafe.Pointer(&bb))
        /* #nosec G103 */
        sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
        bh.Data = sh.Data
        bh.Cap = sh.Len
        bh.Len = sh.Len
        if len(bb) != L {
            b.Fatal()
        }
    }
}

func TestConvertString2Bytes(t *testing.T) {
    teststr := "teststr"
    b := *(*[]byte)(unsafe.Pointer(&teststr))
    newstr := *(*string)(unsafe.Pointer(&b))
    require.Equal(t, teststr, newstr)
}
