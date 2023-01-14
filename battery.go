// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "strconv"
)

var (
    BatteryModelXC = map[string]string{
        "08": "72V30AH",
    }
)

type Battery struct {
    SN     string       `json:"sn"`     // 电池编号
    Brand  BatteryBrand `json:"brand"`  // 电池厂家
    Model  string       `json:"model"`  // 电池型号
    Year   int          `json:"year"`   // 生产年份
    Month  int          `json:"month"`  // 生产月份
    Serial string       `json:"serial"` // 流水号
}

// ParseBatterySN 解析电池编号
func ParseBatterySN(sn string) (bat *Battery) {
    if len(sn) < 16 {
        return &Battery{}
    }

    bat = &Battery{
        Brand:  BatteryBrand(sn[0:2]),
        Model:  BatteryModelXC[sn[3:5]],
        Serial: sn[12:],
        SN:     sn,
    }

    year, _ := strconv.ParseInt(sn[8:10], 10, 64)
    month, _ := strconv.ParseInt(sn[10:12], 10, 64)

    bat.Year = 2000 + int(year)
    bat.Month = int(month)
    return
}
