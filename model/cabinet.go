// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package model

// Cabinet is the model entity for the Cabinet schema.
type Cabinet struct {
    ID uint64 `json:"id,omitempty"`
    // 是否在线
    Online bool `json:"online,omitempty"`
    // 品牌
    Brand string `json:"brand,omitempty"`
    // 电柜编号
    Serial string `json:"serial,omitempty"`
    // 状态
    Status string `json:"status,omitempty"`
    // 电柜是否启用
    Enable bool `json:"enable,omitempty"`
    // 经度
    Lng *float64 `json:"lng,omitempty"`
    // 纬度
    Lat *float64 `json:"lat,omitempty"`
    // GSM信号强度
    Gsm *float64 `json:"gsm,omitempty"`
    // 换电柜总电压 (V)
    Voltage *float64 `json:"voltage,omitempty"`
    // 换电柜总电流 (A)
    Current *float64 `json:"current,omitempty"`
    // 柜体温度值 (换电柜温度)
    Temperature *float64 `json:"temperature,omitempty"`
    // 总用电量
    Electricity *float64 `json:"electricity,omitempty"`
}
