// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-18
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"unsafe"
)

func ConvertBytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ConvertString2Bytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
}
