// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: notifier.proto

package notifier

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

type LoadConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs map[string]string `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LoadConfigsRequest) Reset() {
	*x = LoadConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadConfigsRequest) ProtoMessage() {}

func (x *LoadConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadConfigsRequest.ProtoReflect.Descriptor instead.
func (*LoadConfigsRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *LoadConfigsRequest) GetConfigs() map[string]string {
	if x != nil {
		return x.Configs
	}
	return nil
}

type LoadConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadConfigsResponse) Reset() {
	*x = LoadConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadConfigsResponse) ProtoMessage() {}

func (x *LoadConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadConfigsResponse.ProtoReflect.Descriptor instead.
func (*LoadConfigsResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{1}
}

type UpdateX509SVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateX509SVIDRequest) Reset() {
	*x = UpdateX509SVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateX509SVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateX509SVIDRequest) ProtoMessage() {}

func (x *UpdateX509SVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateX509SVIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateX509SVIDRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{2}
}

type UpdateX509SVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateX509SVIDResponse) Reset() {
	*x = UpdateX509SVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateX509SVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateX509SVIDResponse) ProtoMessage() {}

func (x *UpdateX509SVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateX509SVIDResponse.ProtoReflect.Descriptor instead.
func (*UpdateX509SVIDResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{3}
}

var File_notifier_proto protoreflect.FileDescriptor

var file_notifier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x4c,
	0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x35, 0x30, 0x39,
	0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaf, 0x01, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x35, 0x30, 0x39, 0x53,
	0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x35, 0x30, 0x39,
	0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifier_proto_rawDescOnce sync.Once
	file_notifier_proto_rawDescData = file_notifier_proto_rawDesc
)

func file_notifier_proto_rawDescGZIP() []byte {
	file_notifier_proto_rawDescOnce.Do(func() {
		file_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifier_proto_rawDescData)
	})
	return file_notifier_proto_rawDescData
}

var file_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notifier_proto_goTypes = []interface{}{
	(*LoadConfigsRequest)(nil),     // 0: notifier.LoadConfigsRequest
	(*LoadConfigsResponse)(nil),    // 1: notifier.LoadConfigsResponse
	(*UpdateX509SVIDRequest)(nil),  // 2: notifier.UpdateX509SVIDRequest
	(*UpdateX509SVIDResponse)(nil), // 3: notifier.UpdateX509SVIDResponse
	nil,                            // 4: notifier.LoadConfigsRequest.ConfigsEntry
}
var file_notifier_proto_depIdxs = []int32{
	4, // 0: notifier.LoadConfigsRequest.configs:type_name -> notifier.LoadConfigsRequest.ConfigsEntry
	0, // 1: notifier.Notifier.LoadConfigs:input_type -> notifier.LoadConfigsRequest
	2, // 2: notifier.Notifier.UpdateX509SVID:input_type -> notifier.UpdateX509SVIDRequest
	1, // 3: notifier.Notifier.LoadConfigs:output_type -> notifier.LoadConfigsResponse
	3, // 4: notifier.Notifier.UpdateX509SVID:output_type -> notifier.UpdateX509SVIDResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notifier_proto_init() }
func file_notifier_proto_init() {
	if File_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadConfigsRequest); i {
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
		file_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadConfigsResponse); i {
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
		file_notifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateX509SVIDRequest); i {
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
		file_notifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateX509SVIDResponse); i {
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
			RawDescriptor: file_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifier_proto_goTypes,
		DependencyIndexes: file_notifier_proto_depIdxs,
		MessageInfos:      file_notifier_proto_msgTypes,
	}.Build()
	File_notifier_proto = out.File
	file_notifier_proto_rawDesc = nil
	file_notifier_proto_goTypes = nil
	file_notifier_proto_depIdxs = nil
}
