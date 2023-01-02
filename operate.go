// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "time"
)

// DetectBattery 电池检测
type DetectBattery uint8

const (
    DetectBatteryIgnore DetectBattery = iota // 忽略
    DetectBatteryPutin                       // 检测是否放入
    DetectBatteryPutout                      // 检测是否取走
)

func (o DetectBattery) Text() string {
    switch o {
    default:
        return " - "
    case DetectBatteryIgnore:
        return "忽略"
    case DetectBatteryPutin:
        return "放入"
    case DetectBatteryPutout:
        return "取走"
    }
}

// DetectDoor 仓门检测
type DetectDoor uint8

const (
    DetectDoorIgnore DetectDoor = iota // 忽略
    DetectDoorOpen                     // 检测是否开门
    DetectDoorClose                    // 检测是否关门
)

func (o DetectDoor) Text() string {
    switch o {
    default:
        return " - "
    case DetectDoorIgnore:
        return "忽略"
    case DetectDoorOpen:
        return "开启"
    case DetectDoorClose:
        return "关闭"
    }
}

const (
    OperateUnknown    Operate = "unknown"     // 未知
    OperateDetect     Operate = "detect"      // 检测
    OperateDoorOpen   Operate = "door_open"   // 开仓
    OperateBinDisable Operate = "bin_disable" // 仓位禁用
    OperateBinEnable  Operate = "bin_enable"  // 仓位启用
)

type Operate string

func (o Operate) Text() string {
    switch o {
    default:
        return " - "
    case OperateDoorOpen:
        return "开仓"
    case OperateBinDisable:
        return "禁用仓位"
    case OperateBinEnable:
        return "启用仓位"
    }
}

func (o Operate) IsCommand() bool {
    switch o {
    default:
        return false
    case OperateDoorOpen, OperateBinDisable, OperateBinEnable:
        return true
    }
}

func (o *Operate) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *o = Operate(v)
    }
    return nil
}

func (o Operate) Value() (driver.Value, error) {
    return o, nil
}

type OperateStepResult struct {
    UUID      string     `json:"uuid"`
    Operate   Operate    `json:"operate"`
    Step      int        `json:"step"`                // 操作步骤
    Business  Business   `json:"business"`            // 业务类型
    StartAt   *time.Time `json:"startAt"`             // 开始时间
    StopAt    *time.Time `json:"stopAt"`              // 结束时间
    Success   bool       `json:"success"`             // 是否成功
    Before    *BinInfo   `json:"before"`              // 操作前仓位信息
    After     *BinInfo   `json:"after"`               // 操作后仓位信息
    Duration  float64    `json:"duration,omitempty"`  // 耗时
    Message   string     `json:"message,omitempty"`   // 消息
    BatterySN string     `json:"batterySn,omitempty"` // 在位电池编号
}

type OperateBinRequest struct {
    Operate Operate `json:"operate" validate:"required"`
    Ordinal *int    `json:"ordinal" validate:"required"`
    Serial  string  `json:"serial" validate:"required"`
}
