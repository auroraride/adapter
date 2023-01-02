// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "github.com/google/uuid"

type ExchangeUsableRequest struct {
    Serial string  `json:"serial" validate:"required"` // 电柜编号
    Minsoc float64 `json:"minsoc" validate:"required"` // 换电最小电量
    Lock   int64   `json:"lock" validate:"required"`   // 扫码锁定时间
    Model  string  `json:"model" validate:"required"`  // 电池型号
}

type ExchangeRequest struct {
    UUID    uuid.UUID `json:"uuid" validate:"required"`
    Battery string    `json:"battery" validate:"required"` // 当前电池编号, 若放入电池型号不匹配, 则中断换电流程
    Expires int64     `json:"expires" validate:"required"` // 扫码有效期(s), 例如: 30s
    Timeout int64     `json:"timeout" validate:"required"` // 换电步骤超时(s), 例如: 120s
    Minsoc  float64   `json:"minsoc" validate:"required"`  // 换电最小电量
}

type ExchangeResponse struct {
    Success       bool                   `json:"success"`                 // 是否换电成功
    PutoutBattery string                 `json:"putoutBattery,omitempty"` // 取走电池编号
    PutinBattery  string                 `json:"putinBattery"`            // 放入电池编号
    Results       []*ExchangeStepMessage `json:"results,omitempty"`       // 步骤详情
    Error         string                 `json:"error,omitempty"`         // 错误消息
}
