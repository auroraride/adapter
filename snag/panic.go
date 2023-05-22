// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package snag

import (
	"fmt"

	"go.uber.org/zap"
)

func WithRecover(cb func()) {

	defer func() {
		if v := recover(); v != nil {
			zap.L().Error(
				"捕获未处理崩溃",
				zap.Stack("stack"),
				zap.Error(fmt.Errorf("%v", v)),
			)
		}
	}()

	cb()
}

func WithRecoverError(cb func() error) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%v", v)
			zap.L().Error(
				"捕获未处理崩溃",
				zap.Stack("stack"),
				zap.Error(fmt.Errorf("%v", v)),
			)
		}
	}()

	err = cb()
	return
}
