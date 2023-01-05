// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-03
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/google/uuid"
    "time"
)

// BusinuessUsableRequest 获取业务仓位请求
type BusinuessUsableRequest struct {
    Minsoc   float64          `json:"minsoc" validate:"required"` // 最小电量
    Business adapter.Business `json:"business" validate:"required"`
    Serial   string           `json:"serial" validate:"required"`
    Model    string           `json:"model" validate:"required"` // 电池型号
}

type BusinessRequest struct {
    UUID     uuid.UUID        `json:"uuid" validate:"required"`
    Business adapter.Business `json:"business" validate:"required"`                                                       // 业务类别
    Serial   string           `json:"serial" validate:"required"`                                                         // 电柜编号
    Timeout  int64            `json:"timeout" validate:"required"`                                                        // 超时时间(s)
    Battery  string           `json:"verifyBattery,omitempty" validate:"required_if=Business pause Business unsubscribe"` // 需要校验的电池编号 (可为空, 需要校验放入电池编号的时候必须携带, 例如putin操作)
    Model    string           `json:"model" validate:"required"`                                                          // 电池型号
}

func (req *BusinessRequest) String() string {
    return fmt.Sprintf(
        "[电柜: %s, 业务: %s, 电池校验: %s]",
        req.Serial,
        req.Business,
        adapter.Or(req.Battery == "", " - ", req.Battery),
    )
}

type BusinessResponse struct {
    Error   string              `json:"error,omitempty"`
    Results []*BinOperateResult `json:"results"`
}

type BinOperateResult struct {
    UUID      string           `json:"uuid"`
    Operate   Operate          `json:"operate"`
    Step      int              `json:"step"`                // 操作步骤
    Business  adapter.Business `json:"business"`            // 业务类型
    StartAt   *time.Time       `json:"startAt"`             // 开始时间
    StopAt    *time.Time       `json:"stopAt"`              // 结束时间
    Success   bool             `json:"success"`             // 是否成功
    Before    *BinInfo         `json:"before"`              // 操作前仓位信息
    After     *BinInfo         `json:"after"`               // 操作后仓位信息
    Duration  float64          `json:"duration,omitempty"`  // 耗时
    Message   string           `json:"message,omitempty"`   // 消息
    BatterySN string           `json:"batterySn,omitempty"` // 在位电池编号
}

type CabinetBinUsableResponse struct {
    Cabinet     *Cabinet `json:"cabinet"`
    UUID        string   `json:"uuid,omitempty"`
    Fully       *Bin     `json:"fully,omitempty"`       // 满电仓
    Empty       *Bin     `json:"empty,omitempty"`       // 空电仓
    BusinessBin *Bin     `json:"businessBin,omitempty"` // 业务仓位
}

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

type OperateBinRequest struct {
    Operate Operate `json:"operate" validate:"required"`
    Ordinal *int    `json:"ordinal" validate:"required"`
    Serial  string  `json:"serial" validate:"required"`
    Remark  string  `json:"remark" validate:"required"`
}
