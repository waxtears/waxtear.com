// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: TestMicroProto/TestMicro.proto

package TestMicroProto

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

type SayHelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloReq) Reset() {
	*x = SayHelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMicroProto_TestMicro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloReq) ProtoMessage() {}

func (x *SayHelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_TestMicroProto_TestMicro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloReq.ProtoReflect.Descriptor instead.
func (*SayHelloReq) Descriptor() ([]byte, []int) {
	return file_TestMicroProto_TestMicro_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SayHelloRsp) Reset() {
	*x = SayHelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMicroProto_TestMicro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRsp) ProtoMessage() {}

func (x *SayHelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TestMicroProto_TestMicro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRsp.ProtoReflect.Descriptor instead.
func (*SayHelloRsp) Descriptor() ([]byte, []int) {
	return file_TestMicroProto_TestMicro_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_TestMicroProto_TestMicro_proto protoreflect.FileDescriptor

var file_TestMicroProto_TestMicro_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x21, 0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0x37, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0c, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x20, 0x5a,
	0x1e, 0x2f, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TestMicroProto_TestMicro_proto_rawDescOnce sync.Once
	file_TestMicroProto_TestMicro_proto_rawDescData = file_TestMicroProto_TestMicro_proto_rawDesc
)

func file_TestMicroProto_TestMicro_proto_rawDescGZIP() []byte {
	file_TestMicroProto_TestMicro_proto_rawDescOnce.Do(func() {
		file_TestMicroProto_TestMicro_proto_rawDescData = protoimpl.X.CompressGZIP(file_TestMicroProto_TestMicro_proto_rawDescData)
	})
	return file_TestMicroProto_TestMicro_proto_rawDescData
}

var file_TestMicroProto_TestMicro_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TestMicroProto_TestMicro_proto_goTypes = []interface{}{
	(*SayHelloReq)(nil), // 0: SayHelloReq
	(*SayHelloRsp)(nil), // 1: SayHelloRsp
}
var file_TestMicroProto_TestMicro_proto_depIdxs = []int32{
	0, // 0: TestService.SayHello:input_type -> SayHelloReq
	1, // 1: TestService.SayHello:output_type -> SayHelloRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TestMicroProto_TestMicro_proto_init() }
func file_TestMicroProto_TestMicro_proto_init() {
	if File_TestMicroProto_TestMicro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TestMicroProto_TestMicro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloReq); i {
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
		file_TestMicroProto_TestMicro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRsp); i {
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
			RawDescriptor: file_TestMicroProto_TestMicro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TestMicroProto_TestMicro_proto_goTypes,
		DependencyIndexes: file_TestMicroProto_TestMicro_proto_depIdxs,
		MessageInfos:      file_TestMicroProto_TestMicro_proto_msgTypes,
	}.Build()
	File_TestMicroProto_TestMicro_proto = out.File
	file_TestMicroProto_TestMicro_proto_rawDesc = nil
	file_TestMicroProto_TestMicro_proto_goTypes = nil
	file_TestMicroProto_TestMicro_proto_depIdxs = nil
}
