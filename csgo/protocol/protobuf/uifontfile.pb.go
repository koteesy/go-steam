// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: uifontfile_format.proto

package protobuf

import (
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

type CUIFontFilePB struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	FontFileName     *string                `protobuf:"bytes,1,opt,name=font_file_name,json=fontFileName" json:"font_file_name,omitempty"`
	OpentypeFontData []byte                 `protobuf:"bytes,2,opt,name=opentype_font_data,json=opentypeFontData" json:"opentype_font_data,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CUIFontFilePB) Reset() {
	*x = CUIFontFilePB{}
	mi := &file_uifontfile_format_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CUIFontFilePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePB) ProtoMessage() {}

func (x *CUIFontFilePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{0}
}

func (x *CUIFontFilePB) GetFontFileName() string {
	if x != nil && x.FontFileName != nil {
		return *x.FontFileName
	}
	return ""
}

func (x *CUIFontFilePB) GetOpentypeFontData() []byte {
	if x != nil {
		return x.OpentypeFontData
	}
	return nil
}

type CUIFontFilePackagePB struct {
	state              protoimpl.MessageState                         `protogen:"open.v1"`
	PackageVersion     *uint32                                        `protobuf:"varint,1,req,name=package_version,json=packageVersion" json:"package_version,omitempty"`
	EncryptedFontFiles []*CUIFontFilePackagePB_CUIEncryptedFontFilePB `protobuf:"bytes,2,rep,name=encrypted_font_files,json=encryptedFontFiles" json:"encrypted_font_files,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CUIFontFilePackagePB) Reset() {
	*x = CUIFontFilePackagePB{}
	mi := &file_uifontfile_format_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CUIFontFilePackagePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePackagePB) ProtoMessage() {}

func (x *CUIFontFilePackagePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePackagePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePackagePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{1}
}

func (x *CUIFontFilePackagePB) GetPackageVersion() uint32 {
	if x != nil && x.PackageVersion != nil {
		return *x.PackageVersion
	}
	return 0
}

func (x *CUIFontFilePackagePB) GetEncryptedFontFiles() []*CUIFontFilePackagePB_CUIEncryptedFontFilePB {
	if x != nil {
		return x.EncryptedFontFiles
	}
	return nil
}

type CUIFontFilePackagePB_CUIEncryptedFontFilePB struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	EncryptedContents []byte                 `protobuf:"bytes,1,opt,name=encrypted_contents,json=encryptedContents" json:"encrypted_contents,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) Reset() {
	*x = CUIFontFilePackagePB_CUIEncryptedFontFilePB{}
	mi := &file_uifontfile_format_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) ProtoMessage() {}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePackagePB_CUIEncryptedFontFilePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) GetEncryptedContents() []byte {
	if x != nil {
		return x.EncryptedContents
	}
	return nil
}

var File_uifontfile_format_proto protoreflect.FileDescriptor

const file_uifontfile_format_proto_rawDesc = "" +
	"\n" +
	"\x17uifontfile_format.proto\"c\n" +
	"\rCUIFontFilePB\x12$\n" +
	"\x0efont_file_name\x18\x01 \x01(\tR\ffontFileName\x12,\n" +
	"\x12opentype_font_data\x18\x02 \x01(\fR\x10opentypeFontData\"\xe8\x01\n" +
	"\x14CUIFontFilePackagePB\x12'\n" +
	"\x0fpackage_version\x18\x01 \x02(\rR\x0epackageVersion\x12^\n" +
	"\x14encrypted_font_files\x18\x02 \x03(\v2,.CUIFontFilePackagePB.CUIEncryptedFontFilePBR\x12encryptedFontFiles\x1aG\n" +
	"\x16CUIEncryptedFontFilePB\x12-\n" +
	"\x12encrypted_contents\x18\x01 \x01(\fR\x11encryptedContents"

var (
	file_uifontfile_format_proto_rawDescOnce sync.Once
	file_uifontfile_format_proto_rawDescData []byte
)

func file_uifontfile_format_proto_rawDescGZIP() []byte {
	file_uifontfile_format_proto_rawDescOnce.Do(func() {
		file_uifontfile_format_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_uifontfile_format_proto_rawDesc), len(file_uifontfile_format_proto_rawDesc)))
	})
	return file_uifontfile_format_proto_rawDescData
}

var file_uifontfile_format_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_uifontfile_format_proto_goTypes = []any{
	(*CUIFontFilePB)(nil),                               // 0: CUIFontFilePB
	(*CUIFontFilePackagePB)(nil),                        // 1: CUIFontFilePackagePB
	(*CUIFontFilePackagePB_CUIEncryptedFontFilePB)(nil), // 2: CUIFontFilePackagePB.CUIEncryptedFontFilePB
}
var file_uifontfile_format_proto_depIdxs = []int32{
	2, // 0: CUIFontFilePackagePB.encrypted_font_files:type_name -> CUIFontFilePackagePB.CUIEncryptedFontFilePB
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uifontfile_format_proto_init() }
func file_uifontfile_format_proto_init() {
	if File_uifontfile_format_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_uifontfile_format_proto_rawDesc), len(file_uifontfile_format_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uifontfile_format_proto_goTypes,
		DependencyIndexes: file_uifontfile_format_proto_depIdxs,
		MessageInfos:      file_uifontfile_format_proto_msgTypes,
	}.Build()
	File_uifontfile_format_proto = out.File
	file_uifontfile_format_proto_goTypes = nil
	file_uifontfile_format_proto_depIdxs = nil
}
