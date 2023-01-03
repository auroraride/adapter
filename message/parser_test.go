// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package message

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    "reflect"
    "testing"
)

func TestParser(t *testing.T) {
    sn := "XCB0862022110001"

    msg := &cabdef.BatteryMessage{
        Battery: adapter.ParseBatterySN(sn),
        Cabinet: "TEST",
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

    if typ != TypeCabkitBattery {
        t.Fail()
    }

    if reflect.DeepEqual(data, msg) {
        t.Fail()
    }
}
