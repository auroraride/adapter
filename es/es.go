// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-10, by liasica

package es

import (
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type Elastic struct {
	client     *elasticsearch.TypedClient
	datastream string
}

var instance *Elastic

func logTag() zap.Field {
	return zap.String("tag", "ELASTIC")
}

// Create 创建ES实例
func Create(apiKey, datastream string, addresses []string) error {
	c, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		APIKey:    apiKey,
		Addresses: addresses,
	})
	if err != nil {
		return err
	}
	instance = &Elastic{
		client:     c,
		datastream: datastream,
	}
	return nil
}

func GetInstance() *Elastic {
	return instance
}

func (e *Elastic) GetClient() *elasticsearch.TypedClient {
	return e.client
}

func (e *Elastic) GetIndex() string {
	return e.datastream + time.Now().Format("2006.01.02")
}

func (e *Elastic) GetIndexWizard() string {
	return e.datastream + "*"
}
