// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBatterySN(t *testing.T) {
	bat, err := ParseBatterySN("XCB1191023030243")
	require.NoError(t, err)
	target := Battery{
		Brand: BatteryBrandXC,
		Model: "72V40AH",
		SN:    "XCB1191023030243",
	}
	t.Logf("bat: %#v", bat)

	if !reflect.DeepEqual(bat, target) {
		t.Fail()
	}

	bat, err = ParseBatterySN("BT106002512NNTB211118182")
	require.NoError(t, err)
	target = Battery{
		Brand: BatteryBrandTB,
		Model: "60V25AH",
		SN:    "BT106002512NNTB211118182",
	}
	t.Logf("bat: %#v", bat)

	if !reflect.DeepEqual(bat, target) {
		t.Fail()
	}
}

func BenchmarkParseBatterySN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseBatterySN("XCB0862022110001")
	}
}
