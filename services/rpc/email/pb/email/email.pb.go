// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: email.proto

package email

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	cart "k8scommerce/services/rpc/cart/pb/cart"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId            int64        `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	OrderId            string       `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Item               []*cart.Item `protobuf:"bytes,3,rep,name=item,proto3" json:"item,omitempty"`
	SubTotal           int64        `protobuf:"varint,4,opt,name=subTotal,proto3" json:"subTotal,omitempty"`
	Tax                int64        `protobuf:"varint,5,opt,name=tax,proto3" json:"tax,omitempty"`
	Vat                int64        `protobuf:"varint,6,opt,name=vat,proto3" json:"vat,omitempty"`
	Tax2               int64        `protobuf:"varint,7,opt,name=tax2,proto3" json:"tax2,omitempty"`
	Tax3               int64        `protobuf:"varint,8,opt,name=tax3,proto3" json:"tax3,omitempty"`
	ShippingAmount     int64        `protobuf:"varint,9,opt,name=shippingAmount,proto3" json:"shippingAmount,omitempty"`
	TotalAmount        int64        `protobuf:"varint,10,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Currency           string       `protobuf:"bytes,11,opt,name=currency,proto3" json:"currency,omitempty"`
	ShippingTrackingId string       `protobuf:"bytes,12,opt,name=shippingTrackingId,proto3" json:"shippingTrackingId,omitempty"`
	ShippingCarrier    string       `protobuf:"bytes,13,opt,name=shippingCarrier,proto3" json:"shippingCarrier,omitempty"`
	ShippingAddressId  string       `protobuf:"bytes,14,opt,name=shippingAddressId,proto3" json:"shippingAddressId,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetItem() []*cart.Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *Order) GetSubTotal() int64 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *Order) GetTax() int64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Order) GetVat() int64 {
	if x != nil {
		return x.Vat
	}
	return 0
}

func (x *Order) GetTax2() int64 {
	if x != nil {
		return x.Tax2
	}
	return 0
}

func (x *Order) GetTax3() int64 {
	if x != nil {
		return x.Tax3
	}
	return 0
}

func (x *Order) GetShippingAmount() int64 {
	if x != nil {
		return x.ShippingAmount
	}
	return 0
}

func (x *Order) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Order) GetShippingTrackingId() string {
	if x != nil {
		return x.ShippingTrackingId
	}
	return ""
}

func (x *Order) GetShippingCarrier() string {
	if x != nil {
		return x.ShippingCarrier
	}
	return ""
}

func (x *Order) GetShippingAddressId() string {
	if x != nil {
		return x.ShippingAddressId
	}
	return ""
}

type SendOrderConfirmationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Order *Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SendOrderConfirmationRequest) Reset() {
	*x = SendOrderConfirmationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderConfirmationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderConfirmationRequest) ProtoMessage() {}

func (x *SendOrderConfirmationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderConfirmationRequest.ProtoReflect.Descriptor instead.
func (*SendOrderConfirmationRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendOrderConfirmationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendOrderConfirmationRequest) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xb1, 0x03, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x76, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x78, 0x32, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x61, 0x78, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x78, 0x33, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x61, 0x78, 0x33, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x11, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x58, 0x0a,
	0x1c, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x5b, 0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x62, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_email_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: email.Empty
	(*Order)(nil),                        // 1: email.Order
	(*SendOrderConfirmationRequest)(nil), // 2: email.SendOrderConfirmationRequest
	(*cart.Item)(nil),                    // 3: cart.Item
}
var file_email_proto_depIdxs = []int32{
	3, // 0: email.Order.item:type_name -> cart.Item
	1, // 1: email.SendOrderConfirmationRequest.order:type_name -> email.Order
	2, // 2: email.EmailClient.SendOrderConfirmation:input_type -> email.SendOrderConfirmationRequest
	0, // 3: email.EmailClient.SendOrderConfirmation:output_type -> email.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderConfirmationRequest); i {
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
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailClientClient is the client API for EmailClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClientClient interface {
	SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type emailClientClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClientClient(cc grpc.ClientConnInterface) EmailClientClient {
	return &emailClientClient{cc}
}

func (c *emailClientClient) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/email.EmailClient/SendOrderConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailClientServer is the server API for EmailClient service.
type EmailClientServer interface {
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error)
}

// UnimplementedEmailClientServer can be embedded to have forward compatible implementations.
type UnimplementedEmailClientServer struct {
}

func (*UnimplementedEmailClientServer) SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrderConfirmation not implemented")
}

func RegisterEmailClientServer(s *grpc.Server, srv EmailClientServer) {
	s.RegisterService(&_EmailClient_serviceDesc, srv)
}

func _EmailClient_SendOrderConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailClientServer).SendOrderConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.EmailClient/SendOrderConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailClientServer).SendOrderConfirmation(ctx, req.(*SendOrderConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.EmailClient",
	HandlerType: (*EmailClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrderConfirmation",
			Handler:    _EmailClient_SendOrderConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
