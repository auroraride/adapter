// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on aurtcp by liasica, magicrolan@qq.com.

package cabdef

import "github.com/auroraride/adapter"

type CabinetStatus string

// CabinetStatus values.
const (
	StatusInitializing CabinetStatus = "initializing"
	StatusIdle         CabinetStatus = "idle"
	StatusBusy         CabinetStatus = "busy"
	StatusExchange     CabinetStatus = "exchange"
	StatusAbnormal     CabinetStatus = "abnormal"
)

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Cabinet is the model entity for the Cabinet schema.
type Cabinet struct {
	ID uint64 `json:"id,omitempty"`
	// 是否在线
	Online bool `json:"online,omitempty"`
	// 品牌
	Brand adapter.CabinetBrand `json:"brand,omitempty"`
	// 电柜编号
	Serial string `json:"serial,omitempty"`
	// 状态
	Status CabinetStatus `json:"status,omitempty"`
	// 电柜是否启用
	Enable bool `json:"enable,omitempty"`
	// 经度
	Lng *float64 `json:"lng,omitempty,omitempty"`
	// 纬度
	Lat *float64 `json:"lat,omitempty,omitempty"`
	// GSM信号强度
	Gsm *float64 `json:"gsm,omitempty,omitempty"`
	// 换电柜总电压 (V)
	Voltage *float64 `json:"voltage,omitempty,omitempty"`
	// 换电柜总电流 (A)
	Current *float64 `json:"current,omitempty,omitempty"`
	// 柜体温度值 (换电柜温度)
	Temperature *float64 `json:"temperature,omitempty,omitempty"`
	// 总用电量
	Electricity *float64 `json:"electricity,omitempty,omitempty"`
	// 电柜位置
	Location *GeoPoint `json:"location,omitempty,omitempty"`
}
