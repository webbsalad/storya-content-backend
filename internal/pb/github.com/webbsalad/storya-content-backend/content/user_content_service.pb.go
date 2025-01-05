// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: api/content/user_content_service.proto

package content

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetUserItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContentType ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=content.ContentType" json:"content_type,omitempty"`
}

func (x *GetUserItemsRequest) Reset() {
	*x = GetUserItemsRequest{}
	mi := &file_api_content_user_content_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserItemsRequest) ProtoMessage() {}

func (x *GetUserItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserItemsRequest.ProtoReflect.Descriptor instead.
func (*GetUserItemsRequest) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserItemsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserItemsRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_MOVIE
}

type GetUserItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserItems []*UserItem `protobuf:"bytes,1,rep,name=userItems,proto3" json:"userItems,omitempty"`
}

func (x *GetUserItemsResponse) Reset() {
	*x = GetUserItemsResponse{}
	mi := &file_api_content_user_content_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserItemsResponse) ProtoMessage() {}

func (x *GetUserItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserItemsResponse.ProtoReflect.Descriptor instead.
func (*GetUserItemsResponse) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserItemsResponse) GetUserItems() []*UserItem {
	if x != nil {
		return x.UserItems
	}
	return nil
}

type GetValuedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContentType ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=content.ContentType" json:"content_type,omitempty"`
	Value       Value       `protobuf:"varint,3,opt,name=value,proto3,enum=content.Value" json:"value,omitempty"`
}

func (x *GetValuedRequest) Reset() {
	*x = GetValuedRequest{}
	mi := &file_api_content_user_content_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetValuedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValuedRequest) ProtoMessage() {}

func (x *GetValuedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValuedRequest.ProtoReflect.Descriptor instead.
func (*GetValuedRequest) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetValuedRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetValuedRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_MOVIE
}

func (x *GetValuedRequest) GetValue() Value {
	if x != nil {
		return x.Value
	}
	return Value_Like
}

type GetValuedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetValuedResponse) Reset() {
	*x = GetValuedResponse{}
	mi := &file_api_content_user_content_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetValuedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValuedResponse) ProtoMessage() {}

func (x *GetValuedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValuedResponse.ProtoReflect.Descriptor instead.
func (*GetValuedResponse) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetValuedResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId      string      `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ContentType ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=content.ContentType" json:"content_type,omitempty"`
	Value       Value       `protobuf:"varint,3,opt,name=value,proto3,enum=content.Value" json:"value,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_api_content_user_content_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *AddRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_MOVIE
}

func (x *AddRequest) GetValue() Value {
	if x != nil {
		return x.Value
	}
	return Value_Like
}

type RemoveItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId      string      `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ContentType ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=content.ContentType" json:"content_type,omitempty"`
}

func (x *RemoveItemRequest) Reset() {
	*x = RemoveItemRequest{}
	mi := &file_api_content_user_content_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemRequest) ProtoMessage() {}

func (x *RemoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_user_content_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemRequest) Descriptor() ([]byte, []int) {
	return file_api_content_user_content_service_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *RemoveItemRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_MOVIE
}

var File_api_content_user_content_service_proto protoreflect.FileDescriptor

var file_api_content_user_content_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x94, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8e, 0x01,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6f,
	0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x32,
	0xc4, 0x03, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x7d, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x75,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x7d, 0x12, 0x56, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x3a, 0x01, 0x2a, 0x32, 0x17, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x7b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x12, 0x67, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f, 0x7b, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x62, 0x73, 0x61, 0x6c, 0x61, 0x64, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_content_user_content_service_proto_rawDescOnce sync.Once
	file_api_content_user_content_service_proto_rawDescData = file_api_content_user_content_service_proto_rawDesc
)

func file_api_content_user_content_service_proto_rawDescGZIP() []byte {
	file_api_content_user_content_service_proto_rawDescOnce.Do(func() {
		file_api_content_user_content_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_content_user_content_service_proto_rawDescData)
	})
	return file_api_content_user_content_service_proto_rawDescData
}

var file_api_content_user_content_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_content_user_content_service_proto_goTypes = []any{
	(*GetUserItemsRequest)(nil),  // 0: content.GetUserItemsRequest
	(*GetUserItemsResponse)(nil), // 1: content.GetUserItemsResponse
	(*GetValuedRequest)(nil),     // 2: content.GetValuedRequest
	(*GetValuedResponse)(nil),    // 3: content.GetValuedResponse
	(*AddRequest)(nil),           // 4: content.AddRequest
	(*RemoveItemRequest)(nil),    // 5: content.RemoveItemRequest
	(ContentType)(0),             // 6: content.ContentType
	(*UserItem)(nil),             // 7: content.UserItem
	(Value)(0),                   // 8: content.Value
	(*Item)(nil),                 // 9: content.Item
	(*emptypb.Empty)(nil),        // 10: google.protobuf.Empty
}
var file_api_content_user_content_service_proto_depIdxs = []int32{
	6,  // 0: content.GetUserItemsRequest.content_type:type_name -> content.ContentType
	7,  // 1: content.GetUserItemsResponse.userItems:type_name -> content.UserItem
	6,  // 2: content.GetValuedRequest.content_type:type_name -> content.ContentType
	8,  // 3: content.GetValuedRequest.value:type_name -> content.Value
	9,  // 4: content.GetValuedResponse.items:type_name -> content.Item
	6,  // 5: content.AddRequest.content_type:type_name -> content.ContentType
	8,  // 6: content.AddRequest.value:type_name -> content.Value
	6,  // 7: content.RemoveItemRequest.content_type:type_name -> content.ContentType
	0,  // 8: content.UserContentService.GetUserItems:input_type -> content.GetUserItemsRequest
	2,  // 9: content.UserContentService.GetValued:input_type -> content.GetValuedRequest
	4,  // 10: content.UserContentService.Add:input_type -> content.AddRequest
	5,  // 11: content.UserContentService.Remove:input_type -> content.RemoveItemRequest
	1,  // 12: content.UserContentService.GetUserItems:output_type -> content.GetUserItemsResponse
	3,  // 13: content.UserContentService.GetValued:output_type -> content.GetValuedResponse
	10, // 14: content.UserContentService.Add:output_type -> google.protobuf.Empty
	10, // 15: content.UserContentService.Remove:output_type -> google.protobuf.Empty
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_content_user_content_service_proto_init() }
func file_api_content_user_content_service_proto_init() {
	if File_api_content_user_content_service_proto != nil {
		return
	}
	file_api_content_content_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_content_user_content_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_content_user_content_service_proto_goTypes,
		DependencyIndexes: file_api_content_user_content_service_proto_depIdxs,
		MessageInfos:      file_api_content_user_content_service_proto_msgTypes,
	}.Build()
	File_api_content_user_content_service_proto = out.File
	file_api_content_user_content_service_proto_rawDesc = nil
	file_api_content_user_content_service_proto_goTypes = nil
	file_api_content_user_content_service_proto_depIdxs = nil
}
