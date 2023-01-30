// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "bytes"
    "database/sql/driver"
    "encoding/binary"
    "encoding/hex"
    "fmt"
)

// Geometry 坐标
// https://github.com/go-pg/pg/issues/829#issuecomment-505882885
type Geometry struct {
    Lng, Lat float64
}

func (g *Geometry) Scan(val interface{}) error {
    b, err := hex.DecodeString(ConvertBytes2String(val.([]byte)))
    if err != nil {
        return err
    }
    r := bytes.NewReader(b)
    var wkbByteOrder uint8
    if err = binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
        return err
    }

    var byteOrder binary.ByteOrder
    switch wkbByteOrder {
    case 0:
        byteOrder = binary.BigEndian
    case 1:
        byteOrder = binary.LittleEndian
    default:
        return fmt.Errorf("invalid byte order %d", wkbByteOrder)
    }

    var wkbGeometryType uint64
    if err = binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
        return err
    }

    var point [2]float64
    if err = binary.Read(r, byteOrder, &point); err != nil {
        return err
    }

    g.Lng = point[0]
    g.Lat = point[1]

    return nil
}

func (g *Geometry) String() string {
    return fmt.Sprintf("SRID=4326;POINT(%v %v)", g.Lng, g.Lat)
}

func (g Geometry) Value() (driver.Value, error) {
    return g.String(), nil
}
