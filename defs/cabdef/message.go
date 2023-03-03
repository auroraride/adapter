// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import (
    jsoniter "github.com/json-iterator/go"
)

type ExchangeStepMessage BinOperateResult

func (m *ExchangeStepMessage) MarshalBinary() ([]byte, error) {
    return jsoniter.Marshal(m)
}

func (m *ExchangeStepMessage) UnmarshalBinary(data []byte) error {
    return jsoniter.Unmarshal(data, m)
}
