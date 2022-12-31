// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/goccy/go-json"
    "time"
)

type Messenger interface {
    MarshalBinary() ([]byte, error)
    UnmarshalBinary(data []byte) error
}

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
    *Battery
    Cabinet string `json:"cabinet"` // 所属电柜
}

func (m *BatteryMessage) MarshalBinary() ([]byte, error) {
    return json.Marshal(m)
}

func (m *BatteryMessage) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}

type ExchangeStepMessage struct {
    UUID     string       `json:"uuid"`
    StartAt  *time.Time   `json:"startAt"`           // 开始时间
    StopAt   *time.Time   `json:"stopAt"`            // 结束时间
    Success  bool         `json:"success"`           // 是否成功
    Step     ExchangeStep `json:"step"`              // 步骤
    Before   *BinInfo     `json:"before"`            // 操作前仓位信息
    After    *BinInfo     `json:"after"`             // 操作后仓位信息
    Duration float64      `json:"duration"`          // 耗时
    Message  string       `json:"message,omitempty"` // 消息
}

func (m *ExchangeStepMessage) MarshalBinary() ([]byte, error) {
    return json.Marshal(m)
}

func (m *ExchangeStepMessage) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}
