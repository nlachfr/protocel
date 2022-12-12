// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: testdata/authorize/extended/extended.proto

package extended

import (
	_ "github.com/Neakxs/protocel/authorize"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExtendedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ExtendedRequest) Reset() {
	*x = ExtendedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_authorize_extended_extended_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedRequest) ProtoMessage() {}

func (x *ExtendedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_authorize_extended_extended_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedRequest.ProtoReflect.Descriptor instead.
func (*ExtendedRequest) Descriptor() ([]byte, []int) {
	return file_testdata_authorize_extended_extended_proto_rawDescGZIP(), []int{0}
}

func (x *ExtendedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_testdata_authorize_extended_extended_proto protoreflect.FileDescriptor

var file_testdata_authorize_extended_extended_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x1a,
	0x19, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x5b,
	0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x22, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xb7, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x61, 0x6b, 0x78,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0xca, 0x47, 0x7c, 0x0a, 0x29, 0x0a, 0x0e, 0x12, 0x0c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x15, 0x0a,
	0x09, 0x67, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x08, 0x0a, 0x02, 0x08, 0x06,
	0x12, 0x02, 0x08, 0x06, 0x12, 0x4f, 0x0a, 0x2a, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x1f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x67, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x28,
	0x6e, 0x61, 0x6d, 0x65, 0x29, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_authorize_extended_extended_proto_rawDescOnce sync.Once
	file_testdata_authorize_extended_extended_proto_rawDescData = file_testdata_authorize_extended_extended_proto_rawDesc
)

func file_testdata_authorize_extended_extended_proto_rawDescGZIP() []byte {
	file_testdata_authorize_extended_extended_proto_rawDescOnce.Do(func() {
		file_testdata_authorize_extended_extended_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_authorize_extended_extended_proto_rawDescData)
	})
	return file_testdata_authorize_extended_extended_proto_rawDescData
}

var file_testdata_authorize_extended_extended_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testdata_authorize_extended_extended_proto_goTypes = []interface{}{
	(*ExtendedRequest)(nil), // 0: testdata.extended.ExtendedRequest
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
}
var file_testdata_authorize_extended_extended_proto_depIdxs = []int32{
	0, // 0: testdata.extended.ExtendedService.Extended:input_type -> testdata.extended.ExtendedRequest
	1, // 1: testdata.extended.ExtendedService.Extended:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_authorize_extended_extended_proto_init() }
func file_testdata_authorize_extended_extended_proto_init() {
	if File_testdata_authorize_extended_extended_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_authorize_extended_extended_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendedRequest); i {
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
			RawDescriptor: file_testdata_authorize_extended_extended_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testdata_authorize_extended_extended_proto_goTypes,
		DependencyIndexes: file_testdata_authorize_extended_extended_proto_depIdxs,
		MessageInfos:      file_testdata_authorize_extended_extended_proto_msgTypes,
	}.Build()
	File_testdata_authorize_extended_extended_proto = out.File
	file_testdata_authorize_extended_extended_proto_rawDesc = nil
	file_testdata_authorize_extended_extended_proto_goTypes = nil
	file_testdata_authorize_extended_extended_proto_depIdxs = nil
}
