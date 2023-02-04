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
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/batdef"
    "google.golang.org/protobuf/types/descriptorpb"
    "time"
)

// Reign holds the schema definition for the Reign entity.
type Reign struct {
    ent.Schema
}

func (Reign) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "reign"},
        entsql.WithComments(true),
        entproto.Message(),
    }
}

// Fields of the Reign.
func (Reign) Fields() []ent.Field {
    return []ent.Field{
        field.Other("action", batdef.ReignActionUnknown).
            Default(batdef.ReignActionUnknown).
            SchemaType(map[string]string{dialect.Postgres: postgres.TypeSmallInt}).
            Comment("动作 1:入仓 2:离仓").Annotations(entproto.Field(2, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_INT64))),
        field.String("sn").Comment("电池编号").Annotations(entproto.Field(3)),
        field.Int("battery_id").Comment("电池ID").Annotations(entproto.Field(4)),
        field.Time("created_at").Immutable().Default(time.Now).Comment("记录时间").Annotations(entproto.Field(5)),
        field.String("serial").Comment("电柜编号").Annotations(entproto.Field(6)),
        field.Int("ordinal").Comment("仓位序号").Annotations(entproto.Field(7)),
        field.String("cabinet_name").Optional().Nillable().Comment("电柜名称").Annotations(entproto.Field(8)),
        field.String("remark").Optional().Nillable().Comment("备注信息").Annotations(entproto.Field(9)),
        field.Other("geom", &adapter.Geometry{}).
            SchemaType(map[string]string{dialect.Postgres: "geometry(POINT, 4326)"}).
            Comment("坐标").Annotations(entproto.Field(10, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
    }
}

// Edges of the Reign.
func (Reign) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("battery", Battery.Type).Field("battery_id").
            Ref("reigns").
            Unique().
            Required().
            Annotations(entsql.WithComments(true), entproto.Field(11)).
            Comment("所属电池"),
    }
}

func (Reign) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("sn").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
            entsql.OpClass("gin_trgm_ops"),
        ),
        index.Fields("battery_id"),
        index.Fields("created_at"),
        index.Fields("serial"),
        index.Fields("geom").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIST",
            }),
        ),
    }
}
