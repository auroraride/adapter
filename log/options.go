// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-03
// Based on adapter by liasica, magicrolan@qq.com.

package log

import "go.uber.org/zap"

type Option interface {
	apply(*zap.Logger)
}

type optionFunc func(*zap.Logger)

func (f optionFunc) apply(logger *zap.Logger) {
	f(logger)
}
