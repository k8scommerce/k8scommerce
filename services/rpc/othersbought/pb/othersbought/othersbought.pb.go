// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: othersbought.proto

package othersbought

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	catalog "github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"
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
type GetOthersBoughtBySkuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId int64  `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Sku     string `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *GetOthersBoughtBySkuRequest) Reset() {
	*x = GetOthersBoughtBySkuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_othersbought_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOthersBoughtBySkuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOthersBoughtBySkuRequest) ProtoMessage() {}

func (x *GetOthersBoughtBySkuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_othersbought_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOthersBoughtBySkuRequest.ProtoReflect.Descriptor instead.
func (*GetOthersBoughtBySkuRequest) Descriptor() ([]byte, []int) {
	return file_othersbought_proto_rawDescGZIP(), []int{0}
}

func (x *GetOthersBoughtBySkuRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *GetOthersBoughtBySkuRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type GetOthersBoughtBySkuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variants      []*catalog.Variant `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
	StatusCode    int64              `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string             `protobuf:"bytes,3,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *GetOthersBoughtBySkuResponse) Reset() {
	*x = GetOthersBoughtBySkuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_othersbought_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOthersBoughtBySkuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOthersBoughtBySkuResponse) ProtoMessage() {}

func (x *GetOthersBoughtBySkuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_othersbought_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOthersBoughtBySkuResponse.ProtoReflect.Descriptor instead.
func (*GetOthersBoughtBySkuResponse) Descriptor() ([]byte, []int) {
	return file_othersbought_proto_rawDescGZIP(), []int{1}
}

func (x *GetOthersBoughtBySkuResponse) GetVariants() []*catalog.Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *GetOthersBoughtBySkuResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetOthersBoughtBySkuResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

var File_othersbought_proto protoreflect.FileDescriptor

var file_othersbought_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x62, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x1a, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x42, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b,
	0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x92, 0x01, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x42, 0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x83, 0x01, 0x0a, 0x12, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x42, 0x79, 0x53, 0x6b, 0x75,
	0x12, 0x29, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x42,
	0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x62, 0x2f, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_othersbought_proto_rawDescOnce sync.Once
	file_othersbought_proto_rawDescData = file_othersbought_proto_rawDesc
)

func file_othersbought_proto_rawDescGZIP() []byte {
	file_othersbought_proto_rawDescOnce.Do(func() {
		file_othersbought_proto_rawDescData = protoimpl.X.CompressGZIP(file_othersbought_proto_rawDescData)
	})
	return file_othersbought_proto_rawDescData
}

var file_othersbought_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_othersbought_proto_goTypes = []interface{}{
	(*GetOthersBoughtBySkuRequest)(nil),  // 0: othersbought.GetOthersBoughtBySkuRequest
	(*GetOthersBoughtBySkuResponse)(nil), // 1: othersbought.GetOthersBoughtBySkuResponse
	(*catalog.Variant)(nil),              // 2: catalog.Variant
}
var file_othersbought_proto_depIdxs = []int32{
	2, // 0: othersbought.GetOthersBoughtBySkuResponse.variants:type_name -> catalog.Variant
	0, // 1: othersbought.OthersBoughtClient.GetOthersBoughtBySku:input_type -> othersbought.GetOthersBoughtBySkuRequest
	1, // 2: othersbought.OthersBoughtClient.GetOthersBoughtBySku:output_type -> othersbought.GetOthersBoughtBySkuResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_othersbought_proto_init() }
func file_othersbought_proto_init() {
	if File_othersbought_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_othersbought_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOthersBoughtBySkuRequest); i {
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
		file_othersbought_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOthersBoughtBySkuResponse); i {
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
			RawDescriptor: file_othersbought_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_othersbought_proto_goTypes,
		DependencyIndexes: file_othersbought_proto_depIdxs,
		MessageInfos:      file_othersbought_proto_msgTypes,
	}.Build()
	File_othersbought_proto = out.File
	file_othersbought_proto_rawDesc = nil
	file_othersbought_proto_goTypes = nil
	file_othersbought_proto_depIdxs = nil
}
