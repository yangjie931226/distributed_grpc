// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: registry.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RegistionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//服务名
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	//服务地址
	ServiceUrl       string   `protobuf:"bytes,2,opt,name=serviceUrl,proto3" json:"serviceUrl,omitempty"`
	RequiresService  []string `protobuf:"bytes,3,rep,name=requiresService,proto3" json:"requiresService,omitempty"`
	ServiceUpdateUrl string   `protobuf:"bytes,4,opt,name=serviceUpdateUrl,proto3" json:"serviceUpdateUrl,omitempty"`
	HeartbeatUrl     string   `protobuf:"bytes,5,opt,name=heartbeatUrl,proto3" json:"heartbeatUrl,omitempty"`
}

func (x *RegistionRequest) Reset() {
	*x = RegistionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistionRequest) ProtoMessage() {}

func (x *RegistionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistionRequest.ProtoReflect.Descriptor instead.
func (*RegistionRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RegistionRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegistionRequest) GetServiceUrl() string {
	if x != nil {
		return x.ServiceUrl
	}
	return ""
}

func (x *RegistionRequest) GetRequiresService() []string {
	if x != nil {
		return x.RequiresService
	}
	return nil
}

func (x *RegistionRequest) GetServiceUpdateUrl() string {
	if x != nil {
		return x.ServiceUpdateUrl
	}
	return ""
}

func (x *RegistionRequest) GetHeartbeatUrl() string {
	if x != nil {
		return x.HeartbeatUrl
	}
	return ""
}

type RegistionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	//错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RegistionReply) Reset() {
	*x = RegistionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistionReply) ProtoMessage() {}

func (x *RegistionReply) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistionReply.ProtoReflect.Descriptor instead.
func (*RegistionReply) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegistionReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegistionReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_registry_proto protoreflect.FileDescriptor

var file_registry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xce, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x36, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x8c, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_proto_rawDescOnce sync.Once
	file_registry_proto_rawDescData = file_registry_proto_rawDesc
)

func file_registry_proto_rawDescGZIP() []byte {
	file_registry_proto_rawDescOnce.Do(func() {
		file_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_proto_rawDescData)
	})
	return file_registry_proto_rawDescData
}

var file_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_registry_proto_goTypes = []interface{}{
	(*RegistionRequest)(nil), // 0: pb.RegistionRequest
	(*RegistionReply)(nil),   // 1: pb.RegistionReply
}
var file_registry_proto_depIdxs = []int32{
	0, // 0: pb.RegistryService.AddRegistion:input_type -> pb.RegistionRequest
	0, // 1: pb.RegistryService.RemoveRegistion:input_type -> pb.RegistionRequest
	1, // 2: pb.RegistryService.AddRegistion:output_type -> pb.RegistionReply
	1, // 3: pb.RegistryService.RemoveRegistion:output_type -> pb.RegistionReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_registry_proto_init() }
func file_registry_proto_init() {
	if File_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistionRequest); i {
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
		file_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistionReply); i {
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
			RawDescriptor: file_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_proto_goTypes,
		DependencyIndexes: file_registry_proto_depIdxs,
		MessageInfos:      file_registry_proto_msgTypes,
	}.Build()
	File_registry_proto = out.File
	file_registry_proto_rawDesc = nil
	file_registry_proto_goTypes = nil
	file_registry_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistryServiceClient is the client API for RegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryServiceClient interface {
	AddRegistion(ctx context.Context, in *RegistionRequest, opts ...grpc.CallOption) (*RegistionReply, error)
	RemoveRegistion(ctx context.Context, in *RegistionRequest, opts ...grpc.CallOption) (*RegistionReply, error)
}

type registryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryServiceClient(cc grpc.ClientConnInterface) RegistryServiceClient {
	return &registryServiceClient{cc}
}

func (c *registryServiceClient) AddRegistion(ctx context.Context, in *RegistionRequest, opts ...grpc.CallOption) (*RegistionReply, error) {
	out := new(RegistionReply)
	err := c.cc.Invoke(ctx, "/pb.RegistryService/AddRegistion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) RemoveRegistion(ctx context.Context, in *RegistionRequest, opts ...grpc.CallOption) (*RegistionReply, error) {
	out := new(RegistionReply)
	err := c.cc.Invoke(ctx, "/pb.RegistryService/RemoveRegistion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServiceServer is the server API for RegistryService service.
type RegistryServiceServer interface {
	AddRegistion(context.Context, *RegistionRequest) (*RegistionReply, error)
	RemoveRegistion(context.Context, *RegistionRequest) (*RegistionReply, error)
}

// UnimplementedRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServiceServer struct {
}

func (*UnimplementedRegistryServiceServer) AddRegistion(context.Context, *RegistionRequest) (*RegistionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRegistion not implemented")
}
func (*UnimplementedRegistryServiceServer) RemoveRegistion(context.Context, *RegistionRequest) (*RegistionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRegistion not implemented")
}

func RegisterRegistryServiceServer(s *grpc.Server, srv RegistryServiceServer) {
	s.RegisterService(&_RegistryService_serviceDesc, srv)
}

func _RegistryService_AddRegistion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).AddRegistion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegistryService/AddRegistion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).AddRegistion(ctx, req.(*RegistionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_RemoveRegistion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).RemoveRegistion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegistryService/RemoveRegistion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).RemoveRegistion(ctx, req.(*RegistionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegistryService",
	HandlerType: (*RegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRegistion",
			Handler:    _RegistryService_AddRegistion_Handler,
		},
		{
			MethodName: "RemoveRegistion",
			Handler:    _RegistryService_RemoveRegistion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}
