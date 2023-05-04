// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-25
// Based on adapter by liasica, magicrolan@qq.com.

package app

var (
	Quit chan bool
)

func init() {
	Quit = make(chan bool)
}
