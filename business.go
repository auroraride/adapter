// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "fmt"
    "github.com/google/uuid"
    "io"
    "strconv"
)

// Business 全部业务类型
type Business string

const (
    BusinessOperate     Business = "operate"     // 运维操作
    BusinessExchange    Business = "exchange"    // 换电
    BusinessActive      Business = "active"      // 激活
    BusinessPause       Business = "pause"       // 寄存
    BusinessContinue    Business = "continue"    // 取消寄存
    BusinessUnsubscribe Business = "unsubscribe" // 退订
)

var (
    RiderBusiness = []Business{BusinessActive, BusinessPause, BusinessContinue, BusinessUnsubscribe}
)

func (b Business) Text() string {
    switch b {
    case BusinessOperate:
        return "操作"
    case BusinessExchange:
        return "换电"
    case BusinessActive:
        return "激活"
    case BusinessPause:
        return "寄存"
    case BusinessContinue:
        return "取消寄存"
    case BusinessUnsubscribe:
        return "退订"
    }
    return " - "
}

// BatteryNeed 业务是否需要电池
func (b Business) BatteryNeed() bool {
    switch b {
    case BusinessPause, BusinessUnsubscribe:
        return true
    }
    return false
}

func (b Business) String() string {
    return string(b)
}

func (Business) Values() []string {
    return []string{
        BusinessOperate.String(),
        BusinessExchange.String(),
        BusinessActive.String(),
        BusinessPause.String(),
        BusinessContinue.String(),
        BusinessUnsubscribe.String(),
    }
}

func (b *Business) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *b = Business(v)
    }
    return nil
}

func (b Business) Value() (driver.Value, error) {
    return b, nil
}

func BusinessValidator(t Business) error {
    switch t {
    case BusinessExchange, BusinessOperate, BusinessActive, BusinessPause, BusinessContinue, BusinessUnsubscribe:
        return nil
    default:
        return fmt.Errorf("未知的业务类别: %q", t)
    }
}

// MarshalGQL implements graphql.Marshaler interface.
func (b Business) MarshalGQL(w io.Writer) {
    _, _ = io.WriteString(w, strconv.Quote(b.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (b *Business) UnmarshalGQL(val interface{}) error {
    str, ok := val.(string)
    if !ok {
        return fmt.Errorf("enum %T must be a string", val)
    }
    *b = Business(str)
    if err := BusinessValidator(*b); err != nil {
        return fmt.Errorf("%s is not a valid Business", str)
    }
    return nil
}

// BusinuessUsableRequest 获取业务仓位请求
type BusinuessUsableRequest struct {
    Minsoc   float64  `json:"minsoc" validate:"required"` // 最小电量
    Business Business `json:"business" validate:"required"`
    Serial   string   `json:"serial" validate:"required"`
    Model    string   `json:"model" validate:"required"` // 电池型号
}

type BusinessRequest struct {
    UUID     uuid.UUID `json:"uuid" validate:"required"`
    Business Business  `json:"business" validate:"required"`                                                       // 业务类别
    Serial   string    `json:"serial" validate:"required"`                                                         // 电柜编号
    Timeout  int64     `json:"timeout" validate:"required"`                                                        // 超时时间(s)
    Battery  string    `json:"verifyBattery,omitempty" validate:"required_if=Business pause Business unsubscribe"` // 需要校验的电池编号 (可为空, 需要校验放入电池编号的时候必须携带, 例如putin操作)
    Model    string    `json:"model" validate:"required"`                                                          // 电池型号
}

func (req *BusinessRequest) String() string {
    return fmt.Sprintf(
        "[电柜: %s, 业务: %s, 电池校验: %s]",
        req.Serial,
        req.Business,
        Or(req.Battery == "", " - ", req.Battery),
    )
}

type BusinessResponse struct {
    Error   string               `json:"error,omitempty"`
    Results []*OperateStepResult `json:"results"`
}
