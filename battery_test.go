// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/stretchr/testify/require"
    "reflect"
    "testing"
)

func TestParseBatterySN(t *testing.T) {
    sn := "XCB0862022110001"
    t.Logf("battery sn: %s", sn)

    bat, err := ParseBatterySN(sn)
    require.NoError(t, err)
    target := &Battery{
        Brand:  BatteryBrandXC,
        Model:  "72V30AH",
        Year:   2022,
        Month:  11,
        Serial: "0001",
        SN:     sn,
    }

    t.Logf("bat: %#v", bat)

    if !reflect.DeepEqual(bat, target) {
        t.Fail()
    }
}

func BenchmarkName(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, _ = ParseBatterySN("XCB0862022110001")
    }
}
