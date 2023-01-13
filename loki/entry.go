// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-13
// Based on adapter by liasica, magicrolan@qq.com.

package loki

import (
    jsoniter "github.com/json-iterator/go"
    "sync"
)

type EntryJob struct {
    Job string `json:"job"`
}

type EntryStream struct {
    // Stream map[string]string `json:"stream"`
    Stream EntryJob   `json:"stream"`
    Values [][]string `json:"values"`
}

type Entry struct {
    Streams []EntryStream `json:"streams"`
}

var entryPool = sync.Pool{
    New: func() interface{} {
        return new(Entry)
    },
}

func NewEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func ReleaseEntry(entry *Entry) {
    entry.Streams = nil
    entryPool.Put(entry)
}

func (e *Entry) String() string {
    str, _ := jsoniter.MarshalToString(e)
    return str
}

func (e *Entry) Bytes() []byte {
    b, _ := jsoniter.Marshal(e)
    return b
}
