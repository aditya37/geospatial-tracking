// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: geotracking.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestGetDeviceLogByDeviceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *RequestGetDeviceLogByDeviceId) Reset() {
	*x = RequestGetDeviceLogByDeviceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geotracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetDeviceLogByDeviceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetDeviceLogByDeviceId) ProtoMessage() {}

func (x *RequestGetDeviceLogByDeviceId) ProtoReflect() protoreflect.Message {
	mi := &file_geotracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetDeviceLogByDeviceId.ProtoReflect.Descriptor instead.
func (*RequestGetDeviceLogByDeviceId) Descriptor() ([]byte, []int) {
	return file_geotracking_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGetDeviceLogByDeviceId) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ResponseGetDeviceLogByDeviceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResponseGetDeviceLogByDeviceId) Reset() {
	*x = ResponseGetDeviceLogByDeviceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geotracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetDeviceLogByDeviceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetDeviceLogByDeviceId) ProtoMessage() {}

func (x *ResponseGetDeviceLogByDeviceId) ProtoReflect() protoreflect.Message {
	mi := &file_geotracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetDeviceLogByDeviceId.ProtoReflect.Descriptor instead.
func (*ResponseGetDeviceLogByDeviceId) Descriptor() ([]byte, []int) {
	return file_geotracking_proto_rawDescGZIP(), []int{1}
}

var File_geotracking_proto protoreflect.FileDescriptor

var file_geotracking_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x65, 0x6f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x1d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x42, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x42, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x32, 0x74, 0x0a, 0x0b, 0x47, 0x65,
	0x6f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42,
	0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geotracking_proto_rawDescOnce sync.Once
	file_geotracking_proto_rawDescData = file_geotracking_proto_rawDesc
)

func file_geotracking_proto_rawDescGZIP() []byte {
	file_geotracking_proto_rawDescOnce.Do(func() {
		file_geotracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_geotracking_proto_rawDescData)
	})
	return file_geotracking_proto_rawDescData
}

var file_geotracking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_geotracking_proto_goTypes = []interface{}{
	(*RequestGetDeviceLogByDeviceId)(nil),  // 0: proto.RequestGetDeviceLogByDeviceId
	(*ResponseGetDeviceLogByDeviceId)(nil), // 1: proto.ResponseGetDeviceLogByDeviceId
}
var file_geotracking_proto_depIdxs = []int32{
	0, // 0: proto.Geotracking.GetDeviceLogByDeviceId:input_type -> proto.RequestGetDeviceLogByDeviceId
	1, // 1: proto.Geotracking.GetDeviceLogByDeviceId:output_type -> proto.ResponseGetDeviceLogByDeviceId
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geotracking_proto_init() }
func file_geotracking_proto_init() {
	if File_geotracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geotracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetDeviceLogByDeviceId); i {
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
		file_geotracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetDeviceLogByDeviceId); i {
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
			RawDescriptor: file_geotracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geotracking_proto_goTypes,
		DependencyIndexes: file_geotracking_proto_depIdxs,
		MessageInfos:      file_geotracking_proto_msgTypes,
	}.Build()
	File_geotracking_proto = out.File
	file_geotracking_proto_rawDesc = nil
	file_geotracking_proto_goTypes = nil
	file_geotracking_proto_depIdxs = nil
}
