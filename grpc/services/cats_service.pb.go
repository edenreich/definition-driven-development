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
// source: services/cats_service.proto

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

type CreateCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedCat *CreatedCat `protobuf:"bytes,1,opt,name=createdCat,proto3" json:"createdCat,omitempty"`
}

func (x *CreateCatRequest) Reset() {
	*x = CreateCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cats_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatRequest) ProtoMessage() {}

func (x *CreateCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cats_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatRequest.ProtoReflect.Descriptor instead.
func (*CreateCatRequest) Descriptor() ([]byte, []int) {
	return file_services_cats_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCatRequest) GetCreatedCat() *CreatedCat {
	if x != nil {
		return x.CreatedCat
	}
	return nil
}

type GetCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCatRequest) Reset() {
	*x = GetCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cats_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatRequest) ProtoMessage() {}

func (x *GetCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cats_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatRequest.ProtoReflect.Descriptor instead.
func (*GetCatRequest) Descriptor() ([]byte, []int) {
	return file_services_cats_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Cat `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCatsResponse) Reset() {
	*x = GetCatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cats_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatsResponse) ProtoMessage() {}

func (x *GetCatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cats_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatsResponse.ProtoReflect.Descriptor instead.
func (*GetCatsResponse) Descriptor() ([]byte, []int) {
	return file_services_cats_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCatsResponse) GetData() []*Cat {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedCat *UpdatedCat `protobuf:"bytes,2,opt,name=updatedCat,proto3" json:"updatedCat,omitempty"`
}

func (x *UpdateCatRequest) Reset() {
	*x = UpdateCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cats_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatRequest) ProtoMessage() {}

func (x *UpdateCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cats_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatRequest) Descriptor() ([]byte, []int) {
	return file_services_cats_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCatRequest) GetUpdatedCat() *UpdatedCat {
	if x != nil {
		return x.UpdatedCat
	}
	return nil
}

var File_services_cats_service_proto protoreflect.FileDescriptor

var file_services_cats_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x61, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61,
	0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x74, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x74, 0x32, 0xcc, 0x02, 0x0a, 0x0b, 0x43,
	0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x73,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x72, 0x65, 0x69, 0x63,
	0x68, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x6e, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x50, 0x02, 0x50, 0x03, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_cats_service_proto_rawDescOnce sync.Once
	file_services_cats_service_proto_rawDescData = file_services_cats_service_proto_rawDesc
)

func file_services_cats_service_proto_rawDescGZIP() []byte {
	file_services_cats_service_proto_rawDescOnce.Do(func() {
		file_services_cats_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_cats_service_proto_rawDescData)
	})
	return file_services_cats_service_proto_rawDescData
}

var file_services_cats_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_cats_service_proto_goTypes = []interface{}{
	(*CreateCatRequest)(nil), // 0: grpc.services.catsservice.CreateCatRequest
	(*GetCatRequest)(nil),    // 1: grpc.services.catsservice.GetCatRequest
	(*GetCatsResponse)(nil),  // 2: grpc.services.catsservice.GetCatsResponse
	(*UpdateCatRequest)(nil), // 3: grpc.services.catsservice.UpdateCatRequest
	(*CreatedCat)(nil),       // 4: grpc.CreatedCat
	(*Cat)(nil),              // 5: grpc.Cat
	(*UpdatedCat)(nil),       // 6: grpc.UpdatedCat
	(*emptypb.Empty)(nil),    // 7: google.protobuf.Empty
}
var file_services_cats_service_proto_depIdxs = []int32{
	4, // 0: grpc.services.catsservice.CreateCatRequest.createdCat:type_name -> grpc.CreatedCat
	5, // 1: grpc.services.catsservice.GetCatsResponse.data:type_name -> grpc.Cat
	6, // 2: grpc.services.catsservice.UpdateCatRequest.updatedCat:type_name -> grpc.UpdatedCat
	0, // 3: grpc.services.catsservice.CatsService.CreateCat:input_type -> grpc.services.catsservice.CreateCatRequest
	1, // 4: grpc.services.catsservice.CatsService.GetCat:input_type -> grpc.services.catsservice.GetCatRequest
	7, // 5: grpc.services.catsservice.CatsService.GetCats:input_type -> google.protobuf.Empty
	3, // 6: grpc.services.catsservice.CatsService.UpdateCat:input_type -> grpc.services.catsservice.UpdateCatRequest
	7, // 7: grpc.services.catsservice.CatsService.CreateCat:output_type -> google.protobuf.Empty
	7, // 8: grpc.services.catsservice.CatsService.GetCat:output_type -> google.protobuf.Empty
	2, // 9: grpc.services.catsservice.CatsService.GetCats:output_type -> grpc.services.catsservice.GetCatsResponse
	7, // 10: grpc.services.catsservice.CatsService.UpdateCat:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_cats_service_proto_init() }
func file_services_cats_service_proto_init() {
	if File_services_cats_service_proto != nil {
		return
	}
	file_models_cat_proto_init()
	file_models_created_cat_proto_init()
	file_models_updated_cat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_cats_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatRequest); i {
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
		file_services_cats_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatRequest); i {
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
		file_services_cats_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatsResponse); i {
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
		file_services_cats_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatRequest); i {
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
			RawDescriptor: file_services_cats_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_cats_service_proto_goTypes,
		DependencyIndexes: file_services_cats_service_proto_depIdxs,
		MessageInfos:      file_services_cats_service_proto_msgTypes,
	}.Build()
	File_services_cats_service_proto = out.File
	file_services_cats_service_proto_rawDesc = nil
	file_services_cats_service_proto_goTypes = nil
	file_services_cats_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatsServiceClient is the client API for CatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatsServiceClient interface {
	CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCatsResponse, error)
	UpdateCat(ctx context.Context, in *UpdateCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type catsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatsServiceClient(cc grpc.ClientConnInterface) CatsServiceClient {
	return &catsServiceClient{cc}
}

func (c *catsServiceClient) CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/grpc.services.catsservice.CatsService/CreateCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/grpc.services.catsservice.CatsService/GetCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) GetCats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCatsResponse, error) {
	out := new(GetCatsResponse)
	err := c.cc.Invoke(ctx, "/grpc.services.catsservice.CatsService/GetCats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) UpdateCat(ctx context.Context, in *UpdateCatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/grpc.services.catsservice.CatsService/UpdateCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatsServiceServer is the server API for CatsService service.
type CatsServiceServer interface {
	CreateCat(context.Context, *CreateCatRequest) (*emptypb.Empty, error)
	GetCat(context.Context, *GetCatRequest) (*emptypb.Empty, error)
	GetCats(context.Context, *emptypb.Empty) (*GetCatsResponse, error)
	UpdateCat(context.Context, *UpdateCatRequest) (*emptypb.Empty, error)
}

// UnimplementedCatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatsServiceServer struct {
}

func (*UnimplementedCatsServiceServer) CreateCat(context.Context, *CreateCatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCat not implemented")
}
func (*UnimplementedCatsServiceServer) GetCat(context.Context, *GetCatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCat not implemented")
}
func (*UnimplementedCatsServiceServer) GetCats(context.Context, *emptypb.Empty) (*GetCatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCats not implemented")
}
func (*UnimplementedCatsServiceServer) UpdateCat(context.Context, *UpdateCatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCat not implemented")
}

func RegisterCatsServiceServer(s *grpc.Server, srv CatsServiceServer) {
	s.RegisterService(&_CatsService_serviceDesc, srv)
}

func _CatsService_CreateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).CreateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.services.catsservice.CatsService/CreateCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).CreateCat(ctx, req.(*CreateCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_GetCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).GetCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.services.catsservice.CatsService/GetCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).GetCat(ctx, req.(*GetCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_GetCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).GetCats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.services.catsservice.CatsService/GetCats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).GetCats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_UpdateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).UpdateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.services.catsservice.CatsService/UpdateCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).UpdateCat(ctx, req.(*UpdateCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.services.catsservice.CatsService",
	HandlerType: (*CatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCat",
			Handler:    _CatsService_CreateCat_Handler,
		},
		{
			MethodName: "GetCat",
			Handler:    _CatsService_GetCat_Handler,
		},
		{
			MethodName: "GetCats",
			Handler:    _CatsService_GetCats_Handler,
		},
		{
			MethodName: "UpdateCat",
			Handler:    _CatsService_UpdateCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/cats_service.proto",
}
