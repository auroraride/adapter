// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-15, by liasica

package adapter

import (
	"testing"
	"time"
)

func TestAddMonth(t *testing.T) {
	tests := []struct {
		input time.Time
		month int
		want  time.Time
	}{
		{
			input: time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC),
			month: 1,
			want:  time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			input: time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			month: 1,
			want:  time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		if got := AddMonth(tt.input, tt.month); got != tt.want {
			t.Errorf("AddMonth(%v, %v) = %v, want %v", tt.input, tt.month, got, tt.want)
		}
	}
}
