// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: backend/proto/Page.proto

package pb

import (
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

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string           `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Links       []string         `protobuf:"bytes,2,rep,name=Links,proto3" json:"Links,omitempty"`
	Title       string           `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Meta        *PageMeta        `protobuf:"bytes,5,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Visited     bool             `protobuf:"varint,6,opt,name=Visited,proto3" json:"Visited,omitempty"`
	Timestamp   string           `protobuf:"bytes,7,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Words       map[string]int32 `protobuf:"bytes,8,rep,name=Words,proto3" json:"Words,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_Page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_Page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_backend_proto_Page_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Page) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Page) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Page) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Page) GetMeta() *PageMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Page) GetVisited() bool {
	if x != nil {
		return x.Visited
	}
	return false
}

func (x *Page) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Page) GetWords() map[string]int32 {
	if x != nil {
		return x.Words
	}
	return nil
}

type PageMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OG       map[string]string `protobuf:"bytes,1,rep,name=OG,proto3" json:"OG,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Keywords []string          `protobuf:"bytes,2,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
	Manifest string            `protobuf:"bytes,3,opt,name=Manifest,proto3" json:"Manifest,omitempty"`
	Ld       string            `protobuf:"bytes,4,opt,name=Ld,proto3" json:"Ld,omitempty"`
}

func (x *PageMeta) Reset() {
	*x = PageMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_Page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageMeta) ProtoMessage() {}

func (x *PageMeta) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_Page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageMeta.ProtoReflect.Descriptor instead.
func (*PageMeta) Descriptor() ([]byte, []int) {
	return file_backend_proto_Page_proto_rawDescGZIP(), []int{1}
}

func (x *PageMeta) GetOG() map[string]string {
	if x != nil {
		return x.OG
	}
	return nil
}

func (x *PageMeta) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *PageMeta) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *PageMeta) GetLd() string {
	if x != nil {
		return x.Ld
	}
	return ""
}

type PageIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=Keys,proto3" json:"Keys,omitempty"`
}

func (x *PageIndex) Reset() {
	*x = PageIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_Page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageIndex) ProtoMessage() {}

func (x *PageIndex) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_Page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageIndex.ProtoReflect.Descriptor instead.
func (*PageIndex) Descriptor() ([]byte, []int) {
	return file_backend_proto_Page_proto_rawDescGZIP(), []int{2}
}

func (x *PageIndex) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_backend_proto_Page_proto protoreflect.FileDescriptor

var file_backend_proto_Page_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x50, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa5,
	0x02, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x56, 0x69, 0x73, 0x69,
	0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x29, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x38, 0x0a, 0x0a,
	0x57, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaf, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x02, 0x4f, 0x47, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4f, 0x47,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x4f, 0x47, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x4c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4c,
	0x64, 0x1a, 0x35, 0x0a, 0x07, 0x4f, 0x47, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_proto_Page_proto_rawDescOnce sync.Once
	file_backend_proto_Page_proto_rawDescData = file_backend_proto_Page_proto_rawDesc
)

func file_backend_proto_Page_proto_rawDescGZIP() []byte {
	file_backend_proto_Page_proto_rawDescOnce.Do(func() {
		file_backend_proto_Page_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_proto_Page_proto_rawDescData)
	})
	return file_backend_proto_Page_proto_rawDescData
}

var file_backend_proto_Page_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_backend_proto_Page_proto_goTypes = []interface{}{
	(*Page)(nil),      // 0: pb.Page
	(*PageMeta)(nil),  // 1: pb.PageMeta
	(*PageIndex)(nil), // 2: pb.PageIndex
	nil,               // 3: pb.Page.WordsEntry
	nil,               // 4: pb.PageMeta.OGEntry
}
var file_backend_proto_Page_proto_depIdxs = []int32{
	1, // 0: pb.Page.Meta:type_name -> pb.PageMeta
	3, // 1: pb.Page.Words:type_name -> pb.Page.WordsEntry
	4, // 2: pb.PageMeta.OG:type_name -> pb.PageMeta.OGEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_backend_proto_Page_proto_init() }
func file_backend_proto_Page_proto_init() {
	if File_backend_proto_Page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_proto_Page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_backend_proto_Page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageMeta); i {
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
		file_backend_proto_Page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageIndex); i {
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
			RawDescriptor: file_backend_proto_Page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_proto_Page_proto_goTypes,
		DependencyIndexes: file_backend_proto_Page_proto_depIdxs,
		MessageInfos:      file_backend_proto_Page_proto_msgTypes,
	}.Build()
	File_backend_proto_Page_proto = out.File
	file_backend_proto_Page_proto_rawDesc = nil
	file_backend_proto_Page_proto_goTypes = nil
	file_backend_proto_Page_proto_depIdxs = nil
}
