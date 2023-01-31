// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-30
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
    "github.com/auroraride/adapter"
    jsoniter "github.com/json-iterator/go"
)

func Marshal[T any](key string, input T) (output map[string]any, err error) {
    b, _ := jsoniter.Marshal(input)
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
    err = jsoniter.Unmarshal(adapter.ConvertString2Bytes(data.(string)), output)
    return
}
