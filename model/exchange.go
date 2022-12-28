// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package model

import (
    "database/sql/driver"
    "time"
)

type ExchangeRequest struct {
    Serial string        `json:"serial" query:"serial" validate:"required"` // 电柜编号
    MinSoc float64       `json:"minSoc" query:"minSoc" validate:"required"` // 换电最小电量
    Lock   time.Duration `json:"lock" query:"lock" validate:"required"`     // 扫码锁定时间
}

type ExchangeUsableResponse struct {
    Cabinet Cabinet `json:"cabinet"`
    UUID    string  `json:"uuid,omitempty"`
    Fully   Bin     `json:"fully"` // 满电仓
    Empty   Bin     `json:"empty"` // 空电仓
}

// ExchangeStep 换电步骤
type ExchangeStep uint8

const (
    ExchangeStepFirst ExchangeStep = iota + 1
    ExchangeStepSecond
    ExchangeStepThird
    ExchangeStepFourth
)

func (s *ExchangeStep) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case int64:
        *s = ExchangeStep(v)
    }
    return nil
}

func (s ExchangeStep) Value() (driver.Value, error) {
    return s, nil
}

func (s ExchangeStep) String() string {
    switch s {
    case ExchangeStepFirst:
        return "第1步, 开启空电仓门"
    case ExchangeStepSecond:
        return "第2步, 放入电池关仓"
    case ExchangeStepThird:
        return "第3步, 开启满电仓门"
    case ExchangeStepFourth:
        return "第4步, 取出电池关仓"
    }
    return "-"
}
