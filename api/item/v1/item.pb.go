// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/item/v1/item.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	mi := &file_api_item_v1_item_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type GetItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	mi := &file_api_item_v1_item_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemResponse) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *GetItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetItemResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateItemRequest) Reset() {
	*x = CreateItemRequest{}
	mi := &file_api_item_v1_item_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemRequest) ProtoMessage() {}

func (x *CreateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemRequest.ProtoReflect.Descriptor instead.
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{2}
}

func (x *CreateItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateItemRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateItemResponse) Reset() {
	*x = CreateItemResponse{}
	mi := &file_api_item_v1_item_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemResponse) ProtoMessage() {}

func (x *CreateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemResponse.ProtoReflect.Descriptor instead.
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{3}
}

func (x *CreateItemResponse) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	mi := &file_api_item_v1_item_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *UpdateItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateItemRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	mi := &file_api_item_v1_item_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	mi := &file_api_item_v1_item_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type DeleteItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	mi := &file_api_item_v1_item_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	mi := &file_api_item_v1_item_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{8}
}

func (x *ListItemsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListItemsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListItemsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Item                `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	mi := &file_api_item_v1_item_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{9}
}

func (x *ListItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemId        string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_api_item_v1_item_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_item_v1_item_proto_msgTypes[10]
	if x != nil {
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
	return file_api_item_v1_item_proto_rawDescGZIP(), []int{10}
}

func (x *Item) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_item_v1_item_proto protoreflect.FileDescriptor

var file_api_item_v1_item_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xf4, 0x03, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0xf4,
	0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x37,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02,
	0x28, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04, 0x18, 0x64,
	0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x55, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x96,
	0x03, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_item_v1_item_proto_rawDescOnce sync.Once
	file_api_item_v1_item_proto_rawDescData []byte
)

func file_api_item_v1_item_proto_rawDescGZIP() []byte {
	file_api_item_v1_item_proto_rawDescOnce.Do(func() {
		file_api_item_v1_item_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_item_v1_item_proto_rawDesc), len(file_api_item_v1_item_proto_rawDesc)))
	})
	return file_api_item_v1_item_proto_rawDescData
}

var file_api_item_v1_item_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_item_v1_item_proto_goTypes = []any{
	(*GetItemRequest)(nil),     // 0: api.item.v1.GetItemRequest
	(*GetItemResponse)(nil),    // 1: api.item.v1.GetItemResponse
	(*CreateItemRequest)(nil),  // 2: api.item.v1.CreateItemRequest
	(*CreateItemResponse)(nil), // 3: api.item.v1.CreateItemResponse
	(*UpdateItemRequest)(nil),  // 4: api.item.v1.UpdateItemRequest
	(*UpdateItemResponse)(nil), // 5: api.item.v1.UpdateItemResponse
	(*DeleteItemRequest)(nil),  // 6: api.item.v1.DeleteItemRequest
	(*DeleteItemResponse)(nil), // 7: api.item.v1.DeleteItemResponse
	(*ListItemsRequest)(nil),   // 8: api.item.v1.ListItemsRequest
	(*ListItemsResponse)(nil),  // 9: api.item.v1.ListItemsResponse
	(*Item)(nil),               // 10: api.item.v1.Item
}
var file_api_item_v1_item_proto_depIdxs = []int32{
	10, // 0: api.item.v1.ListItemsResponse.items:type_name -> api.item.v1.Item
	0,  // 1: api.item.v1.ItemService.GetItem:input_type -> api.item.v1.GetItemRequest
	2,  // 2: api.item.v1.ItemService.CreateItem:input_type -> api.item.v1.CreateItemRequest
	4,  // 3: api.item.v1.ItemService.UpdateItem:input_type -> api.item.v1.UpdateItemRequest
	6,  // 4: api.item.v1.ItemService.DeleteItem:input_type -> api.item.v1.DeleteItemRequest
	8,  // 5: api.item.v1.ItemService.ListItems:input_type -> api.item.v1.ListItemsRequest
	1,  // 6: api.item.v1.ItemService.GetItem:output_type -> api.item.v1.GetItemResponse
	3,  // 7: api.item.v1.ItemService.CreateItem:output_type -> api.item.v1.CreateItemResponse
	5,  // 8: api.item.v1.ItemService.UpdateItem:output_type -> api.item.v1.UpdateItemResponse
	7,  // 9: api.item.v1.ItemService.DeleteItem:output_type -> api.item.v1.DeleteItemResponse
	9,  // 10: api.item.v1.ItemService.ListItems:output_type -> api.item.v1.ListItemsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_item_v1_item_proto_init() }
func file_api_item_v1_item_proto_init() {
	if File_api_item_v1_item_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_item_v1_item_proto_rawDesc), len(file_api_item_v1_item_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_item_v1_item_proto_goTypes,
		DependencyIndexes: file_api_item_v1_item_proto_depIdxs,
		MessageInfos:      file_api_item_v1_item_proto_msgTypes,
	}.Build()
	File_api_item_v1_item_proto = out.File
	file_api_item_v1_item_proto_goTypes = nil
	file_api_item_v1_item_proto_depIdxs = nil
}
