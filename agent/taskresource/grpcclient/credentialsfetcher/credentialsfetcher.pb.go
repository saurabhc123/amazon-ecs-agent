// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: credentialsfetcher/credentialsfetcher.proto

package credentialsfetcher

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

type CreateKerberosLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredspecContents []string `protobuf:"bytes,1,rep,name=credspec_contents,json=credspecContents,proto3" json:"credspec_contents,omitempty"`
}

func (x *CreateKerberosLeaseRequest) Reset() {
	*x = CreateKerberosLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKerberosLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKerberosLeaseRequest) ProtoMessage() {}

func (x *CreateKerberosLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKerberosLeaseRequest.ProtoReflect.Descriptor instead.
func (*CreateKerberosLeaseRequest) Descriptor() ([]byte, []int) {
	return file_credentialsfetcher_credentialsfetcher_proto_rawDescGZIP(), []int{0}
}

func (x *CreateKerberosLeaseRequest) GetCredspecContents() []string {
	if x != nil {
		return x.CredspecContents
	}
	return nil
}

type CreateKerberosLeaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId                  string   `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	CreatedKerberosFilePaths []string `protobuf:"bytes,2,rep,name=created_kerberos_file_paths,json=createdKerberosFilePaths,proto3" json:"created_kerberos_file_paths,omitempty"`
}

func (x *CreateKerberosLeaseResponse) Reset() {
	*x = CreateKerberosLeaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKerberosLeaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKerberosLeaseResponse) ProtoMessage() {}

func (x *CreateKerberosLeaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKerberosLeaseResponse.ProtoReflect.Descriptor instead.
func (*CreateKerberosLeaseResponse) Descriptor() ([]byte, []int) {
	return file_credentialsfetcher_credentialsfetcher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKerberosLeaseResponse) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *CreateKerberosLeaseResponse) GetCreatedKerberosFilePaths() []string {
	if x != nil {
		return x.CreatedKerberosFilePaths
	}
	return nil
}

type DeleteKerberosLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
}

func (x *DeleteKerberosLeaseRequest) Reset() {
	*x = DeleteKerberosLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKerberosLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKerberosLeaseRequest) ProtoMessage() {}

func (x *DeleteKerberosLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKerberosLeaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteKerberosLeaseRequest) Descriptor() ([]byte, []int) {
	return file_credentialsfetcher_credentialsfetcher_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteKerberosLeaseRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

type DeleteKerberosLeaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId                  string   `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	DeletedKerberosFilePaths []string `protobuf:"bytes,2,rep,name=deleted_kerberos_file_paths,json=deletedKerberosFilePaths,proto3" json:"deleted_kerberos_file_paths,omitempty"`
}

func (x *DeleteKerberosLeaseResponse) Reset() {
	*x = DeleteKerberosLeaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKerberosLeaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKerberosLeaseResponse) ProtoMessage() {}

func (x *DeleteKerberosLeaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credentialsfetcher_credentialsfetcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKerberosLeaseResponse.ProtoReflect.Descriptor instead.
func (*DeleteKerberosLeaseResponse) Descriptor() ([]byte, []int) {
	return file_credentialsfetcher_credentialsfetcher_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteKerberosLeaseResponse) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *DeleteKerberosLeaseResponse) GetDeletedKerberosFilePaths() []string {
	if x != nil {
		return x.DeletedKerberosFilePaths
	}
	return nil
}

var File_credentialsfetcher_credentialsfetcher_proto protoreflect.FileDescriptor

var file_credentialsfetcher_credentialsfetcher_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x22, 0x49, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65,
	0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64,
	0x73, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0x37, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b,
	0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x77,
	0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73,
	0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x32, 0x88, 0x02, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x72, 0x62,
	0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x72, 0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73,
	0x65, 0x12, 0x2e, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x72,
	0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x72,
	0x62, 0x65, 0x72, 0x6f, 0x73, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credentialsfetcher_credentialsfetcher_proto_rawDescOnce sync.Once
	file_credentialsfetcher_credentialsfetcher_proto_rawDescData = file_credentialsfetcher_credentialsfetcher_proto_rawDesc
)

func file_credentialsfetcher_credentialsfetcher_proto_rawDescGZIP() []byte {
	file_credentialsfetcher_credentialsfetcher_proto_rawDescOnce.Do(func() {
		file_credentialsfetcher_credentialsfetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_credentialsfetcher_credentialsfetcher_proto_rawDescData)
	})
	return file_credentialsfetcher_credentialsfetcher_proto_rawDescData
}

var file_credentialsfetcher_credentialsfetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_credentialsfetcher_credentialsfetcher_proto_goTypes = []interface{}{
	(*CreateKerberosLeaseRequest)(nil),  // 0: credentialsfetcher.CreateKerberosLeaseRequest
	(*CreateKerberosLeaseResponse)(nil), // 1: credentialsfetcher.CreateKerberosLeaseResponse
	(*DeleteKerberosLeaseRequest)(nil),  // 2: credentialsfetcher.DeleteKerberosLeaseRequest
	(*DeleteKerberosLeaseResponse)(nil), // 3: credentialsfetcher.DeleteKerberosLeaseResponse
}
var file_credentialsfetcher_credentialsfetcher_proto_depIdxs = []int32{
	0, // 0: credentialsfetcher.CredentialsFetcherService.AddKerberosLease:input_type -> credentialsfetcher.CreateKerberosLeaseRequest
	2, // 1: credentialsfetcher.CredentialsFetcherService.DeleteKerberosLease:input_type -> credentialsfetcher.DeleteKerberosLeaseRequest
	1, // 2: credentialsfetcher.CredentialsFetcherService.AddKerberosLease:output_type -> credentialsfetcher.CreateKerberosLeaseResponse
	3, // 3: credentialsfetcher.CredentialsFetcherService.DeleteKerberosLease:output_type -> credentialsfetcher.DeleteKerberosLeaseResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_credentialsfetcher_credentialsfetcher_proto_init() }
func file_credentialsfetcher_credentialsfetcher_proto_init() {
	if File_credentialsfetcher_credentialsfetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credentialsfetcher_credentialsfetcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKerberosLeaseRequest); i {
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
		file_credentialsfetcher_credentialsfetcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKerberosLeaseResponse); i {
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
		file_credentialsfetcher_credentialsfetcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKerberosLeaseRequest); i {
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
		file_credentialsfetcher_credentialsfetcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKerberosLeaseResponse); i {
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
			RawDescriptor: file_credentialsfetcher_credentialsfetcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credentialsfetcher_credentialsfetcher_proto_goTypes,
		DependencyIndexes: file_credentialsfetcher_credentialsfetcher_proto_depIdxs,
		MessageInfos:      file_credentialsfetcher_credentialsfetcher_proto_msgTypes,
	}.Build()
	File_credentialsfetcher_credentialsfetcher_proto = out.File
	file_credentialsfetcher_credentialsfetcher_proto_rawDesc = nil
	file_credentialsfetcher_credentialsfetcher_proto_goTypes = nil
	file_credentialsfetcher_credentialsfetcher_proto_depIdxs = nil
}
