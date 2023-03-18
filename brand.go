// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
)

type CabinetBrand string

const (
    CabinetBrandUnknown CabinetBrand = "UNKNOWN"
    CabinetBrandKaixin  CabinetBrand = "KAIXIN"
    CabinetBrandYundong CabinetBrand = "YUNDONG"
    CabinetBrandTuobang CabinetBrand = "TUOBANG"
)

func (b CabinetBrand) RpcName() string {
    switch b {
    default:
        return ""
    case CabinetBrandKaixin:
        return "kxcab"
    case CabinetBrandYundong:
        return "ydcab"
    case CabinetBrandTuobang:
        return "tbcab"
    }
}

func (b CabinetBrand) String() string {
    return string(b)
}

func (b *CabinetBrand) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *b = CabinetBrand(v)
    }
    return nil
}

func (b CabinetBrand) Value() (driver.Value, error) {
    return b, nil
}

type BatteryBrand string

const (
    BatteryBrandUnknown BatteryBrand = "UNKNOWN" // 未知
    BatteryBrandXC      BatteryBrand = "XC"      // 星创电池
    BatteryBrandTB      BatteryBrand = "TB"      // 拓邦电池
)

func (b BatteryBrand) RpcName() string {
    switch b {
    default:
        return ""
    case BatteryBrandXC:
        return "xcbms"
    case BatteryBrandTB:
        return "tbbms"
    }
}

func (b BatteryBrand) String() string {
    return string(b)
}

func (b *BatteryBrand) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *b = BatteryBrand(v)
    }
    return nil
}

func (b BatteryBrand) Value() (driver.Value, error) {
    return b, nil
}
