// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-07
// Based on adapter by liasica, magicrolan@qq.com.

package timestamppb

import "time"

func Now() *Timestamp {
    return New(time.Now().In(time.Local))
}

func New(t time.Time) *Timestamp {
    return &Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())}
}

func (x *Timestamp) AsTime() time.Time {
    return time.Unix(x.GetSeconds(), int64(x.GetNanos())).In(time.Local)
}
