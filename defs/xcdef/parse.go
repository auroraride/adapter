// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package xcdef

import (
    "database/sql/driver"
    "strings"
)

type Faults []Fault

type Fault uint8

func (f Fault) String() string {
    switch f {
    default:
        return " - "
    case 0:
        return "总压低"
    case 1:
        return "总压高"
    case 2:
        return "单体低"
    case 3:
        return "单体高"
    case 6:
        return "放电过流"
    case 7:
        return "充电过流"
    case 8:
        return "SOC低"
    case 11:
        return "充电高温"
    case 12:
        return "充电低温"
    case 13:
        return "放电高温"
    case 14:
        return "放电低温"
    case 15:
        return "短路"
    case 16:
        return "MOS高温"
    }
}

type MosStatus [2]uint8

func (m MosStatus) String() string {
    var builder strings.Builder
    builder.WriteString("充电")
    if m[0] == 1 {
        builder.WriteString("开")
    } else {
        builder.WriteString("关")
    }
    builder.WriteString(" / 放电")
    if m[1] == 1 {
        builder.WriteString("开")
    } else {
        builder.WriteString("关")
    }
    return builder.String()
}

type GPSStatus uint8

const (
    GPSStatusNone GPSStatus = 0
    GPSStatusRaw  GPSStatus = 1
    GPSStatusLBS  GPSStatus = 4
)

func (s GPSStatus) String() string {
    switch s {
    default:
        return " - "
    case GPSStatusNone:
        return "未定位"
    case GPSStatusRaw:
        return "GPS定位"
    case GPSStatusLBS:
        return "LBS定位"
    }
}

func (s *GPSStatus) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case uint8:
        *s = GPSStatus(v)
    }
    return nil
}

func (s GPSStatus) Value() (driver.Value, error) {
    return s, nil
}
