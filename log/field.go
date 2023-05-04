// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
	"encoding/hex"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

func Binary(b []byte) zap.Field {
	return zap.Binary("binary", b)
}

func Payload(payload any) zap.Field {
	return zap.Reflect("payload", payload)
}

func JsonData(data any) zap.Field {
	b, _ := jsoniter.Marshal(data)
	return zap.ByteString("data", b)
}

func ResponseBody(b []byte) zap.Field {
	return zap.ByteString("response", b)
}

func Hex(b []byte) zap.Field {
	return zap.String("hex", hex.EncodeToString(b))
}
