package schema

import (
    "ariga.io/atlas/sql/postgres"
    "entgo.io/contrib/entproto"
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/auroraride/adapter/defs/xcdef"
    "google.golang.org/protobuf/types/descriptorpb"
    "time"
)

// Fault holds the schema definition for the Fault entity.
type Fault struct {
    ent.Schema
}

func (Fault) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "fault"},
        entsql.WithComments(true),
        entproto.Message(),
    }
}

// Fields of the Fault.
func (Fault) Fields() []ent.Field {
    return []ent.Field{
        field.String("sn").Comment("电池编号").Annotations(entproto.Field(2)),
        field.Int("battery_id").Comment("电池ID").Annotations(entproto.Field(3)),
        field.Other("fault", xcdef.Fault(-1)).Comment("故障信息").
            SchemaType(map[string]string{dialect.Postgres: postgres.TypeInt}).
            Annotations(entproto.Field(4, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_INT64))),
        field.Time("begin_at").Immutable().Default(time.Now).Comment("开始时间").Annotations(entproto.Field(5)),
        field.Time("end_at").Optional().Comment("结束时间").Annotations(entproto.Field(6)),
    }
}

// Edges of the Fault.
func (Fault) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("battery", Battery.Type).
            Field("battery_id").
            Ref("fault_log").
            Unique().
            Required().
            Annotations(
                entsql.WithComments(true),
                entproto.Field(7),
            ).
            Comment("所属电池"),
    }
}

func (Fault) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("begin_at"),
        index.Fields("end_at"),
        index.Fields("battery_id"),
        index.Fields("sn").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
            entsql.OpClass("gin_trgm_ops"),
        ),
        index.Fields("fault"),
    }
}
