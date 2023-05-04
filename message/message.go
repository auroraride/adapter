// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-03
// Based on adapter by liasica, magicrolan@qq.com.

package message

type Messenger interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}
