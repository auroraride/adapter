// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package adapter

type CabinetStatus string

// CabinetStatus values.
const (
    StatusInitializing CabinetStatus = "initializing"
    StatusIdle         CabinetStatus = "idle"
    StatusBusy         CabinetStatus = "busy"
    StatusExchange     CabinetStatus = "exchange"
    StatusAbnormal     CabinetStatus = "abnormal"
)

// Cabinet is the model entity for the Cabinet schema.
type Cabinet struct {
    ID uint64 `json:"id"`
    // 是否在线
    Online bool `json:"online"`
    // 品牌
    Brand Brand `json:"brand"`
    // 电柜编号
    Serial string `json:"serial"`
    // 状态
    Status CabinetStatus `json:"status"`
    // 电柜是否启用
    Enable bool `json:"enable"`
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
