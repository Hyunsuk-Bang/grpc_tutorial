// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: avg.proto

package proto

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

type AvgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AvgRequest) Reset() {
	*x = AvgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgRequest) ProtoMessage() {}

func (x *AvgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvgRequest.ProtoReflect.Descriptor instead.
func (*AvgRequest) Descriptor() ([]byte, []int) {
	return file_avg_proto_rawDescGZIP(), []int{0}
}

func (x *AvgRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AvgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avg float64 `protobuf:"fixed64,1,opt,name=avg,proto3" json:"avg,omitempty"`
}

func (x *AvgResponse) Reset() {
	*x = AvgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgResponse) ProtoMessage() {}

func (x *AvgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvgResponse.ProtoReflect.Descriptor instead.
func (*AvgResponse) Descriptor() ([]byte, []int) {
	return file_avg_proto_rawDescGZIP(), []int{1}
}

func (x *AvgResponse) GetAvg() float64 {
	if x != nil {
		return x.Avg
	}
	return 0
}

var File_avg_proto protoreflect.FileDescriptor

var file_avg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x76, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x76, 0x67,
	0x22, 0x24, 0x0a, 0x0a, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x61, 0x76, 0x67, 0x32, 0x3b, 0x0a, 0x0a, 0x41, 0x76, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x76, 0x67, 0x12,
	0x0f, 0x2e, 0x61, 0x76, 0x67, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x61, 0x76, 0x67, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x76, 0x67, 0x52, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avg_proto_rawDescOnce sync.Once
	file_avg_proto_rawDescData = file_avg_proto_rawDesc
)

func file_avg_proto_rawDescGZIP() []byte {
	file_avg_proto_rawDescOnce.Do(func() {
		file_avg_proto_rawDescData = protoimpl.X.CompressGZIP(file_avg_proto_rawDescData)
	})
	return file_avg_proto_rawDescData
}

var file_avg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_avg_proto_goTypes = []interface{}{
	(*AvgRequest)(nil),  // 0: avg.AvgRequest
	(*AvgResponse)(nil), // 1: avg.AvgResponse
}
var file_avg_proto_depIdxs = []int32{
	0, // 0: avg.AvgService.GetAvg:input_type -> avg.AvgRequest
	1, // 1: avg.AvgService.GetAvg:output_type -> avg.AvgResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_avg_proto_init() }
func file_avg_proto_init() {
	if File_avg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgRequest); i {
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
		file_avg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgResponse); i {
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
			RawDescriptor: file_avg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avg_proto_goTypes,
		DependencyIndexes: file_avg_proto_depIdxs,
		MessageInfos:      file_avg_proto_msgTypes,
	}.Build()
	File_avg_proto = out.File
	file_avg_proto_rawDesc = nil
	file_avg_proto_goTypes = nil
	file_avg_proto_depIdxs = nil
}
