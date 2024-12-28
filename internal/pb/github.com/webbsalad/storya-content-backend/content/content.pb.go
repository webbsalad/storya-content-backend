// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: api/content/content.proto

package content

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContentType int32

const (
	ContentType_MOVIE ContentType = 0
	ContentType_GAME  ContentType = 1
	ContentType_BOOK  ContentType = 2
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "MOVIE",
		1: "GAME",
		2: "BOOK",
	}
	ContentType_value = map[string]int32{
		"MOVIE": 0,
		"GAME":  1,
		"BOOK":  2,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_content_content_proto_enumTypes[0].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_api_content_content_proto_enumTypes[0]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_api_content_content_proto_rawDescGZIP(), []int{0}
}

type Value int32

const (
	Value_Like    Value = 0
	Value_Neutral Value = 1
	Value_Dislike Value = 3
)

// Enum value maps for Value.
var (
	Value_name = map[int32]string{
		0: "Like",
		1: "Neutral",
		3: "Dislike",
	}
	Value_value = map[string]int32{
		"Like":    0,
		"Neutral": 1,
		"Dislike": 3,
	}
)

func (x Value) Enum() *Value {
	p := new(Value)
	*p = x
	return p
}

func (x Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Value) Descriptor() protoreflect.EnumDescriptor {
	return file_api_content_content_proto_enumTypes[1].Descriptor()
}

func (Value) Type() protoreflect.EnumType {
	return &file_api_content_content_proto_enumTypes[1]
}

func (x Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Value.Descriptor instead.
func (Value) EnumDescriptor() ([]byte, []int) {
	return file_api_content_content_proto_rawDescGZIP(), []int{1}
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_api_content_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_api_content_content_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year      int32                  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Type      ContentType            `protobuf:"varint,4,opt,name=type,proto3,enum=content.ContentType" json:"type,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Tags      []*Tag                 `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_api_content_content_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_content_content_proto_msgTypes[1]
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
	return file_api_content_content_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Item) GetType() ContentType {
	if x != nil {
		return x.Type
	}
	return ContentType_MOVIE
}

func (x *Item) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Item) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_api_content_content_proto protoreflect.FileDescriptor

var file_api_content_content_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x2a, 0x2c, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x4d, 0x4f, 0x56, 0x49, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x41, 0x4d, 0x45, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x10, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x62, 0x73, 0x61, 0x6c, 0x61, 0x64,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_content_content_proto_rawDescOnce sync.Once
	file_api_content_content_proto_rawDescData = file_api_content_content_proto_rawDesc
)

func file_api_content_content_proto_rawDescGZIP() []byte {
	file_api_content_content_proto_rawDescOnce.Do(func() {
		file_api_content_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_content_content_proto_rawDescData)
	})
	return file_api_content_content_proto_rawDescData
}

var file_api_content_content_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_content_content_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_content_content_proto_goTypes = []any{
	(ContentType)(0),              // 0: content.ContentType
	(Value)(0),                    // 1: content.Value
	(*Tag)(nil),                   // 2: content.Tag
	(*Item)(nil),                  // 3: content.Item
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_api_content_content_proto_depIdxs = []int32{
	0, // 0: content.Item.type:type_name -> content.ContentType
	4, // 1: content.Item.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: content.Item.tags:type_name -> content.Tag
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_content_content_proto_init() }
func file_api_content_content_proto_init() {
	if File_api_content_content_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_content_content_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_content_content_proto_goTypes,
		DependencyIndexes: file_api_content_content_proto_depIdxs,
		EnumInfos:         file_api_content_content_proto_enumTypes,
		MessageInfos:      file_api_content_content_proto_msgTypes,
	}.Build()
	File_api_content_content_proto = out.File
	file_api_content_content_proto_rawDesc = nil
	file_api_content_content_proto_goTypes = nil
	file_api_content_content_proto_depIdxs = nil
}
