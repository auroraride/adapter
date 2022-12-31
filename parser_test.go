// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "reflect"
    "testing"
)

func TestParser(t *testing.T) {
    sn := "XCB0862022110001"

    msg := &BatteryMessage{
        Battery: ParseBatterySN(sn),
        Serial:  "TEST",
    }

    b, err := Pack(msg)
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    typ, data, err := Unpack(b)
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    if typ != TypeBattery {
        t.Fail()
    }

    if reflect.DeepEqual(data, msg) {
        t.Fail()
    }
}
