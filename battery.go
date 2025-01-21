// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "errors"

const (
	batteryXcLength = 16 // 星创电池长度
	batteryTbLength = 24 // 拓邦电池长度
)

var (
	BatteryModelXC = map[string]string{
		"08": "72V30AH",
		"11": "72V35AH",
		"12": "60V30AH",
		"16": "72V35AH",
		"17": "60V45AH",
	}
)

type Battery struct {
	SN    string       `json:"sn"`    // 电池编号
	Brand BatteryBrand `json:"brand"` // 电池厂家
	Model string       `json:"model"` // 电池型号
}

// ParseBatterySN 解析电池编号
func ParseBatterySN(sn string) (bat Battery, err error) {
	if len(sn) < 16 {
		return bat, errors.New("电池编码位数错误 (" + sn + ")")
	}

	b := make([]byte, len(sn))
	for i := range b {
		c := sn[i]
		switch {
		case c >= 'a' && c <= 'z':
			c -= 'a' - 'A'
		case c < '0', c > 'z', c > '9' && c < 'A', c > 'Z' && c < 'a':
			return Battery{}, errors.New("电池编码错误 (" + sn + ")")
		}
		b[i] = c
	}
	sn = ConvertBytes2String(b)

	// 按字符串长度简单区分拓邦和星创 >.<
	switch len(sn) {
	case batteryXcLength:
		bat.Brand = BatteryBrandXC
		bat.Model = BatteryModelXC[sn[3:5]]
	case batteryTbLength:
		bat.Brand = BatteryBrandTB
		bat.Model = sn[4:6] + "V" + sn[7:9] + "AH"
	default:
		return Battery{}, errors.New("电池品牌编码解析失败 (" + sn + ")")
	}

	bat.SN = sn

	return
}
