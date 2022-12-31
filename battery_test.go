// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "reflect"
    "testing"
)

func TestParseBatterySN(t *testing.T) {
    sn := "XCB0862022110001"
    t.Logf("battery sn: %s", sn)

    bat := ParseBatterySN(sn)
    target := &Battery{
        Brand:  BatteryBrandXC,
        Model:  "72V30AH",
        Year:   2022,
        Month:  11,
        Serial: "0001",
    }

    t.Logf("bat: %#v", bat)

    if !reflect.DeepEqual(bat, target) {
        t.Fail()
    }
}
