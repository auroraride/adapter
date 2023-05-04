// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-13
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import "testing"

func TestEntry(t *testing.T) {
	entry := NewEntry()
	t.Logf("%#v", entry)
	entry.Streams = append(entry.Streams, EntryStream{
		Stream: EntryJob{Job: "testjob"},
		Values: [][]string{{"value"}},
	})
	t.Logf("%#v", entry)
	ReleaseEntry(entry)
	t.Logf("%#v", NewEntry())
}
