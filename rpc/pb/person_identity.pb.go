// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.3
// source: person_identity.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PersonIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdCardNumber string                   `protobuf:"bytes,1,opt,name=idCardNumber,proto3" json:"idCardNumber,omitempty"`
	Name         string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OcrResult    *PersonIdentityOcrResult `protobuf:"bytes,3,opt,name=ocrResult,proto3,oneof" json:"ocrResult,omitempty"`
}

func (x *PersonIdentity) Reset() {
	*x = PersonIdentity{}
	mi := &file_person_identity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonIdentity) ProtoMessage() {}

func (x *PersonIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_person_identity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonIdentity.ProtoReflect.Descriptor instead.
func (*PersonIdentity) Descriptor() ([]byte, []int) {
	return file_person_identity_proto_rawDescGZIP(), []int{0}
}

func (x *PersonIdentity) GetIdCardNumber() string {
	if x != nil {
		return x.IdCardNumber
	}
	return ""
}

func (x *PersonIdentity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonIdentity) GetOcrResult() *PersonIdentityOcrResult {
	if x != nil {
		return x.OcrResult
	}
	return nil
}

type PersonIdentityOcrResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex             string  `protobuf:"bytes,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Nation          string  `protobuf:"bytes,3,opt,name=nation,proto3" json:"nation,omitempty"`
	Birth           string  `protobuf:"bytes,4,opt,name=birth,proto3" json:"birth,omitempty"`
	Address         string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	IdCardNumber    string  `protobuf:"bytes,6,opt,name=idCardNumber,proto3" json:"idCardNumber,omitempty"`
	ValidStartDate  string  `protobuf:"bytes,7,opt,name=validStartDate,proto3" json:"validStartDate,omitempty"`
	ValidExpireDate string  `protobuf:"bytes,8,opt,name=validExpireDate,proto3" json:"validExpireDate,omitempty"`
	Authority       string  `protobuf:"bytes,9,opt,name=authority,proto3" json:"authority,omitempty"`
	PortraitCrop    string  `protobuf:"bytes,10,opt,name=portraitCrop,proto3" json:"portraitCrop,omitempty"`
	NationalCrop    string  `protobuf:"bytes,11,opt,name=nationalCrop,proto3" json:"nationalCrop,omitempty"`
	PortraitClarity float64 `protobuf:"fixed64,12,opt,name=portraitClarity,proto3" json:"portraitClarity,omitempty"`
	NationalClarity float64 `protobuf:"fixed64,13,opt,name=nationalClarity,proto3" json:"nationalClarity,omitempty"`
}

func (x *PersonIdentityOcrResult) Reset() {
	*x = PersonIdentityOcrResult{}
	mi := &file_person_identity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonIdentityOcrResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonIdentityOcrResult) ProtoMessage() {}

func (x *PersonIdentityOcrResult) ProtoReflect() protoreflect.Message {
	mi := &file_person_identity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonIdentityOcrResult.ProtoReflect.Descriptor instead.
func (*PersonIdentityOcrResult) Descriptor() ([]byte, []int) {
	return file_person_identity_proto_rawDescGZIP(), []int{1}
}

func (x *PersonIdentityOcrResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetIdCardNumber() string {
	if x != nil {
		return x.IdCardNumber
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetValidStartDate() string {
	if x != nil {
		return x.ValidStartDate
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetValidExpireDate() string {
	if x != nil {
		return x.ValidExpireDate
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetPortraitCrop() string {
	if x != nil {
		return x.PortraitCrop
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetNationalCrop() string {
	if x != nil {
		return x.NationalCrop
	}
	return ""
}

func (x *PersonIdentityOcrResult) GetPortraitClarity() float64 {
	if x != nil {
		return x.PortraitClarity
	}
	return 0
}

func (x *PersonIdentityOcrResult) GetNationalClarity() float64 {
	if x != nil {
		return x.NationalClarity
	}
	return 0
}

var File_person_identity_proto protoreflect.FileDescriptor

var file_person_identity_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x96, 0x01, 0x0a, 0x0e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x63, 0x72,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f, 0x63, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0xb7, 0x03, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x43, 0x72,
	0x6f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x43, 0x72, 0x6f, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x43, 0x72, 0x6f, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x6f, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x43, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x43, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72,
	0x6f, 0x72, 0x61, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_identity_proto_rawDescOnce sync.Once
	file_person_identity_proto_rawDescData = file_person_identity_proto_rawDesc
)

func file_person_identity_proto_rawDescGZIP() []byte {
	file_person_identity_proto_rawDescOnce.Do(func() {
		file_person_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_identity_proto_rawDescData)
	})
	return file_person_identity_proto_rawDescData
}

var file_person_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_person_identity_proto_goTypes = []any{
	(*PersonIdentity)(nil),          // 0: pb.PersonIdentity
	(*PersonIdentityOcrResult)(nil), // 1: pb.PersonIdentityOcrResult
}
var file_person_identity_proto_depIdxs = []int32{
	1, // 0: pb.PersonIdentity.ocrResult:type_name -> pb.PersonIdentityOcrResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_person_identity_proto_init() }
func file_person_identity_proto_init() {
	if File_person_identity_proto != nil {
		return
	}
	file_person_identity_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_identity_proto_goTypes,
		DependencyIndexes: file_person_identity_proto_depIdxs,
		MessageInfos:      file_person_identity_proto_msgTypes,
	}.Build()
	File_person_identity_proto = out.File
	file_person_identity_proto_rawDesc = nil
	file_person_identity_proto_goTypes = nil
	file_person_identity_proto_depIdxs = nil
}
