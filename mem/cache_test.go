// Copyright (C) cabservd. 2025-present.
//
// Created at 2025-02-06, by liasica

package mem

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	db := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   15,
	})
	CacheSetup(db)

	c1 := Cache()
	c2 := Cache()
	t.Logf("c1: %p, s2: %p", c1, c2)
	require.Equal(t, c1, c2)

	serial := "KHD40-19040000"

	items := map[string]int{
		"XCB1262023030318": 1,
		"XCB1262023080582": 2,
		"XCB1262023041836": 3,
		"XCB1262023041464": 4,
		"XCB1262023052066": 5,
		"XCB1262023030678": 6,
	}

	ctx := context.Background()

	for sn, ordinal := range items {
		c1.AddCabinetBatteryCache(ctx, serial, ordinal, sn)
	}

	require.Equal(t, c1.CountCabinetBatteryCache(ctx, serial), int64(len(items)))
	require.Equal(t, c1.IsMemberCabinetBatteryCache(ctx, serial, "XCB1262023030318"), true)
	c1.RemoveCabinetBatteryCache(ctx, serial, "XCB1262023030318")
	require.Equal(t, c1.IsMemberCabinetBatteryCache(ctx, serial, "XCB1262023030318"), false)
	m := c1.ListCabinetBatteryCache(ctx, serial)
	t.Log(m)
	require.Equal(t, m["XCB1262023030678"], items["XCB1262023030678"])
}
