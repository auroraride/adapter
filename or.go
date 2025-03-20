// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

func Or[T any](condition bool, yes T, no T) T {
	if condition {
		return yes
	}
	return no
}

func OrFunc[T any](condition func() bool, yes func() T, no func() T) T {
	if condition() {
		return yes()
	}
	return no()
}
