// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigure(t *testing.T) {
	type Config struct {
		Configure
		Version int
		Debug   bool
		Message string
	}
	dc := &Config{
		Version: 1,
		Debug:   true,
		Message: "Default Config",
	}

	dcb, err := jsoniter.Marshal(dc)
	assert.NoError(t, err)

	cfg := new(Config)
	err = LoadConfigure[*Config](cfg, "/tmp/adapter_config_temp.yaml", dcb)
	assert.NoError(t, err)

	assert.Equal(t, dc, cfg)

	cfg.Version += 1
	assert.NotEqual(t, dc, cfg)
}
