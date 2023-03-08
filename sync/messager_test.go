// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-08
// Based on adapter by liasica, magicrolan@qq.com.

package sync

import (
    "github.com/auroraride/adapter/defs/batdef"
    "github.com/stretchr/testify/require"
    "testing"
)

func TestUnmarshal(t *testing.T) {
    input := map[string]any{
        "__DATA__": `{"out":{"sn":"XCB0862022110277","brand":"XC","model":"72V30AH"},"serial":"CHZD08KXHD221115037","ordinal":1}`,
    }
    output, err := Unmarshal[batdef.BatteryFlow]("__DATA__", input)
    require.NoError(t, err)
    t.Log(output)
}
