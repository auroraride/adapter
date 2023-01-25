// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-25
// Based on adapter by liasica, magicrolan@qq.com.

package batdef

type InCabinetRequest struct {
    Cabinet *Cabinet `json:"cabinet" validate:"required"`
    SN      string   `json:"sn" validate:"required"`
}
