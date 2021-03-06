// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: pb/inventory.proto

package inventory_pb

import (
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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID      string  `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	ItemName    string  `protobuf:"bytes,3,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string  `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Quantity    int64   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price       float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Item) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddInventoryItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *AddInventoryItemRequest) Reset() {
	*x = AddInventoryItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInventoryItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInventoryItemRequest) ProtoMessage() {}

func (x *AddInventoryItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInventoryItemRequest.ProtoReflect.Descriptor instead.
func (*AddInventoryItemRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *AddInventoryItemRequest) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetInventoryItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInventoryItemRequest) Reset() {
	*x = GetInventoryItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInventoryItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInventoryItemRequest) ProtoMessage() {}

func (x *GetInventoryItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInventoryItemRequest.ProtoReflect.Descriptor instead.
func (*GetInventoryItemRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetInventoryItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInventoryItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInventoryItemRequest) Reset() {
	*x = DeleteInventoryItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInventoryItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInventoryItemRequest) ProtoMessage() {}

func (x *DeleteInventoryItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInventoryItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteInventoryItemRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteInventoryItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateInventoryItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateInventoryItemRequest) Reset() {
	*x = UpdateInventoryItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInventoryItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryItemRequest) ProtoMessage() {}

func (x *UpdateInventoryItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateInventoryItemRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInventoryItemRequest) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListInventoryItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Lim    int64  `protobuf:"varint,2,opt,name=lim,proto3" json:"lim,omitempty"`
	Off    int64  `protobuf:"varint,3,opt,name=off,proto3" json:"off,omitempty"`
}

func (x *ListInventoryItemsRequest) Reset() {
	*x = ListInventoryItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInventoryItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInventoryItemsRequest) ProtoMessage() {}

func (x *ListInventoryItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInventoryItemsRequest.ProtoReflect.Descriptor instead.
func (*ListInventoryItemsRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *ListInventoryItemsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ListInventoryItemsRequest) GetLim() int64 {
	if x != nil {
		return x.Lim
	}
	return 0
}

func (x *ListInventoryItemsRequest) GetOff() int64 {
	if x != nil {
		return x.Off
	}
	return 0
}

type ListInventoryItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListInventoryItemsResponse) Reset() {
	*x = ListInventoryItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInventoryItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInventoryItemsResponse) ProtoMessage() {}

func (x *ListInventoryItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInventoryItemsResponse.ProtoReflect.Descriptor instead.
func (*ListInventoryItemsResponse) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *ListInventoryItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pb_inventory_proto protoreflect.FileDescriptor

var file_pb_inventory_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50,
	0x42, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc,
	0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x40, 0x0a,
	0x17, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x50, 0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22,
	0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x50, 0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x57, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6c, 0x69, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6f, 0x66, 0x66, 0x22, 0x45, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50,
	0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc8, 0x03,
	0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x50, 0x42, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x50, 0x42, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12,
	0x67, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x50, 0x42, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_inventory_proto_rawDescOnce sync.Once
	file_pb_inventory_proto_rawDescData = file_pb_inventory_proto_rawDesc
)

func file_pb_inventory_proto_rawDescGZIP() []byte {
	file_pb_inventory_proto_rawDescOnce.Do(func() {
		file_pb_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_inventory_proto_rawDescData)
	})
	return file_pb_inventory_proto_rawDescData
}

var file_pb_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_inventory_proto_goTypes = []interface{}{
	(*Item)(nil),                       // 0: inventoryPB.Item
	(*AddInventoryItemRequest)(nil),    // 1: inventoryPB.AddInventoryItemRequest
	(*GetInventoryItemRequest)(nil),    // 2: inventoryPB.GetInventoryItemRequest
	(*DeleteInventoryItemRequest)(nil), // 3: inventoryPB.DeleteInventoryItemRequest
	(*UpdateInventoryItemRequest)(nil), // 4: inventoryPB.UpdateInventoryItemRequest
	(*ListInventoryItemsRequest)(nil),  // 5: inventoryPB.ListInventoryItemsRequest
	(*ListInventoryItemsResponse)(nil), // 6: inventoryPB.ListInventoryItemsResponse
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_pb_inventory_proto_depIdxs = []int32{
	0, // 0: inventoryPB.AddInventoryItemRequest.item:type_name -> inventoryPB.Item
	0, // 1: inventoryPB.UpdateInventoryItemRequest.item:type_name -> inventoryPB.Item
	0, // 2: inventoryPB.ListInventoryItemsResponse.items:type_name -> inventoryPB.Item
	1, // 3: inventoryPB.InventoryService.AddInventoryItem:input_type -> inventoryPB.AddInventoryItemRequest
	2, // 4: inventoryPB.InventoryService.GetInventoryItem:input_type -> inventoryPB.GetInventoryItemRequest
	3, // 5: inventoryPB.InventoryService.DeleteInventoryItem:input_type -> inventoryPB.DeleteInventoryItemRequest
	4, // 6: inventoryPB.InventoryService.UpdateInventoryItem:input_type -> inventoryPB.UpdateInventoryItemRequest
	5, // 7: inventoryPB.InventoryService.ListInventoryItems:input_type -> inventoryPB.ListInventoryItemsRequest
	0, // 8: inventoryPB.InventoryService.AddInventoryItem:output_type -> inventoryPB.Item
	0, // 9: inventoryPB.InventoryService.GetInventoryItem:output_type -> inventoryPB.Item
	7, // 10: inventoryPB.InventoryService.DeleteInventoryItem:output_type -> google.protobuf.Empty
	0, // 11: inventoryPB.InventoryService.UpdateInventoryItem:output_type -> inventoryPB.Item
	6, // 12: inventoryPB.InventoryService.ListInventoryItems:output_type -> inventoryPB.ListInventoryItemsResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_inventory_proto_init() }
func file_pb_inventory_proto_init() {
	if File_pb_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_pb_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInventoryItemRequest); i {
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
		file_pb_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInventoryItemRequest); i {
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
		file_pb_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInventoryItemRequest); i {
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
		file_pb_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInventoryItemRequest); i {
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
		file_pb_inventory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInventoryItemsRequest); i {
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
		file_pb_inventory_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInventoryItemsResponse); i {
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
			RawDescriptor: file_pb_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_inventory_proto_goTypes,
		DependencyIndexes: file_pb_inventory_proto_depIdxs,
		MessageInfos:      file_pb_inventory_proto_msgTypes,
	}.Build()
	File_pb_inventory_proto = out.File
	file_pb_inventory_proto_rawDesc = nil
	file_pb_inventory_proto_goTypes = nil
	file_pb_inventory_proto_depIdxs = nil
}
