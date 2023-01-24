// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-22
// Based on adapter by liasica, magicrolan@qq.com.

package xcdef

type Cabinet struct {
    ID     uint64 `json:"id"`
    Serial string `json:"serial"` // 电柜编码
    Name   string `json:"name"`   // 名称
    CityID uint64 `json:"cityId"` // 城市ID
}
