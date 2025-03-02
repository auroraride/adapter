// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-10, by liasica

package es

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

const (
	DefaultFieldTimestamp = "@timestamp"
)

// CreateDocument 创建文档
func (e *Elastic) CreateDocument(doc any) {
	index := e.GetIndex()
	res, err := e.client.Index(index).Document(doc).Do(context.Background())
	if err != nil {
		b, _ := json.Marshal(doc)
		zap.L().Error("document创建失败", logTag(), zap.Error(err), zap.String("index", index), zap.Reflect("res", res), zap.ByteString("doc", b))
	}
}
