// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-27
// Based on adapter by liasica, magicrolan@qq.com.

package batdef

import (
    "database/sql/driver"
    jsoniter "github.com/json-iterator/go"
)

type ReignAction uint8

const (
    ReignActionUnknown ReignAction = iota
    ReignActionIn                  // 入仓
    ReignActionOut                 // 离仓
)

func (s ReignAction) String() string {
    switch s {
    default:
        return " - "
    case ReignActionIn:
        return "入仓"
    case ReignActionOut:
        return "离仓"
    }
}

func (s *ReignAction) Scan(src interface{}) error {
    *s = ReignAction(uint8(src.(int64)))
    return nil
}

func (s ReignAction) Value() (driver.Value, error) {
    return int64(s), nil
}

func (s *ReignAction) MarshalBinary() (data []byte, err error) {
    return jsoniter.Marshal(s)
}

func (s *ReignAction) UnmarshalBinary(data []byte) (err error) {
    return jsoniter.Unmarshal(data, s)
}

// Reign 电池在位
type Reign struct {
    SN      string      `json:"sn" validate:"required"`
    Serial  string      `json:"serial" validate:"required"`  // 电柜编码
    Ordinal *int        `json:"ordinal" validate:"required"` // 仓位序号
    Action  ReignAction `json:"action" validate:"required"`  // 动作

    Lng    *float64 `json:"lng"`
    Lat    *float64 `json:"lat"`
    Remark *string  `json:"remark"`
}

// Clone 克隆电池在位结构体, 删除电池编码和动作
func (r *Reign) Clone(sn string, action ReignAction) *Reign {
    return &Reign{
        Serial:  r.Serial,
        Ordinal: r.Ordinal,
        Lng:     r.Lng,
        Lat:     r.Lat,
        Remark:  r.Remark,
    }
}
