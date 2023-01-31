// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import (
    "github.com/auroraride/adapter"
)

type CabinetMessage struct {
    Full    bool     `json:"full"`
    Serial  string   `json:"serial"`
    Cabinet *Cabinet `json:"cabinet,omitempty"`
    Bins    []*Bin   `json:"bins,omitempty"`
}

type BatteryMessage struct {
    *adapter.Battery
    Cabinet string `json:"cabinet,omitempty"` // 所属电柜, 可能为空
}

type ExchangeStepMessage BinOperateResult
