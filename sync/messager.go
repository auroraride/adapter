// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-30
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/auroraride/adapter"
)

func Marshal[T any](key string, input T) (output map[string]any, b []byte, err error) {
	b, _ = jsoniter.Marshal(input)
	output = map[string]any{
		key: b,
	}
	return
}

func Unmarshal[T any](key string, input map[string]any) (output *T, err error) {
	data, ok := input[key]
	if !ok {
		err = adapter.ErrorData
		return
	}
	output = new(T)
	b := adapter.ConvertString2Bytes(data.(string))
	err = jsoniter.Unmarshal(b, output)
	return
}
