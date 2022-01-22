// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/department/v1/department.proto

package v1

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

type GetDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDepartmentRequest) Reset() {
	*x = GetDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_department_v1_department_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentRequest) ProtoMessage() {}

func (x *GetDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_department_v1_department_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentRequest.ProtoReflect.Descriptor instead.
func (*GetDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_api_department_v1_department_proto_rawDescGZIP(), []int{0}
}

func (x *GetDepartmentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDepartmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	No   string `protobuf:"bytes,2,opt,name=no,proto3" json:"no,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDepartmentReply) Reset() {
	*x = GetDepartmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_department_v1_department_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentReply) ProtoMessage() {}

func (x *GetDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_department_v1_department_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentReply.ProtoReflect.Descriptor instead.
func (*GetDepartmentReply) Descriptor() ([]byte, []int) {
	return file_api_department_v1_department_proto_rawDescGZIP(), []int{1}
}

func (x *GetDepartmentReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDepartmentReply) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *GetDepartmentReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_department_v1_department_proto protoreflect.FileDescriptor

var file_api_department_v1_department_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6d, 0x0a, 0x0a, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x32, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1b, 0x67, 0x62, 0x6d, 0x65, 0x72, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_department_v1_department_proto_rawDescOnce sync.Once
	file_api_department_v1_department_proto_rawDescData = file_api_department_v1_department_proto_rawDesc
)

func file_api_department_v1_department_proto_rawDescGZIP() []byte {
	file_api_department_v1_department_proto_rawDescOnce.Do(func() {
		file_api_department_v1_department_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_department_v1_department_proto_rawDescData)
	})
	return file_api_department_v1_department_proto_rawDescData
}

var file_api_department_v1_department_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_department_v1_department_proto_goTypes = []interface{}{
	(*GetDepartmentRequest)(nil), // 0: api.department.v1.GetDepartmentRequest
	(*GetDepartmentReply)(nil),   // 1: api.department.v1.GetDepartmentReply
}
var file_api_department_v1_department_proto_depIdxs = []int32{
	0, // 0: api.department.v1.Department.GetDepartment:input_type -> api.department.v1.GetDepartmentRequest
	1, // 1: api.department.v1.Department.GetDepartment:output_type -> api.department.v1.GetDepartmentReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_department_v1_department_proto_init() }
func file_api_department_v1_department_proto_init() {
	if File_api_department_v1_department_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_department_v1_department_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepartmentRequest); i {
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
		file_api_department_v1_department_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepartmentReply); i {
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
			RawDescriptor: file_api_department_v1_department_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_department_v1_department_proto_goTypes,
		DependencyIndexes: file_api_department_v1_department_proto_depIdxs,
		MessageInfos:      file_api_department_v1_department_proto_msgTypes,
	}.Build()
	File_api_department_v1_department_proto = out.File
	file_api_department_v1_department_proto_rawDesc = nil
	file_api_department_v1_department_proto_goTypes = nil
	file_api_department_v1_department_proto_depIdxs = nil
}
