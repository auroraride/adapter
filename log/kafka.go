// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-02
// Based on adapter by liasica, magicrolan@qq.com.

package log

import (
	"github.com/IBM/sarama"
)

type KafkaWriter struct {
	Topic     string
	Producer  sarama.SyncProducer
	Partition int32
}

func NewKafkaWriter(pr *KafkaWriter) *KafkaWriter {
	return &KafkaWriter{
		Producer:  pr.Producer,
		Topic:     pr.Topic,
		Partition: pr.Partition,
	}
}

func (lk *KafkaWriter) Write(p []byte) (n int, err error) {
	// 构建消息
	msg := &sarama.ProducerMessage{
		Topic:     lk.Topic,
		Value:     sarama.ByteEncoder(p),
		Partition: lk.Partition,
	}

	// 发送消息
	_, _, err = lk.Producer.SendMessage(msg)
	if err != nil {
		return
	}

	return
}
