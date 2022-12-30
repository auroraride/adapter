// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on adapter by liasica, magicrolan@qq.com.

package model

import (
    "database/sql/driver"
    "fmt"
)

type OperateType string

func (s *OperateType) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *s = OperateType(v)
    }
    return nil
}

func (s OperateType) Value() (driver.Value, error) {
    return s, nil
}

const (
    OperateTypeUnknown    OperateType = "unknown"
    OperateTypeBinOpen    OperateType = "binOpen"    // 仓位开门
    OperateTypeBinDisable OperateType = "binDisable" // 仓位禁用
    OperateTypeBinEnable  OperateType = "binEnable"  // 仓位启用
)

type OperateRequest struct {
    Type    OperateType `json:"type" validate:"required"`    // 控制类型
    Serial  string      `json:"serial" validate:"required"`  // 待控制的电柜编号
    Ordinal *int        `json:"ordinal" validate:"required"` // 待控制的仓位序号
}

func (o *OperateRequest) String() string {
    op := "-"
    switch o.Type {
    case OperateTypeBinOpen:
        op = "开仓"
    case OperateTypeBinDisable:
        op = "禁用"
    case OperateTypeBinEnable:
        op = "启用"
    }
    return fmt.Sprintf("[%s-%d] 仓控 %s", o.Serial, o.Ordinal, op)
}
