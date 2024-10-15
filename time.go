// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-15, by liasica

package adapter

import "time"

func AddMonth(t time.Time, m int) time.Time {
	x := t.AddDate(0, m, 0)
	if d := x.Day(); d != t.Day() {
		return x.AddDate(0, 0, -d)
	}
	return x
}
