// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package model

import (
    "database/sql/driver"
    "time"
)

// ExchangeStep 换电步骤
type ExchangeStep uint8

const (
    ExchangeStepFirst ExchangeStep = iota + 1
    ExchangeStepSecond
    ExchangeStepThird
    ExchangeStepFourth
)

func (s ExchangeStep) Index() int {
    return int(s) - 1
}

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

var ExchangeSteps = []ExchangeStep{
    ExchangeStepFirst,
    ExchangeStepSecond,
    ExchangeStepThird,
    ExchangeStepFourth,
}

// DetectDoor 仓门检测
type DetectDoor uint8

const (
    DetectDoorIgnore DetectDoor = iota // 忽略
    DetectDoorOpen                     // 开仓
    DetectDoorClose                    // 关仓
)

func (d DetectDoor) String() string {
    switch d {
    default:
        return "忽略"
    case DetectDoorOpen:
        return "开仓"
    case DetectDoorClose:
        return "关仓"
    }
}

// DetectBattery 检测电池
type DetectBattery uint8

const (
    DetectBatteryIgnore DetectBattery = iota // 忽略
    DetectBatteryPutin                       // 放入
    DetectBatteryPutout                      // 取走
)

func (d DetectBattery) String() string {
    switch d {
    default:
        return "忽略"
    case DetectBatteryPutin:
        return "放入"
    case DetectBatteryPutout:
        return "取走"
    }
}

type ExchangeStepConfigure struct {
    Step    ExchangeStep
    Door    DetectDoor
    Battery DetectBattery
}

var ExchangeStepConfigures = []ExchangeStepConfigure{
    {
        Step:    ExchangeStepFirst,
        Battery: DetectBatteryIgnore,
        Door:    DetectDoorOpen,
    },
    {
        Step:    ExchangeStepSecond,
        Battery: DetectBatteryPutin,
        Door:    DetectDoorClose,
    },
    {
        Step:    ExchangeStepThird,
        Battery: DetectBatteryIgnore,
        Door:    DetectDoorOpen,
    },
    {
        Step:    ExchangeStepFourth,
        Battery: DetectBatteryPutout,
        Door:    DetectDoorClose,
    },
}

type ExchangeUsableRequest struct {
    Serial string  `json:"serial" validate:"required"` // 电柜编号
    Minsoc float64 `json:"minsoc" validate:"required"` // 换电最小电量
    Lock   int64   `json:"lock" validate:"required"`   // 扫码锁定时间
}

type ExchangeUsableResponse struct {
    Cabinet *Cabinet `json:"cabinet"`
    UUID    string   `json:"uuid,omitempty"`
    Fully   *Bin     `json:"fully"` // 满电仓
    Empty   *Bin     `json:"empty"` // 空电仓
}

type ExchangeRequest struct {
    UUID    string  `json:"uuid" validate:"required"`
    Expires int64   `json:"expires" validate:"required"` // 扫码有效期(s), 例如: 30s
    TimeOut int64   `json:"timeOut" validate:"required"` // 换电步骤超时(s), 例如: 120s
    Minsoc  float64 `json:"minsoc" validate:"required"`  // 换电最小电量
}

type ExchangeStepResult struct {
    StartAt  *time.Time    `json:"startAt"`  // 开始时间
    StopAt   *time.Time    `json:"stopAt"`   // 结束时间
    Success  bool          `json:"success"`  // 是否成功
    Step     *ExchangeStep `json:"step"`     // 步骤
    Before   *BinInfo      `json:"before"`   // 操作前仓位信息
    After    *BinInfo      `json:"after"`    // 操作后仓位信息
    Duration *float64      `json:"duration"` // 耗时
}

type ExchangeResponse struct {
    Results []*ExchangeStepResult `json:"results,omitempty"` // 步骤详情
}
