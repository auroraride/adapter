// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-30
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/auroraride/adapter"
)

type testdata struct {
	ID int `json:"id" mapstructure:"id"`
}

func TestRun(t *testing.T) {
	s := New[testdata](
		redis.NewClient(&redis.Options{}),
		adapter.Development,
		"TEST",
		func(data []*testdata) {
			fmt.Printf("%#v\n", data)
		},
	)
	go s.Run()

	i := 0
	ticker := time.NewTicker(10 * time.Second)
	for ; true; <-ticker.C {
		i += 1
		s.Push(&testdata{ID: i})
	}
}
