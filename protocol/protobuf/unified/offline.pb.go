// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: steammessages_offline.steamclient.proto

package unified

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

type COffline_GetOfflineLogonTicket_Request struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Priority          *uint32                `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	PerformEncryption *bool                  `protobuf:"varint,2,opt,name=perform_encryption,json=performEncryption" json:"perform_encryption,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *COffline_GetOfflineLogonTicket_Request) Reset() {
	*x = COffline_GetOfflineLogonTicket_Request{}
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *COffline_GetOfflineLogonTicket_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetOfflineLogonTicket_Request) ProtoMessage() {}

func (x *COffline_GetOfflineLogonTicket_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COffline_GetOfflineLogonTicket_Request.ProtoReflect.Descriptor instead.
func (*COffline_GetOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_offline_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *COffline_GetOfflineLogonTicket_Request) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *COffline_GetOfflineLogonTicket_Request) GetPerformEncryption() bool {
	if x != nil && x.PerformEncryption != nil {
		return *x.PerformEncryption
	}
	return false
}

type COffline_GetOfflineLogonTicket_Response struct {
	state            protoimpl.MessageState               `protogen:"open.v1"`
	SerializedTicket []byte                               `protobuf:"bytes,1,opt,name=serialized_ticket,json=serializedTicket" json:"serialized_ticket,omitempty"`
	Signature        []byte                               `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	EncryptedTicket  *Offline_Ticket `protobuf:"bytes,3,opt,name=encrypted_ticket,json=encryptedTicket" json:"encrypted_ticket,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *COffline_GetOfflineLogonTicket_Response) Reset() {
	*x = COffline_GetOfflineLogonTicket_Response{}
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *COffline_GetOfflineLogonTicket_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetOfflineLogonTicket_Response) ProtoMessage() {}

func (x *COffline_GetOfflineLogonTicket_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COffline_GetOfflineLogonTicket_Response.ProtoReflect.Descriptor instead.
func (*COffline_GetOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_offline_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *COffline_GetOfflineLogonTicket_Response) GetSerializedTicket() []byte {
	if x != nil {
		return x.SerializedTicket
	}
	return nil
}

func (x *COffline_GetOfflineLogonTicket_Response) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *COffline_GetOfflineLogonTicket_Response) GetEncryptedTicket() *Offline_Ticket {
	if x != nil {
		return x.EncryptedTicket
	}
	return nil
}

type COffline_GetUnsignedOfflineLogonTicket_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) Reset() {
	*x = COffline_GetUnsignedOfflineLogonTicket_Request{}
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetUnsignedOfflineLogonTicket_Request) ProtoMessage() {}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COffline_GetUnsignedOfflineLogonTicket_Request.ProtoReflect.Descriptor instead.
func (*COffline_GetUnsignedOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_offline_steamclient_proto_rawDescGZIP(), []int{2}
}

type COffline_OfflineLogonTicket struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Accountid           *uint32                `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Rtime32CreationTime *uint32                `protobuf:"fixed32,2,opt,name=rtime32_creation_time,json=rtime32CreationTime" json:"rtime32_creation_time,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *COffline_OfflineLogonTicket) Reset() {
	*x = COffline_OfflineLogonTicket{}
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *COffline_OfflineLogonTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_OfflineLogonTicket) ProtoMessage() {}

func (x *COffline_OfflineLogonTicket) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COffline_OfflineLogonTicket.ProtoReflect.Descriptor instead.
func (*COffline_OfflineLogonTicket) Descriptor() ([]byte, []int) {
	return file_steammessages_offline_steamclient_proto_rawDescGZIP(), []int{3}
}

func (x *COffline_OfflineLogonTicket) GetAccountid() uint32 {
	if x != nil && x.Accountid != nil {
		return *x.Accountid
	}
	return 0
}

func (x *COffline_OfflineLogonTicket) GetRtime32CreationTime() uint32 {
	if x != nil && x.Rtime32CreationTime != nil {
		return *x.Rtime32CreationTime
	}
	return 0
}

type COffline_GetUnsignedOfflineLogonTicket_Response struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Ticket        *COffline_OfflineLogonTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) Reset() {
	*x = COffline_GetUnsignedOfflineLogonTicket_Response{}
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetUnsignedOfflineLogonTicket_Response) ProtoMessage() {}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COffline_GetUnsignedOfflineLogonTicket_Response.ProtoReflect.Descriptor instead.
func (*COffline_GetUnsignedOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_offline_steamclient_proto_rawDescGZIP(), []int{4}
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) GetTicket() *COffline_OfflineLogonTicket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

var File_steammessages_offline_steamclient_proto protoreflect.FileDescriptor

const file_steammessages_offline_steamclient_proto_rawDesc = "" +
	"\n" +
	"'steammessages_offline.steamclient.proto\x1a\x18steammessages_base.proto\x1a,steammessages_unified_base.steamclient.proto\x1a\x14offline_ticket.proto\"s\n" +
	"&COffline_GetOfflineLogonTicket_Request\x12\x1a\n" +
	"\bpriority\x18\x01 \x01(\rR\bpriority\x12-\n" +
	"\x12perform_encryption\x18\x02 \x01(\bR\x11performEncryption\"\xb0\x01\n" +
	"'COffline_GetOfflineLogonTicket_Response\x12+\n" +
	"\x11serialized_ticket\x18\x01 \x01(\fR\x10serializedTicket\x12\x1c\n" +
	"\tsignature\x18\x02 \x01(\fR\tsignature\x12:\n" +
	"\x10encrypted_ticket\x18\x03 \x01(\v2\x0f.Offline_TicketR\x0fencryptedTicket\"0\n" +
	".COffline_GetUnsignedOfflineLogonTicket_Request\"o\n" +
	"\x1bCOffline_OfflineLogonTicket\x12\x1c\n" +
	"\taccountid\x18\x01 \x01(\rR\taccountid\x122\n" +
	"\x15rtime32_creation_time\x18\x02 \x01(\aR\x13rtime32CreationTime\"g\n" +
	"/COffline_GetUnsignedOfflineLogonTicket_Response\x124\n" +
	"\x06ticket\x18\x01 \x01(\v2\x1c.COffline_OfflineLogonTicketR\x06ticket2\xfa\x01\n" +
	"\aOffline\x12j\n" +
	"\x15GetOfflineLogonTicket\x12'.COffline_GetOfflineLogonTicket_Request\x1a(.COffline_GetOfflineLogonTicket_Response\x12\x82\x01\n" +
	"\x1dGetUnsignedOfflineLogonTicket\x12/.COffline_GetUnsignedOfflineLogonTicket_Request\x1a0.COffline_GetUnsignedOfflineLogonTicket_ResponseB5Z0github.com/koteesy/go-steam/v3/protocol/protobuf\x80\x01\x01"

var (
	file_steammessages_offline_steamclient_proto_rawDescOnce sync.Once
	file_steammessages_offline_steamclient_proto_rawDescData []byte
)

func file_steammessages_offline_steamclient_proto_rawDescGZIP() []byte {
	file_steammessages_offline_steamclient_proto_rawDescOnce.Do(func() {
		file_steammessages_offline_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_steammessages_offline_steamclient_proto_rawDesc), len(file_steammessages_offline_steamclient_proto_rawDesc)))
	})
	return file_steammessages_offline_steamclient_proto_rawDescData
}

var file_steammessages_offline_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_steammessages_offline_steamclient_proto_goTypes = []any{
	(*COffline_GetOfflineLogonTicket_Request)(nil),          // 0: COffline_GetOfflineLogonTicket_Request
	(*COffline_GetOfflineLogonTicket_Response)(nil),         // 1: COffline_GetOfflineLogonTicket_Response
	(*COffline_GetUnsignedOfflineLogonTicket_Request)(nil),  // 2: COffline_GetUnsignedOfflineLogonTicket_Request
	(*COffline_OfflineLogonTicket)(nil),                     // 3: COffline_OfflineLogonTicket
	(*COffline_GetUnsignedOfflineLogonTicket_Response)(nil), // 4: COffline_GetUnsignedOfflineLogonTicket_Response
	(*Offline_Ticket)(nil),             // 5: Offline_Ticket
}
var file_steammessages_offline_steamclient_proto_depIdxs = []int32{
	5, // 0: COffline_GetOfflineLogonTicket_Response.encrypted_ticket:type_name -> Offline_Ticket
	3, // 1: COffline_GetUnsignedOfflineLogonTicket_Response.ticket:type_name -> COffline_OfflineLogonTicket
	0, // 2: Offline.GetOfflineLogonTicket:input_type -> COffline_GetOfflineLogonTicket_Request
	2, // 3: Offline.GetUnsignedOfflineLogonTicket:input_type -> COffline_GetUnsignedOfflineLogonTicket_Request
	1, // 4: Offline.GetOfflineLogonTicket:output_type -> COffline_GetOfflineLogonTicket_Response
	4, // 5: Offline.GetUnsignedOfflineLogonTicket:output_type -> COffline_GetUnsignedOfflineLogonTicket_Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steammessages_offline_steamclient_proto_init() }
func file_steammessages_offline_steamclient_proto_init() {
	if File_steammessages_offline_steamclient_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_steammessages_offline_steamclient_proto_rawDesc), len(file_steammessages_offline_steamclient_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_offline_steamclient_proto_goTypes,
		DependencyIndexes: file_steammessages_offline_steamclient_proto_depIdxs,
		MessageInfos:      file_steammessages_offline_steamclient_proto_msgTypes,
	}.Build()
	File_steammessages_offline_steamclient_proto = out.File
	file_steammessages_offline_steamclient_proto_goTypes = nil
	file_steammessages_offline_steamclient_proto_depIdxs = nil
}
