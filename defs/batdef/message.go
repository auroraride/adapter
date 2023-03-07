// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-06
// Based on adapter by liasica, magicrolan@qq.com.

package batdef

import "github.com/auroraride/adapter"

type BatteryFlow struct {
    Out     *adapter.Battery  `json:"out"`     // 取出电池信息
    In      *adapter.Battery  `json:"in"`      // 放入电池信息
    Serial  string            `json:"serial"`  // 电柜编码
    Ordinal int               `json:"ordinal"` // 仓位序号
    Geom    *adapter.Geometry `json:"geom"`    // 坐标
}
