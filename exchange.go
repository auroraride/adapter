// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "github.com/google/uuid"

// // ExchangeStep 换电步骤
// type ExchangeStep uint8
//
// const (
//     ExchangeStepFirst ExchangeStep = iota + 1
//     ExchangeStepSecond
//     ExchangeStepThird
//     ExchangeStepFourth
// )
//
// func (s ExchangeStep) Index() int {
//     return int(s) - 1
// }
//
// func (s *ExchangeStep) Scan(src interface{}) error {
//     switch v := src.(type) {
//     case nil:
//         return nil
//     case int64:
//         *s = ExchangeStep(v)
//     }
//     return nil
// }
//
// func (s ExchangeStep) Value() (driver.Value, error) {
//     return s, nil
// }
//
// func (s ExchangeStep) String() string {
//     switch s {
//     case ExchangeStepFirst:
//         return "第1步, 开启空电仓门"
//     case ExchangeStepSecond:
//         return "第2步, 放入电池关仓"
//     case ExchangeStepThird:
//         return "第3步, 开启满电仓门"
//     case ExchangeStepFourth:
//         return "第4步, 取出电池关仓"
//     }
//     return "-"
// }
//
// var ExchangeSteps = []ExchangeStep{
//     ExchangeStepFirst,
//     ExchangeStepSecond,
//     ExchangeStepThird,
//     ExchangeStepFourth,
// }
//
// type ExchangeStepConfigure struct {
//     Step    ExchangeStep
//     Operate Operate
// }
//
// var ExchangeStepConfigures = []ExchangeStepConfigure{
//     // {
//     //     Step:    ExchangeStepFirst,
//     //     Operate: OperateDoorOpen,
//     // },
//     // {
//     //     Step:    ExchangeStepSecond,
//     //     Operate: OperatePutin,
//     // },
//     // {
//     //     Step:    ExchangeStepThird,
//     //     Operate: OperateDoorOpen,
//     // },
//     // {
//     //     Step:    ExchangeStepFourth,
//     //     Operate: OperatePutout,
//     // },
// }

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
