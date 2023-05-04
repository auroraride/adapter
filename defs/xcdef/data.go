// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package xcdef

import (
	"database/sql/driver"
	"encoding/binary"
	"strings"
)

type Faults []Fault

func (s *Faults) Bytes() (data []byte) {
	for _, fault := range *s {
		data = append(data, uint8(fault))
	}
	return
}

func (s *Faults) FromBytes(data []byte) {
	for _, b := range data {
		*s = append(*s, Fault(b))
	}
	return
}

type Fault int

func (f *Fault) Scan(src interface{}) error {
	*f = Fault(src.(int64))
	return nil
}

func (f Fault) Value() (driver.Value, error) {
	return int64(f), nil
}

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

func NewMosStatus(b []byte) (ms *MosStatus) {
	ms = new(MosStatus)
	ms.FromBytes(b)
	return
}

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

// CanCharge 是否可充电
func (m MosStatus) CanCharge() bool {
	return m[0] == 1
}

// CanDisCharge 是否可放电
func (m MosStatus) CanDisCharge() bool {
	return m[1] == 1
}

func (m *MosStatus) Bytes() (data []byte) {
	for _, item := range *m {
		data = append(data, item)
	}
	return
}

func (m *MosStatus) FromBytes(data []byte) {
	for i, b := range data {
		m[i] = b
	}
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
	*s = GPSStatus(uint8(src.(int64)))
	return nil
}

func (s GPSStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

type MonVoltage [24]uint16

func NewMonVoltage(b []byte) (mv *MonVoltage) {
	mv = new(MonVoltage)
	mv.FromBytes(b)
	return
}

func (m *MonVoltage) Bytes() (data []byte) {
	data = make([]byte, 48)
	for i, item := range m {
		binary.BigEndian.PutUint16(data[i*2:i*2+2], item)
	}
	return
}

func (m *MonVoltage) FromBytes(data []byte) {
	for i := 0; i < 24; i++ {
		m[i] = binary.BigEndian.Uint16(data[i*2 : i*2+2])
	}
	return
}

type Temperature [4]uint16

func NewTemperature(b []byte) (tt *Temperature) {
	tt = new(Temperature)
	tt.FromBytes(b)
	return
}

func (m *Temperature) Bytes() (data []byte) {
	data = make([]byte, 8)
	for i, item := range *m {
		binary.BigEndian.PutUint16(data[i*2:i*2+2], item)
	}
	return
}

func (m *Temperature) FromBytes(data []byte) {
	for i := 0; i < 4; i++ {
		m[i] = binary.BigEndian.Uint16(data[i*2 : i*2+2])
	}
	return
}
