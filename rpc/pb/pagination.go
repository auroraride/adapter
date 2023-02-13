// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-13
// Based on adapter by liasica, magicrolan@qq.com.

package pb

const (
    defaultPageSize    = 20
    defaultMaxPageSize = 100
    defaultCurrent     = 1
)

func (x *PaginationRequest) RealCurrent() (current int) {
    current = int(x.GetCurrent())
    if current < 1 {
        return defaultCurrent
    }
    return
}

func (x *PaginationRequest) RealPageSize() (size int) {
    size = int(x.GetPageSize())
    if size > defaultMaxPageSize {
        return defaultMaxPageSize
    }
    if size == 0 {
        return defaultPageSize
    }
    return
}

func (x *PaginationRequest) RealOffset() int {
    return x.RealPageSize() * (x.RealCurrent() - 1)
}
