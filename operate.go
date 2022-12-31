// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "fmt"
    "github.com/google/uuid"
    "time"
)

const (
    OperateUnknown    Operate = "unknown"
    OperateDoorOpen   Operate = "door_open"   // 开仓
    OperateBinDisable Operate = "bin_disable" // 仓位禁用
    OperateBinEnable  Operate = "bin_enable"  // 仓位启用
    OperatePutin      Operate = "putin"       // 检测电池放入
    OperatePutout     Operate = "putout"      // 检测电池取出
)

type Operate string

func (s *Operate) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *s = Operate(v)
    }
    return nil
}

func (s Operate) Value() (driver.Value, error) {
    return s, nil
}

// OperateRequest 业务操作
type OperateRequest struct {
    Serial  string  `json:"serial" validate:"required"`  // 电柜编号
    Operate Operate `json:"operate" validate:"required"` // 操作类别
    Timeout int64   `json:"timeout" validate:"required"` // 超时时间(s)

    UUID               uuid.UUID     `json:"UUID"`
    Ordinal            *int          `json:"ordinal"`                      // 仓位序号 (操作电柜的时候为空, 操作仓位的时候必不为空)
    Step               *ExchangeStep `json:"step,omitempty"`               // 换电步骤 (可为空)
    VerifyPutinBattery string        `json:"verifyPutinBattery,omitempty"` // 需要校验的电池编号 (可为空, 需要校验放入电池编号的时候必须携带, 例如putin操作)
}

func (b OperateRequest) String() (str string) {
    bat := " - "
    if b.VerifyPutinBattery != "" {
        bat = b.VerifyPutinBattery
    }

    str = fmt.Sprintf(
        "[UUID: %s, 电柜: %s, 仓位: %d, 操作: %s, 换电步骤: %s, 电池校验: %s]",
        b.UUID.String(),
        b.Serial,
        b.Ordinal,
        b.Operate,
        b.Step.String(),
        bat,
    )

    return
}

type OperateResult struct {
    UUID     string     `json:"uuid"`
    StartAt  *time.Time `json:"startAt"`           // 开始时间
    StopAt   *time.Time `json:"stopAt"`            // 结束时间
    Success  bool       `json:"success"`           // 是否成功
    Before   *BinInfo   `json:"before"`            // 操作前仓位信息
    After    *BinInfo   `json:"after"`             // 操作后仓位信息
    Duration float64    `json:"duration"`          // 耗时
    Message  string     `json:"message,omitempty"` // 消息
}
