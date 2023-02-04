package schema

import (
    "entgo.io/contrib/entproto"
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/auroraride/adapter/ent/emixin"
)

// Battery holds the schema definition for the Battery entity.
type Battery struct {
    ent.Schema
}

func (Battery) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "battery"},
        entsql.WithComments(true),
        entproto.Message(),
    }
}

// Fields of the Battery.
func (Battery) Fields() []ent.Field {
    return []ent.Field{
        field.String("sn").Unique().Comment("电池编号").Annotations(entproto.Field(2)),
        field.Uint16("soft_version").Optional().Nillable().Comment("BMS软件版本").Annotations(entproto.Field(3)),
        field.Uint16("hard_version").Optional().Nillable().Comment("BMS硬件版本").Annotations(entproto.Field(4)),
        field.Uint16("soft_4g_version").Optional().Nillable().Comment("4G软件版本").Annotations(entproto.Field(5)),
        field.Uint16("hard_4g_version").Optional().Nillable().Comment("4G硬件版本").Annotations(entproto.Field(6)),
        field.Uint64("sn_4g").Optional().Nillable().Comment("4G板SN").Annotations(entproto.Field(7)),
        field.String("iccid").Optional().Nillable().Comment("SIM卡ICCID").Annotations(entproto.Field(8)),
        field.Uint16("soc").Optional().Nillable().Comment("电池设计容量").Annotations(entproto.Field(9)),
    }
}

// Edges of the Battery.
func (Battery) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("heartbeats", Heartbeat.Type).Annotations(
            entsql.WithComments(true),
            entsql.DescColumns("created_at"),
            entproto.Field(10),
        ).Comment("心跳列表"),

        edge.To("reigns", Reign.Type).Annotations(
            entsql.WithComments(true),
            entsql.DescColumns("created_at"),
            entproto.Field(11),
        ).Comment("在位列表"),
    }
}

func (Battery) Mixin() []ent.Mixin {
    return []ent.Mixin{
        emixin.TimeMixin{},
    }
}

func (Battery) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("sn").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
            entsql.OpClass("gin_trgm_ops"),
        ),
    }
}
