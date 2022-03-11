// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: inventory.proto

package inventory

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

type InventoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId    int64  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Sku        string `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	SupplierId int64  `protobuf:"varint,5,opt,name=supplierId,proto3" json:"supplierId,omitempty"`
	BrandId    int64  `protobuf:"varint,6,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *InventoryItem) Reset() {
	*x = InventoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItem) ProtoMessage() {}

func (x *InventoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItem.ProtoReflect.Descriptor instead.
func (*InventoryItem) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InventoryItem) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *InventoryItem) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *InventoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InventoryItem) GetSupplierId() int64 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *InventoryItem) GetBrandId() int64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId int64 `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *Warehouse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type SockLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId         int64 `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	WarehouseId     int64 `protobuf:"varint,2,opt,name=warehouseId,proto3" json:"warehouseId,omitempty"`
	InventoryItemId int64 `protobuf:"varint,3,opt,name=inventoryItemId,proto3" json:"inventoryItemId,omitempty"`
	Quantity        int32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *SockLevel) Reset() {
	*x = SockLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SockLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SockLevel) ProtoMessage() {}

func (x *SockLevel) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SockLevel.ProtoReflect.Descriptor instead.
func (*SockLevel) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *SockLevel) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *SockLevel) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *SockLevel) GetInventoryItemId() int64 {
	if x != nil {
		return x.InventoryItemId
	}
	return 0
}

func (x *SockLevel) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId int64  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *Supplier) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Supplier) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *Supplier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Brand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId int64  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Brand) Reset() {
	*x = Brand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brand) ProtoMessage() {}

func (x *Brand) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brand.ProtoReflect.Descriptor instead.
func (*Brand) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *Brand) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Brand) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *Brand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// get item quantity
type GetItemQuantityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId int64  `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Sku     string `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Region  string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *GetItemQuantityRequest) Reset() {
	*x = GetItemQuantityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemQuantityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemQuantityRequest) ProtoMessage() {}

func (x *GetItemQuantityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemQuantityRequest.ProtoReflect.Descriptor instead.
func (*GetItemQuantityRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *GetItemQuantityRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *GetItemQuantityRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *GetItemQuantityRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type GetItemQuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId       int64  `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Sku           string `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Quanity       int32  `protobuf:"varint,3,opt,name=quanity,proto3" json:"quanity,omitempty"`
	StatusCode    int64  `protobuf:"varint,4,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string `protobuf:"bytes,5,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *GetItemQuantityResponse) Reset() {
	*x = GetItemQuantityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemQuantityResponse) ProtoMessage() {}

func (x *GetItemQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemQuantityResponse.ProtoReflect.Descriptor instead.
func (*GetItemQuantityResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *GetItemQuantityResponse) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *GetItemQuantityResponse) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *GetItemQuantityResponse) GetQuanity() int32 {
	if x != nil {
		return x.Quanity
	}
	return 0
}

func (x *GetItemQuantityResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetItemQuantityResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x99, 0x01, 0x0a,
	0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x09, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22,
	0x8d, 0x01, 0x0a, 0x09, 0x53, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x48, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0xa5,
	0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6e, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6e, 0x69, 0x74, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6b, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_inventory_proto_goTypes = []interface{}{
	(*InventoryItem)(nil),           // 0: inventory.InventoryItem
	(*Warehouse)(nil),               // 1: inventory.Warehouse
	(*SockLevel)(nil),               // 2: inventory.SockLevel
	(*Supplier)(nil),                // 3: inventory.Supplier
	(*Brand)(nil),                   // 4: inventory.Brand
	(*GetItemQuantityRequest)(nil),  // 5: inventory.GetItemQuantityRequest
	(*GetItemQuantityResponse)(nil), // 6: inventory.GetItemQuantityResponse
}
var file_inventory_proto_depIdxs = []int32{
	5, // 0: inventory.InventoryClient.GetItemQuantity:input_type -> inventory.GetItemQuantityRequest
	6, // 1: inventory.InventoryClient.GetItemQuantity:output_type -> inventory.GetItemQuantityResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryItem); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouse); i {
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
		file_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SockLevel); i {
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
		file_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
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
		file_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brand); i {
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
		file_inventory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemQuantityRequest); i {
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
		file_inventory_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemQuantityResponse); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InventoryClientClient is the client API for InventoryClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryClientClient interface {
	GetItemQuantity(ctx context.Context, in *GetItemQuantityRequest, opts ...grpc.CallOption) (*GetItemQuantityResponse, error)
}

type inventoryClientClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClientClient(cc grpc.ClientConnInterface) InventoryClientClient {
	return &inventoryClientClient{cc}
}

func (c *inventoryClientClient) GetItemQuantity(ctx context.Context, in *GetItemQuantityRequest, opts ...grpc.CallOption) (*GetItemQuantityResponse, error) {
	out := new(GetItemQuantityResponse)
	err := c.cc.Invoke(ctx, "/inventory.InventoryClient/GetItemQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryClientServer is the server API for InventoryClient service.
type InventoryClientServer interface {
	GetItemQuantity(context.Context, *GetItemQuantityRequest) (*GetItemQuantityResponse, error)
}

// UnimplementedInventoryClientServer can be embedded to have forward compatible implementations.
type UnimplementedInventoryClientServer struct {
}

func (*UnimplementedInventoryClientServer) GetItemQuantity(context.Context, *GetItemQuantityRequest) (*GetItemQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemQuantity not implemented")
}

func RegisterInventoryClientServer(s *grpc.Server, srv InventoryClientServer) {
	s.RegisterService(&_InventoryClient_serviceDesc, srv)
}

func _InventoryClient_GetItemQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryClientServer).GetItemQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryClient/GetItemQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryClientServer).GetItemQuantity(ctx, req.(*GetItemQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InventoryClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryClient",
	HandlerType: (*InventoryClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItemQuantity",
			Handler:    _InventoryClient_GetItemQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
