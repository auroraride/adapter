// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "fmt"
)

type Operator string

func (s *Operator) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *s = Operator(v)
    }
    return nil
}

func (s Operator) Value() (driver.Value, error) {
    return s, nil
}

const (
    OperatorUnknown    Operator = "unknown"
    OperatorBinOpen    Operator = "binOpen"    // 仓位开门
    OperatorBinDisable Operator = "binDisable" // 仓位禁用
    OperatorBinEnable  Operator = "binEnable"  // 仓位启用
)

type OperateRequest struct {
    Type    Operator `json:"type" validate:"required"`    // 控制类型
    Serial  string   `json:"serial" validate:"required"`  // 待控制的电柜编号
    Ordinal *int     `json:"ordinal" validate:"required"` // 待控制的仓位序号
}

func (o *OperateRequest) String() string {
    op := "-"
    switch o.Type {
    case OperatorBinOpen:
        op = "开仓"
    case OperatorBinDisable:
        op = "禁用"
    case OperatorBinEnable:
        op = "启用"
    }
    return fmt.Sprintf("[%s-%d] 仓控 %s", o.Serial, o.Ordinal, op)
}
