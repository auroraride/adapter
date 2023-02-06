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
    "github.com/auroraride/adapter/defs/xcdef"
    "google.golang.org/protobuf/types/descriptorpb"
    "time"
)

// Heartbeat holds the schema definition for the Heartbeat entity.
type Heartbeat struct {
    ent.Schema
}

func (Heartbeat) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "heartbeat"},
        entsql.WithComments(true),
        entproto.Message(),
    }
}

// Fields of the Heartbeat.
func (Heartbeat) Fields() []ent.Field {
    return []ent.Field{
        field.String("sn").Comment("电池编号").Annotations(entproto.Field(2)),
        field.Int("battery_id").Comment("电池ID").Annotations(entproto.Field(3)),
        field.Time("created_at").Immutable().Default(time.Now).Annotations(entproto.Field(4)),
        field.Float("voltage").Comment("电池总压 (V)").Annotations(entproto.Field(5)),
        field.Float("current").Comment("电流 (A, 充电为正, 放电为负)").Annotations(entproto.Field(6)),
        field.Uint8("soc").Comment("单位1%").Annotations(entproto.Field(7)),
        field.Uint8("soh").Comment("单位1%").Annotations(entproto.Field(8)),
        field.Bool("in_cabinet").Comment("是否在电柜").Annotations(entproto.Field(9)),
        field.Float("capacity").Comment("剩余容量 (单位AH)").Annotations(entproto.Field(10)),
        field.Uint16("mon_max_voltage").Comment("最大单体电压 (mV)").Annotations(entproto.Field(11)),
        field.Uint8("mon_max_voltage_pos").Comment("最大单体电压位置 (第x串)").Annotations(entproto.Field(12)),
        field.Uint16("mon_min_voltage").Comment("最小单体电压 (mV)").Annotations(entproto.Field(13)),
        field.Uint8("mon_min_voltage_pos").Comment("最小单体电压位置 (第x串)").Annotations(entproto.Field(14)),
        field.Uint16("max_temp").Comment("最大温度 (单位1℃)").Annotations(entproto.Field(15)),
        field.Uint16("min_temp").Comment("最小温度 (单位1℃)").Annotations(entproto.Field(16)),
        field.JSON("faults", &xcdef.Faults{}).Optional().Comment("故障列表").Annotations(
            entproto.Field(17, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES)),
        ),
        field.JSON("mos_status", &xcdef.MosStatus{}).Comment("MOS状态 (Bit0表示充电, Bit1表示放电, 此字段无法判定电池是否充放电状态)").Annotations(entproto.Field(18, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
        field.JSON("mon_voltage", &xcdef.MonVoltage{}).Comment("单体电压 (24个单体电压, 单位mV)").Annotations(entproto.Field(19, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
        field.JSON("temp", &xcdef.Temperature{}).Comment("电池温度 (4个电池温度传感器, 单位1℃)").Annotations(entproto.Field(20, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
        field.Uint16("mos_temp").Comment("MOS温度 (1个MOS温度传感器, 单位1℃)").Annotations(entproto.Field(21)),
        field.Uint16("env_temp").Comment("环境温度 (1个环境温度传感器, 单位1℃)").Annotations(entproto.Field(22)),
        field.Other("geom", &adapter.Geometry{}).
            SchemaType(map[string]string{dialect.Postgres: postgres.TypeGeometry}).
            Comment("坐标").
            Annotations(entproto.Field(23, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
        field.Other("gps", xcdef.GPSStatusNone).
            Default(xcdef.GPSStatusNone).
            SchemaType(map[string]string{dialect.Postgres: postgres.TypeSmallInt}).
            Comment("GPS定位状态 (0=未定位 1=GPS定位 4=LBS定位)").
            Annotations(entproto.Field(24, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_UINT32))),
        field.Uint8("strength").Comment("4G通讯信号强度 (0-100 百分比形式)").Annotations(entproto.Field(25)),
        field.Uint16("cycles").Comment("电池包循环次数 (80%累加一次)").Annotations(entproto.Field(26)),
        field.Uint32("charging_time").Comment("本次充电时长").Annotations(entproto.Field(27)),
        field.Uint32("dis_charging_time").Comment("本次放电时长").Annotations(entproto.Field(28)),
        field.Uint32("using_time").Comment("本次使用时长").Annotations(entproto.Field(29)),
        field.Uint32("total_charging_time").Comment("总充电时长").Annotations(entproto.Field(30)),
        field.Uint32("total_dis_charging_time").Comment("总放电时长").Annotations(entproto.Field(31)),
        field.Uint32("total_using_time").Comment("总使用时长").Annotations(entproto.Field(32)),
    }
}

// Edges of the Heartbeat.
func (Heartbeat) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("battery", Battery.Type).
            Field("battery_id").
            Ref("heartbeats").
            Unique().
            Required().
            Annotations(
                entsql.WithComments(true),
                entproto.Field(33),
            ).
            Comment("所属电池"),
    }
}

func (Heartbeat) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("created_at"),
        index.Fields("battery_id"),
        index.Fields("sn").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
            entsql.OpClass("gin_trgm_ops"),
        ),
        index.Fields("geom").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIST",
            }),
        ),
        index.Fields("faults").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
        ),
    }
}
