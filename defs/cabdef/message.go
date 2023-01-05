// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import (
    "github.com/auroraride/adapter"
    jsoniter "github.com/json-iterator/go"
)

type CabinetMessage struct {
    Full    bool     `json:"full"`
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}

func (m *CabinetMessage) MarshalBinary() ([]byte, error) {
    return jsoniter.Marshal(m)
}

func (m *CabinetMessage) UnmarshalBinary(data []byte) error {
    return jsoniter.Unmarshal(data, m)
}

type BatteryMessage struct {
    *adapter.Battery
    Cabinet string `json:"cabinet,omitempty"` // 所属电柜, 可能为空
}

func (m *BatteryMessage) MarshalBinary() ([]byte, error) {
    return jsoniter.Marshal(m)
}

func (m *BatteryMessage) UnmarshalBinary(data []byte) error {
    return jsoniter.Unmarshal(data, m)
}

type ExchangeStepMessage BinOperateResult

func (m *ExchangeStepMessage) MarshalBinary() ([]byte, error) {
    return jsoniter.Marshal(m)
}

func (m *ExchangeStepMessage) UnmarshalBinary(data []byte) error {
    return jsoniter.Unmarshal(data, m)
}
