// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-10, by liasica

package es

import (
	"context"

	"go.uber.org/zap"
)

const (
	DefaultFieldTimestamp = "@timestamp"
)

// CreateDocument 创建文档
func (e *Elastic) CreateDocument(document any) {
	index := e.GetIndex()
	res, err := e.client.Index(index).Document(document).Do(context.Background())
	if err != nil {
		zap.L().Error("document创建失败", logTag(), zap.Error(err), zap.String("index", index), zap.Reflect("payload", res))
	}
}
