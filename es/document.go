// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-10, by liasica

package es

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

const (
	DefaultFieldTimestamp = "@timestamp"
)

type Document interface {
	GetTimestamp() time.Time
}

// CreateDocument 创建文档
func (e *Elastic) CreateDocument(doc Document) {
	index := e.GetIndex()
	b, _ := jsoniter.Marshal(doc)
	a := make(map[string]any)
	_ = jsoniter.Unmarshal(b, &a)
	if _, ok := a[DefaultFieldTimestamp]; !ok {
		a[DefaultFieldTimestamp] = doc.GetTimestamp()
	}

	res, err := e.client.Index(index).Document(a).Do(context.Background())
	if err != nil {
		zap.L().Error("document创建失败", logTag(), zap.Error(err), zap.String("index", index), zap.Reflect("payload", res))
	}
}
