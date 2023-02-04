// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-22
// Based on adapter by liasica, magicrolan@qq.com.

package emixin

import (
    "entgo.io/contrib/entproto"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "time"
)

// TimeMixin 时间字段
type TimeMixin struct {
    mixin.Schema
    DisableIndex bool
    Optional     bool
}

func (t TimeMixin) Fields() []ent.Field {
    creator := field.Time("created_at").Immutable().Annotations(entproto.Skip())
    updator := field.Time("updated_at").Annotations(entproto.Skip())
    if t.Optional {
        creator.Optional().Nillable()
        updator.Optional().Nillable()
    }
    return []ent.Field{
        // .SchemaType(map[string]string{dialect.Postgres: "timestamp without time zone"})
        creator.Default(time.Now),
        updator.Default(time.Now).UpdateDefault(time.Now),
    }
}

func (t TimeMixin) Indexes() []ent.Index {
    var list []ent.Index
    if !t.DisableIndex {
        list = append(list, index.Fields("created_at"))
    }
    return list
}
