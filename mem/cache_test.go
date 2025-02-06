// Copyright (C) cabservd. 2025-present.
//
// Created at 2025-02-06, by liasica

package mem

import "testing"

func TestCache(t *testing.T) {
	s1 := Cache()
	s2 := Cache()
	t.Logf("s1: %p, s2: %p", s1, s2)
}
