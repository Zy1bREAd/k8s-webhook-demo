// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/containers.proto

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

type ContainerInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image         string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	mi := &file_proto_containers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_containers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_proto_containers_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListContainersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 能被List出Container的数量
	ContainerCount string `protobuf:"bytes,1,opt,name=container_count,json=containerCount,proto3" json:"container_count,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListContainersRequest) Reset() {
	*x = ListContainersRequest{}
	mi := &file_proto_containers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContainersRequest) ProtoMessage() {}

func (x *ListContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_containers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContainersRequest.ProtoReflect.Descriptor instead.
func (*ListContainersRequest) Descriptor() ([]byte, []int) {
	return file_proto_containers_proto_rawDescGZIP(), []int{1}
}

func (x *ListContainersRequest) GetContainerCount() string {
	if x != nil {
		return x.ContainerCount
	}
	return ""
}

type GetContainerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 能被List出Container的数量
	ContainerId   string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContainerRequest) Reset() {
	*x = GetContainerRequest{}
	mi := &file_proto_containers_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerRequest) ProtoMessage() {}

func (x *GetContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_containers_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerRequest.ProtoReflect.Descriptor instead.
func (*GetContainerRequest) Descriptor() ([]byte, []int) {
	return file_proto_containers_proto_rawDescGZIP(), []int{2}
}

func (x *GetContainerRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type ListContainersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The field name should match the noun "Type" in the method name.
	// There will be a maximum number of items returned based on the page_size field in the request.
	ContainerList []*ContainerInfo `protobuf:"bytes,1,rep,name=ContainerList,proto3" json:"ContainerList,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListContainersResponse) Reset() {
	*x = ListContainersResponse{}
	mi := &file_proto_containers_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContainersResponse) ProtoMessage() {}

func (x *ListContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_containers_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContainersResponse.ProtoReflect.Descriptor instead.
func (*ListContainersResponse) Descriptor() ([]byte, []int) {
	return file_proto_containers_proto_rawDescGZIP(), []int{3}
}

func (x *ListContainersResponse) GetContainerList() []*ContainerInfo {
	if x != nil {
		return x.ContainerList
	}
	return nil
}

func (x *ListContainersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetContainerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerList *ContainerInfo         `protobuf:"bytes,1,opt,name=ContainerList,proto3" json:"ContainerList,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContainerResponse) Reset() {
	*x = GetContainerResponse{}
	mi := &file_proto_containers_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerResponse) ProtoMessage() {}

func (x *GetContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_containers_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerResponse.ProtoReflect.Descriptor instead.
func (*GetContainerResponse) Descriptor() ([]byte, []int) {
	return file_proto_containers_proto_rawDescGZIP(), []int{4}
}

func (x *GetContainerResponse) GetContainerList() *ContainerInfo {
	if x != nil {
		return x.ContainerList
	}
	return nil
}

var File_proto_containers_proto protoreflect.FileDescriptor

var file_proto_containers_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xb8, 0x01,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x55, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x20,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_containers_proto_rawDescOnce sync.Once
	file_proto_containers_proto_rawDescData = file_proto_containers_proto_rawDesc
)

func file_proto_containers_proto_rawDescGZIP() []byte {
	file_proto_containers_proto_rawDescOnce.Do(func() {
		file_proto_containers_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_containers_proto_rawDescData)
	})
	return file_proto_containers_proto_rawDescData
}

var file_proto_containers_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_containers_proto_goTypes = []any{
	(*ContainerInfo)(nil),          // 0: Container.ContainerInfo
	(*ListContainersRequest)(nil),  // 1: Container.ListContainersRequest
	(*GetContainerRequest)(nil),    // 2: Container.GetContainerRequest
	(*ListContainersResponse)(nil), // 3: Container.ListContainersResponse
	(*GetContainerResponse)(nil),   // 4: Container.GetContainerResponse
}
var file_proto_containers_proto_depIdxs = []int32{
	0, // 0: Container.ListContainersResponse.ContainerList:type_name -> Container.ContainerInfo
	0, // 1: Container.GetContainerResponse.ContainerList:type_name -> Container.ContainerInfo
	1, // 2: Container.Containers.ListContainers:input_type -> Container.ListContainersRequest
	2, // 3: Container.Containers.GetContainerInfo:input_type -> Container.GetContainerRequest
	3, // 4: Container.Containers.ListContainers:output_type -> Container.ListContainersResponse
	4, // 5: Container.Containers.GetContainerInfo:output_type -> Container.GetContainerResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_containers_proto_init() }
func file_proto_containers_proto_init() {
	if File_proto_containers_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_containers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_containers_proto_goTypes,
		DependencyIndexes: file_proto_containers_proto_depIdxs,
		MessageInfos:      file_proto_containers_proto_msgTypes,
	}.Build()
	File_proto_containers_proto = out.File
	file_proto_containers_proto_rawDesc = nil
	file_proto_containers_proto_goTypes = nil
	file_proto_containers_proto_depIdxs = nil
}
