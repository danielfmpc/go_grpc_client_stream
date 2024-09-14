// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: src/proto/calc-service.proto

package calc

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_calc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_calc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_src_proto_calc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int32   `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Avarage  float64 `protobuf:"fixed64,2,opt,name=avarage,proto3" json:"avarage,omitempty"`
	Total    int32   `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_calc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_calc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_src_proto_calc_service_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Output) GetAvarage() float64 {
	if x != nil {
		return x.Avarage
	}
	return 0
}

func (x *Output) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_src_proto_calc_service_proto protoreflect.FileDescriptor

var file_src_proto_calc_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x63, 0x61, 0x6c, 0x63, 0x22, 0x1d, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x61,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x32, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63,
	0x12, 0x0b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0c, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x28, 0x01, 0x42, 0x0e, 0x5a,
	0x0c, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_calc_service_proto_rawDescOnce sync.Once
	file_src_proto_calc_service_proto_rawDescData = file_src_proto_calc_service_proto_rawDesc
)

func file_src_proto_calc_service_proto_rawDescGZIP() []byte {
	file_src_proto_calc_service_proto_rawDescOnce.Do(func() {
		file_src_proto_calc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_calc_service_proto_rawDescData)
	})
	return file_src_proto_calc_service_proto_rawDescData
}

var file_src_proto_calc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_calc_service_proto_goTypes = []any{
	(*Input)(nil),  // 0: calc.Input
	(*Output)(nil), // 1: calc.Output
}
var file_src_proto_calc_service_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.Calc:input_type -> calc.Input
	1, // 1: calc.CalcService.Calc:output_type -> calc.Output
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_calc_service_proto_init() }
func file_src_proto_calc_service_proto_init() {
	if File_src_proto_calc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_calc_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Input); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_calc_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Output); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_proto_calc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_calc_service_proto_goTypes,
		DependencyIndexes: file_src_proto_calc_service_proto_depIdxs,
		MessageInfos:      file_src_proto_calc_service_proto_msgTypes,
	}.Build()
	File_src_proto_calc_service_proto = out.File
	file_src_proto_calc_service_proto_rawDesc = nil
	file_src_proto_calc_service_proto_goTypes = nil
	file_src_proto_calc_service_proto_depIdxs = nil
}
