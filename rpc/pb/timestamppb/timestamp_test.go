// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-07
// Based on adapter by liasica, magicrolan@qq.com.

package timestamppb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNew(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	str := "2023-03-07 08:46:18"
	tx, err := time.Parse(layout, str)
	require.NoError(t, err)
	tt := timestamppb.New(tx)
	t.Log(tt.AsTime().Format(layout))
	require.Equal(t, tt.AsTime(), tx)
}
