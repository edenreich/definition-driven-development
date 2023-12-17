//
//Pets API
//
//A simple API to manage pets
//
//The version of the OpenAPI document: 1.0.0
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: services/dogs_service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetDogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Dog `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDogsResponse) Reset() {
	*x = GetDogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_dogs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDogsResponse) ProtoMessage() {}

func (x *GetDogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_dogs_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDogsResponse.ProtoReflect.Descriptor instead.
func (*GetDogsResponse) Descriptor() ([]byte, []int) {
	return file_services_dogs_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDogsResponse) GetData() []*Dog {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_services_dogs_service_proto protoreflect.FileDescriptor

var file_services_dogs_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x67, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x67,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x5c, 0x0a, 0x0b, 0x44, 0x6f, 0x67,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x67, 0x73,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x72, 0x65, 0x69, 0x63, 0x68, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x6e, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_dogs_service_proto_rawDescOnce sync.Once
	file_services_dogs_service_proto_rawDescData = file_services_dogs_service_proto_rawDesc
)

func file_services_dogs_service_proto_rawDescGZIP() []byte {
	file_services_dogs_service_proto_rawDescOnce.Do(func() {
		file_services_dogs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_dogs_service_proto_rawDescData)
	})
	return file_services_dogs_service_proto_rawDescData
}

var file_services_dogs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_dogs_service_proto_goTypes = []interface{}{
	(*GetDogsResponse)(nil), // 0: grpc.services.dogsservice.GetDogsResponse
	(*Dog)(nil),             // 1: grpc.Dog
	(*emptypb.Empty)(nil),   // 2: google.protobuf.Empty
}
var file_services_dogs_service_proto_depIdxs = []int32{
	1, // 0: grpc.services.dogsservice.GetDogsResponse.data:type_name -> grpc.Dog
	2, // 1: grpc.services.dogsservice.DogsService.GetDogs:input_type -> google.protobuf.Empty
	0, // 2: grpc.services.dogsservice.DogsService.GetDogs:output_type -> grpc.services.dogsservice.GetDogsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_dogs_service_proto_init() }
func file_services_dogs_service_proto_init() {
	if File_services_dogs_service_proto != nil {
		return
	}
	file_models_dog_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_dogs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDogsResponse); i {
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
			RawDescriptor: file_services_dogs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_dogs_service_proto_goTypes,
		DependencyIndexes: file_services_dogs_service_proto_depIdxs,
		MessageInfos:      file_services_dogs_service_proto_msgTypes,
	}.Build()
	File_services_dogs_service_proto = out.File
	file_services_dogs_service_proto_rawDesc = nil
	file_services_dogs_service_proto_goTypes = nil
	file_services_dogs_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DogsServiceClient is the client API for DogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DogsServiceClient interface {
	GetDogs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDogsResponse, error)
}

type dogsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDogsServiceClient(cc grpc.ClientConnInterface) DogsServiceClient {
	return &dogsServiceClient{cc}
}

func (c *dogsServiceClient) GetDogs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDogsResponse, error) {
	out := new(GetDogsResponse)
	err := c.cc.Invoke(ctx, "/grpc.services.dogsservice.DogsService/GetDogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DogsServiceServer is the server API for DogsService service.
type DogsServiceServer interface {
	GetDogs(context.Context, *emptypb.Empty) (*GetDogsResponse, error)
}

// UnimplementedDogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDogsServiceServer struct {
}

func (*UnimplementedDogsServiceServer) GetDogs(context.Context, *emptypb.Empty) (*GetDogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDogs not implemented")
}

func RegisterDogsServiceServer(s *grpc.Server, srv DogsServiceServer) {
	s.RegisterService(&_DogsService_serviceDesc, srv)
}

func _DogsService_GetDogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DogsServiceServer).GetDogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.services.dogsservice.DogsService/GetDogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DogsServiceServer).GetDogs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.services.dogsservice.DogsService",
	HandlerType: (*DogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDogs",
			Handler:    _DogsService_GetDogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/dogs_service.proto",
}
