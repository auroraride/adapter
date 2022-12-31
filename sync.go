// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

// type DataType string
//
// // TODO 使用字头包封装类型
//
// const (
//     DataTypeCabinetSync  DataType = "cabinet_sync"
//     DataTypeExchangeStep DataType = "exchange_step"
// )
//
// type DataStuff interface {
//     CabinetSyncData | ExchangeStepResult
// }
//
// type CabinetSyncData struct {
//     Full    bool     `json:"full"`
//     Serial  string   `json:"serial"`
//     Cabinet *Cabinet `json:"cabinet,omitempty"`
//     Bins    []*Bin   `json:"bins,omitempty"`
// }
//
// type Data[T DataStuff] struct {
//     Type  DataType `json:"type"`
//     Value *T       `json:"value"`
// }
