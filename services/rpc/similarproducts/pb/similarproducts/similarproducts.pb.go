// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: similarproducts.proto

package similarproducts

import (
	context "context"
	product "github.com/k8s-commerce/k8s-commerce/services/rpc/product/pb/product"
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

// get one by sku
type GetSimilarProductsBySkuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *GetSimilarProductsBySkuRequest) Reset() {
	*x = GetSimilarProductsBySkuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_similarproducts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimilarProductsBySkuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimilarProductsBySkuRequest) ProtoMessage() {}

func (x *GetSimilarProductsBySkuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_similarproducts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimilarProductsBySkuRequest.ProtoReflect.Descriptor instead.
func (*GetSimilarProductsBySkuRequest) Descriptor() ([]byte, []int) {
	return file_similarproducts_proto_rawDescGZIP(), []int{0}
}

func (x *GetSimilarProductsBySkuRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type GetSimilarProductsBySkuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variants      []*product.Variant `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
	StatusCode    int64              `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string             `protobuf:"bytes,3,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *GetSimilarProductsBySkuResponse) Reset() {
	*x = GetSimilarProductsBySkuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_similarproducts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimilarProductsBySkuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimilarProductsBySkuResponse) ProtoMessage() {}

func (x *GetSimilarProductsBySkuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_similarproducts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimilarProductsBySkuResponse.ProtoReflect.Descriptor instead.
func (*GetSimilarProductsBySkuResponse) Descriptor() ([]byte, []int) {
	return file_similarproducts_proto_rawDescGZIP(), []int{1}
}

func (x *GetSimilarProductsBySkuResponse) GetVariants() []*product.Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *GetSimilarProductsBySkuResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetSimilarProductsBySkuResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

var File_similarproducts_proto protoreflect.FileDescriptor

var file_similarproducts_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53,
	0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x95, 0x01, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x15, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x7c, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x12, 0x2f, 0x2e, 0x73, 0x69, 0x6d, 0x69, 0x6c,
	0x61, 0x72, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53,
	0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x69, 0x6d, 0x69,
	0x6c, 0x61, 0x72, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79,
	0x53, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x70,
	0x62, 0x2f, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_similarproducts_proto_rawDescOnce sync.Once
	file_similarproducts_proto_rawDescData = file_similarproducts_proto_rawDesc
)

func file_similarproducts_proto_rawDescGZIP() []byte {
	file_similarproducts_proto_rawDescOnce.Do(func() {
		file_similarproducts_proto_rawDescData = protoimpl.X.CompressGZIP(file_similarproducts_proto_rawDescData)
	})
	return file_similarproducts_proto_rawDescData
}

var file_similarproducts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_similarproducts_proto_goTypes = []interface{}{
	(*GetSimilarProductsBySkuRequest)(nil),  // 0: similarproducts.GetSimilarProductsBySkuRequest
	(*GetSimilarProductsBySkuResponse)(nil), // 1: similarproducts.GetSimilarProductsBySkuResponse
	(*product.Variant)(nil),                 // 2: product.Variant
}
var file_similarproducts_proto_depIdxs = []int32{
	2, // 0: similarproducts.GetSimilarProductsBySkuResponse.variants:type_name -> product.Variant
	0, // 1: similarproducts.SimilarProductsClient.GetSimilarProductsBySku:input_type -> similarproducts.GetSimilarProductsBySkuRequest
	1, // 2: similarproducts.SimilarProductsClient.GetSimilarProductsBySku:output_type -> similarproducts.GetSimilarProductsBySkuResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_similarproducts_proto_init() }
func file_similarproducts_proto_init() {
	if File_similarproducts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_similarproducts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimilarProductsBySkuRequest); i {
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
		file_similarproducts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimilarProductsBySkuResponse); i {
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
			RawDescriptor: file_similarproducts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_similarproducts_proto_goTypes,
		DependencyIndexes: file_similarproducts_proto_depIdxs,
		MessageInfos:      file_similarproducts_proto_msgTypes,
	}.Build()
	File_similarproducts_proto = out.File
	file_similarproducts_proto_rawDesc = nil
	file_similarproducts_proto_goTypes = nil
	file_similarproducts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SimilarProductsClientClient is the client API for SimilarProductsClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimilarProductsClientClient interface {
	GetSimilarProductsBySku(ctx context.Context, in *GetSimilarProductsBySkuRequest, opts ...grpc.CallOption) (*GetSimilarProductsBySkuResponse, error)
}

type similarProductsClientClient struct {
	cc grpc.ClientConnInterface
}

func NewSimilarProductsClientClient(cc grpc.ClientConnInterface) SimilarProductsClientClient {
	return &similarProductsClientClient{cc}
}

func (c *similarProductsClientClient) GetSimilarProductsBySku(ctx context.Context, in *GetSimilarProductsBySkuRequest, opts ...grpc.CallOption) (*GetSimilarProductsBySkuResponse, error) {
	out := new(GetSimilarProductsBySkuResponse)
	err := c.cc.Invoke(ctx, "/similarproducts.SimilarProductsClient/GetSimilarProductsBySku", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimilarProductsClientServer is the server API for SimilarProductsClient service.
type SimilarProductsClientServer interface {
	GetSimilarProductsBySku(context.Context, *GetSimilarProductsBySkuRequest) (*GetSimilarProductsBySkuResponse, error)
}

// UnimplementedSimilarProductsClientServer can be embedded to have forward compatible implementations.
type UnimplementedSimilarProductsClientServer struct {
}

func (*UnimplementedSimilarProductsClientServer) GetSimilarProductsBySku(context.Context, *GetSimilarProductsBySkuRequest) (*GetSimilarProductsBySkuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimilarProductsBySku not implemented")
}

func RegisterSimilarProductsClientServer(s *grpc.Server, srv SimilarProductsClientServer) {
	s.RegisterService(&_SimilarProductsClient_serviceDesc, srv)
}

func _SimilarProductsClient_GetSimilarProductsBySku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimilarProductsBySkuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimilarProductsClientServer).GetSimilarProductsBySku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/similarproducts.SimilarProductsClient/GetSimilarProductsBySku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimilarProductsClientServer).GetSimilarProductsBySku(ctx, req.(*GetSimilarProductsBySkuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimilarProductsClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "similarproducts.SimilarProductsClient",
	HandlerType: (*SimilarProductsClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSimilarProductsBySku",
			Handler:    _SimilarProductsClient_GetSimilarProductsBySku_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "similarproducts.proto",
}
