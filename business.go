// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "fmt"
    "io"
    "strconv"
)

type Business string

const (
    BusinessOperate     Business = "operate"     // 运维操作
    BusinessExchange    Business = "exchange"    // 换电
    BusinessActive      Business = "active"      // 激活
    BusinessPause       Business = "pause"       // 寄存
    BusinessContinue    Business = "continue"    // 取消寄存
    BusinessUnsubscribe Business = "unsubscribe" // 退订
)

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
