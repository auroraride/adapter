// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"os"
)

// CreateDirectoryIfNotExist 若目录不存在则创建
func CreateDirectoryIfNotExist(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		return os.MkdirAll(d, 0755)
	}
	return nil
}
