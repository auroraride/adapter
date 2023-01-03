// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package cabdef

import "database/sql/driver"

type Brand string

const (
    BrandUnknown Brand = "UNKNOWN"
    BrandKaixin  Brand = "KAIXIN"
)

func (b Brand) String() string {
    return string(b)
}

func (b *Brand) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *b = Brand(v)
    }
    return nil
}

func (b Brand) Value() (driver.Value, error) {
    return b, nil
}
