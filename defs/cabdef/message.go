// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import (
    "github.com/auroraride/adapter"
    "github.com/goccy/go-json"
)

type CabinetMessage struct {
    Full    bool     `json:"full"`
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}

func (m *CabinetMessage) MarshalBinary() ([]byte, error) {
    return json.Marshal(m)
}

func (m *CabinetMessage) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}

type BatteryMessage struct {
    *adapter.Battery
    Cabinet string `json:"cabinet"` // 所属电柜
}

func (m *BatteryMessage) MarshalBinary() ([]byte, error) {
    return json.Marshal(m)
}

func (m *BatteryMessage) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}

type ExchangeStepMessage BusinessStepResult

func (m *ExchangeStepMessage) MarshalBinary() ([]byte, error) {
    return json.Marshal(m)
}

func (m *ExchangeStepMessage) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}
