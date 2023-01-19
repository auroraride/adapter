// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-18
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "reflect"
    "unsafe"
)

func ConvertBytes2String(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func ConvertString2Bytes(s string) (b []byte) {
    /* #nosec G103 */
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    /* #nosec G103 */
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh.Data = sh.Data
    bh.Cap = sh.Len
    bh.Len = sh.Len
    return
}
