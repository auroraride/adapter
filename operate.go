// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "fmt"
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

const (
    OperateUnknown    Operate = "unknown"
    OperateBinDetect  Operate = "detect"      // 检测
    OperateBinOpen    Operate = "bin_open"    // 开仓
    OperateBinDisable Operate = "bin_disable" // 仓位禁用
    OperateBinEnable  Operate = "bin_enable"  // 仓位启用
)

type OperateRequest struct {
    Type    Operate `json:"type" validate:"required"`    // 控制类型
    Serial  string  `json:"serial" validate:"required"`  // 待控制的电柜编号
    Ordinal *int    `json:"ordinal" validate:"required"` // 待控制的仓位序号
}

func (o *OperateRequest) String() string {
    op := "-"
    switch o.Type {
    case OperateBinOpen:
        op = "开仓"
    case OperateBinDetect:
        op = "仓位检测"
    case OperateBinDisable:
        op = "仓位禁用"
    case OperateBinEnable:
        op = "仓位启用"
    }
    return fmt.Sprintf("[%s-%d] 操作 %s", o.Serial, o.Ordinal, op)
}

// BusinessOperate 业务操作
type BusinessOperate string

const (
    BusinessOperatePutin  BusinessOperate = "putin"  // 电池放入
    BusinessOperatePutout BusinessOperate = "putout" // 电池取出
)

// BusinessOperateRequest 业务操作请求
type BusinessOperateRequest struct {
    Serial  string          `json:"serial" validate:"required"`  // 电柜编号
    Operate BusinessOperate `json:"operate" validate:"required"` // 操作
    Battery string          `json:"battery,omitempty"`           // 电池编号
    Verify  bool            `json:"verify"`                      // 是否校验电池
}

type BusinessOperateResponse struct {
}
