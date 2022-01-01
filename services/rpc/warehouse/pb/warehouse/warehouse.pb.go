// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: warehouse.proto

package warehouse

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

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Street     string `protobuf:"bytes,3,opt,name=street,proto3" json:"street,omitempty"`
	City       string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State      string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Country    string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	PostalCode string `protobuf:"bytes,7,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Lat        string `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng        string `protobuf:"bytes,9,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[0]
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
	return file_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *Warehouse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warehouse) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Warehouse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Warehouse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Warehouse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Warehouse) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Warehouse) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Warehouse) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

// create Warehouse product
type CreateWarehouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId   int64      `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Warehouse *Warehouse `protobuf:"bytes,2,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
}

func (x *CreateWarehouseRequest) Reset() {
	*x = CreateWarehouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWarehouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWarehouseRequest) ProtoMessage() {}

func (x *CreateWarehouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWarehouseRequest.ProtoReflect.Descriptor instead.
func (*CreateWarehouseRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWarehouseRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *CreateWarehouseRequest) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

type CreateWarehouseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse     *Warehouse `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	StatusCode    int64      `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string     `protobuf:"bytes,3,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *CreateWarehouseResponse) Reset() {
	*x = CreateWarehouseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWarehouseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWarehouseResponse) ProtoMessage() {}

func (x *CreateWarehouseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWarehouseResponse.ProtoReflect.Descriptor instead.
func (*CreateWarehouseResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWarehouseResponse) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *CreateWarehouseResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateWarehouseResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

// get Get All Warehouses By Store Id
type GetAllWarehousesByStoreIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId int64 `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
}

func (x *GetAllWarehousesByStoreIdRequest) Reset() {
	*x = GetAllWarehousesByStoreIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllWarehousesByStoreIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllWarehousesByStoreIdRequest) ProtoMessage() {}

func (x *GetAllWarehousesByStoreIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllWarehousesByStoreIdRequest.ProtoReflect.Descriptor instead.
func (*GetAllWarehousesByStoreIdRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllWarehousesByStoreIdRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type GetAllWarehousesByStoreIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse     *Warehouse `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	StatusCode    int64      `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string     `protobuf:"bytes,3,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *GetAllWarehousesByStoreIdResponse) Reset() {
	*x = GetAllWarehousesByStoreIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllWarehousesByStoreIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllWarehousesByStoreIdResponse) ProtoMessage() {}

func (x *GetAllWarehousesByStoreIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllWarehousesByStoreIdResponse.ProtoReflect.Descriptor instead.
func (*GetAllWarehousesByStoreIdResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllWarehousesByStoreIdResponse) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *GetAllWarehousesByStoreIdResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetAllWarehousesByStoreIdResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

// get Warehouse by Id
type GetWarehouseByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWarehouseByIdRequest) Reset() {
	*x = GetWarehouseByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehouseByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehouseByIdRequest) ProtoMessage() {}

func (x *GetWarehouseByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehouseByIdRequest.ProtoReflect.Descriptor instead.
func (*GetWarehouseByIdRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{5}
}

func (x *GetWarehouseByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWarehouseByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse     *Warehouse `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	StatusCode    int64      `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string     `protobuf:"bytes,3,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *GetWarehouseByIdResponse) Reset() {
	*x = GetWarehouseByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehouseByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehouseByIdResponse) ProtoMessage() {}

func (x *GetWarehouseByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehouseByIdResponse.ProtoReflect.Descriptor instead.
func (*GetWarehouseByIdResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{6}
}

func (x *GetWarehouseByIdResponse) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *GetWarehouseByIdResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetWarehouseByIdResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

var File_warehouse_proto protoreflect.FileDescriptor

var file_warehouse_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x0a,
	0x09, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x67,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x42, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73,
	0x42, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc0, 0x02,
	0x0a, 0x0f, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x58, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x42,
	0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x42, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x73, 0x42, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x62, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_proto_rawDescOnce sync.Once
	file_warehouse_proto_rawDescData = file_warehouse_proto_rawDesc
)

func file_warehouse_proto_rawDescGZIP() []byte {
	file_warehouse_proto_rawDescOnce.Do(func() {
		file_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_proto_rawDescData)
	})
	return file_warehouse_proto_rawDescData
}

var file_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_warehouse_proto_goTypes = []interface{}{
	(*Warehouse)(nil),                         // 0: warehouse.Warehouse
	(*CreateWarehouseRequest)(nil),            // 1: warehouse.CreateWarehouseRequest
	(*CreateWarehouseResponse)(nil),           // 2: warehouse.CreateWarehouseResponse
	(*GetAllWarehousesByStoreIdRequest)(nil),  // 3: warehouse.GetAllWarehousesByStoreIdRequest
	(*GetAllWarehousesByStoreIdResponse)(nil), // 4: warehouse.GetAllWarehousesByStoreIdResponse
	(*GetWarehouseByIdRequest)(nil),           // 5: warehouse.GetWarehouseByIdRequest
	(*GetWarehouseByIdResponse)(nil),          // 6: warehouse.GetWarehouseByIdResponse
}
var file_warehouse_proto_depIdxs = []int32{
	0, // 0: warehouse.CreateWarehouseRequest.warehouse:type_name -> warehouse.Warehouse
	0, // 1: warehouse.CreateWarehouseResponse.warehouse:type_name -> warehouse.Warehouse
	0, // 2: warehouse.GetAllWarehousesByStoreIdResponse.warehouse:type_name -> warehouse.Warehouse
	0, // 3: warehouse.GetWarehouseByIdResponse.warehouse:type_name -> warehouse.Warehouse
	1, // 4: warehouse.WarehouseClient.CreateWarehouse:input_type -> warehouse.CreateWarehouseRequest
	3, // 5: warehouse.WarehouseClient.GetAllWarehousesByStoreId:input_type -> warehouse.GetAllWarehousesByStoreIdRequest
	5, // 6: warehouse.WarehouseClient.GetWarehouseById:input_type -> warehouse.GetWarehouseByIdRequest
	2, // 7: warehouse.WarehouseClient.CreateWarehouse:output_type -> warehouse.CreateWarehouseResponse
	4, // 8: warehouse.WarehouseClient.GetAllWarehousesByStoreId:output_type -> warehouse.GetAllWarehousesByStoreIdResponse
	6, // 9: warehouse.WarehouseClient.GetWarehouseById:output_type -> warehouse.GetWarehouseByIdResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_warehouse_proto_init() }
func file_warehouse_proto_init() {
	if File_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWarehouseRequest); i {
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
		file_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWarehouseResponse); i {
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
		file_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllWarehousesByStoreIdRequest); i {
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
		file_warehouse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllWarehousesByStoreIdResponse); i {
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
		file_warehouse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehouseByIdRequest); i {
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
		file_warehouse_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehouseByIdResponse); i {
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
			RawDescriptor: file_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_proto_goTypes,
		DependencyIndexes: file_warehouse_proto_depIdxs,
		MessageInfos:      file_warehouse_proto_msgTypes,
	}.Build()
	File_warehouse_proto = out.File
	file_warehouse_proto_rawDesc = nil
	file_warehouse_proto_goTypes = nil
	file_warehouse_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WarehouseClientClient is the client API for WarehouseClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarehouseClientClient interface {
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error)
	GetAllWarehousesByStoreId(ctx context.Context, in *GetAllWarehousesByStoreIdRequest, opts ...grpc.CallOption) (*GetAllWarehousesByStoreIdResponse, error)
	GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*GetWarehouseByIdResponse, error)
}

type warehouseClientClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseClientClient(cc grpc.ClientConnInterface) WarehouseClientClient {
	return &warehouseClientClient{cc}
}

func (c *warehouseClientClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error) {
	out := new(CreateWarehouseResponse)
	err := c.cc.Invoke(ctx, "/warehouse.WarehouseClient/CreateWarehouse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClientClient) GetAllWarehousesByStoreId(ctx context.Context, in *GetAllWarehousesByStoreIdRequest, opts ...grpc.CallOption) (*GetAllWarehousesByStoreIdResponse, error) {
	out := new(GetAllWarehousesByStoreIdResponse)
	err := c.cc.Invoke(ctx, "/warehouse.WarehouseClient/GetAllWarehousesByStoreId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClientClient) GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*GetWarehouseByIdResponse, error) {
	out := new(GetWarehouseByIdResponse)
	err := c.cc.Invoke(ctx, "/warehouse.WarehouseClient/GetWarehouseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseClientServer is the server API for WarehouseClient service.
type WarehouseClientServer interface {
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	GetAllWarehousesByStoreId(context.Context, *GetAllWarehousesByStoreIdRequest) (*GetAllWarehousesByStoreIdResponse, error)
	GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseByIdResponse, error)
}

// UnimplementedWarehouseClientServer can be embedded to have forward compatible implementations.
type UnimplementedWarehouseClientServer struct {
}

func (*UnimplementedWarehouseClientServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (*UnimplementedWarehouseClientServer) GetAllWarehousesByStoreId(context.Context, *GetAllWarehousesByStoreIdRequest) (*GetAllWarehousesByStoreIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWarehousesByStoreId not implemented")
}
func (*UnimplementedWarehouseClientServer) GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouseById not implemented")
}

func RegisterWarehouseClientServer(s *grpc.Server, srv WarehouseClientServer) {
	s.RegisterService(&_WarehouseClient_serviceDesc, srv)
}

func _WarehouseClient_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseClientServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warehouse.WarehouseClient/CreateWarehouse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseClientServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseClient_GetAllWarehousesByStoreId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllWarehousesByStoreIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseClientServer).GetAllWarehousesByStoreId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warehouse.WarehouseClient/GetAllWarehousesByStoreId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseClientServer).GetAllWarehousesByStoreId(ctx, req.(*GetAllWarehousesByStoreIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseClient_GetWarehouseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarehouseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseClientServer).GetWarehouseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warehouse.WarehouseClient/GetWarehouseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseClientServer).GetWarehouseById(ctx, req.(*GetWarehouseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WarehouseClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.WarehouseClient",
	HandlerType: (*WarehouseClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWarehouse",
			Handler:    _WarehouseClient_CreateWarehouse_Handler,
		},
		{
			MethodName: "GetAllWarehousesByStoreId",
			Handler:    _WarehouseClient_GetAllWarehousesByStoreId_Handler,
		},
		{
			MethodName: "GetWarehouseById",
			Handler:    _WarehouseClient_GetWarehouseById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}
