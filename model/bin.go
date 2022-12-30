// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package model

// Bin is the model entity for the Bin schema.
type Bin struct {
    ID uint64 `json:"id"`
    // 品牌
    Brand string `json:"brand"`
    // 电柜设备序列号
    Serial string `json:"serial"`
    // 仓位名称(N号仓)
    Name string `json:"name"`
    // 仓位序号(从1开始)
    Ordinal int `json:"ordinal"`
    // 仓门是否开启
    Open bool `json:"open"`
    // 仓位是否启用
    Enable bool `json:"enable"`
    // 仓位是否健康
    Health bool `json:"health"`
    // 是否有电池
    BatteryExists bool `json:"battery_exists"`
    // 电池序列号
    BatterySn string `json:"battery_sn"`
    // 当前电压
    Voltage float64 `json:"voltage"`
    // 当前电流
    Current float64 `json:"current"`
    // 电池电量
    Soc float64 `json:"soc"`
    // 电池健康程度
    Soh float64 `json:"soh"`
    // 仓位备注
    Remark *string `json:"remark,omitempty"`
}

type BinInfo struct {
    Ordinal       int     `json:"ordinal"`        // 仓位序号
    BatterySN     string  `json:"batterySn"`      // 电池编码
    Voltage       float64 `json:"voltage"`        // 电压
    Current       float64 `json:"current"`        // 电流
    Soc           float64 `json:"soc"`            // 电量
    Soh           float64 `json:"soh"`            // 健康
    Health        bool    `json:"health"`         // 健康
    Enable        bool    `json:"enable"`         // 启用
    Open          bool    `json:"open"`           // 开启
    BatteryExists bool    `json:"battery_exists"` // 电池在位
}

func (b *Bin) Info() *BinInfo {
    return &BinInfo{
        Ordinal:       b.Ordinal,
        BatterySN:     b.BatterySn,
        Voltage:       b.Voltage,
        Current:       b.Current,
        Soc:           b.Soc,
        Soh:           b.Soh,
        Health:        b.Health,
        Enable:        b.Enable,
        Open:          b.Open,
        BatteryExists: b.BatteryExists,
    }
}

type BinOperateEnable struct {
    Enable  *bool   `json:"enable" validate:"required"`  // 是否启用
    Serial  *string `json:"serial" validate:"required"`  // 电柜编号
    Ordinal *int    `json:"ordinal" validate:"required"` // 仓位序号
}
