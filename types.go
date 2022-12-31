// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

const (
    BinIDKey     = "bin-%d"
    CabinetIDKey = "cabinet-%d"
)

type PointInt *int

type Bool bool

const (
    True  Bool = true
    False Bool = false
)

func (b Bool) String() string {
    switch b {
    case True:
        return "是"
    default:
        return "否"
    }
}

type (
    VoidFunc      func()
    BytesCallback func(b []byte)
)
