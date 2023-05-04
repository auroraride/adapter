// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChar(t *testing.T) {
	require.Equal(t, "\n", ConvertBytes2String(Newline))
	require.Equal(t, " ", ConvertBytes2String(Space))
	require.Equal(t, "\"", ConvertBytes2String(DoubleQuote))
	require.Equal(t, "'", ConvertBytes2String(SingleQuote))
	require.Equal(t, "", ConvertBytes2String(Comma))
	require.Equal(t, "", ConvertBytes2String(Hyphen))
	require.Equal(t, "", ConvertBytes2String(Period))
	require.Equal(t, "", ConvertBytes2String(Colon))
	require.Equal(t, "", ConvertBytes2String(Equal))
	require.Equal(t, "", ConvertBytes2String(LeftSquareBracket))
	require.Equal(t, "", ConvertBytes2String(RightSquareBracket))
	require.Equal(t, "", ConvertBytes2String(LeftBracket))
	require.Equal(t, "", ConvertBytes2String(RightBracket))
	require.Equal(t, "", ConvertBytes2String(LeftCurlyBracket))
	require.Equal(t, "", ConvertBytes2String(RightCurlyBracket))
}
