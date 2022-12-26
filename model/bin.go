// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package model

// Bin is the model entity for the Bin schema.
type Bin struct {
    ID uint64 `json:"id,omitempty"`
    // 唯一标识
    UUID string `json:"uuid,omitempty"`
    // CabinetID holds the value of the "cabinet_id" field.
    CabinetID uint64 `json:"cabinet_id,omitempty"`
    // 品牌
    Brand string `json:"brand,omitempty"`
    // 电柜设备序列号
    Serial string `json:"serial,omitempty"`
    // 仓位名称(N号仓)
    Name string `json:"name,omitempty"`
    // 仓位序号(从1开始)
    Ordinal int `json:"ordinal,omitempty"`
    // 仓门是否开启
    Open bool `json:"open,omitempty"`
    // 仓位是否启用
    Enable bool `json:"enable,omitempty"`
    // 仓位是否健康
    Health bool `json:"health,omitempty"`
    // 是否有电池
    BatteryExists bool `json:"battery_exists,omitempty"`
    // 电池序列号
    BatterySn string `json:"battery_sn,omitempty"`
    // 当前电压
    Voltage float64 `json:"voltage,omitempty"`
    // 当前电流
    Current float64 `json:"current,omitempty"`
    // 电池电量
    Soc float64 `json:"soc,omitempty"`
    // 电池健康程度
    Soh float64 `json:"soh,omitempty"`
    // 仓位备注
    Remark *string `json:"remark,omitempty"`
}
