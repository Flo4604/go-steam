// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: steammessages_offline.steamclient.proto

package unified

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

type COffline_GetOfflineLogonTicket_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority          *uint32 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	PerformEncryption *bool   `protobuf:"varint,2,opt,name=perform_encryption,json=performEncryption" json:"perform_encryption,omitempty"`
}

func (x *COffline_GetOfflineLogonTicket_Request) Reset() {
	*x = COffline_GetOfflineLogonTicket_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_offline_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COffline_GetOfflineLogonTicket_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetOfflineLogonTicket_Request) ProtoMessage() {}

func (x *COffline_GetOfflineLogonTicket_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerializedTicket []byte                               `protobuf:"bytes,1,opt,name=serialized_ticket,json=serializedTicket" json:"serialized_ticket,omitempty"`
	Signature        []byte                               `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	EncryptedTicket  *Offline_Ticket `protobuf:"bytes,3,opt,name=encrypted_ticket,json=encryptedTicket" json:"encrypted_ticket,omitempty"`
}

func (x *COffline_GetOfflineLogonTicket_Response) Reset() {
	*x = COffline_GetOfflineLogonTicket_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_offline_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COffline_GetOfflineLogonTicket_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetOfflineLogonTicket_Response) ProtoMessage() {}

func (x *COffline_GetOfflineLogonTicket_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) Reset() {
	*x = COffline_GetUnsignedOfflineLogonTicket_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_offline_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetUnsignedOfflineLogonTicket_Request) ProtoMessage() {}

func (x *COffline_GetUnsignedOfflineLogonTicket_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accountid           *uint32 `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Rtime32CreationTime *uint32 `protobuf:"fixed32,2,opt,name=rtime32_creation_time,json=rtime32CreationTime" json:"rtime32_creation_time,omitempty"`
}

func (x *COffline_OfflineLogonTicket) Reset() {
	*x = COffline_OfflineLogonTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_offline_steamclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COffline_OfflineLogonTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_OfflineLogonTicket) ProtoMessage() {}

func (x *COffline_OfflineLogonTicket) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *COffline_OfflineLogonTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) Reset() {
	*x = COffline_GetUnsignedOfflineLogonTicket_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_offline_steamclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COffline_GetUnsignedOfflineLogonTicket_Response) ProtoMessage() {}

func (x *COffline_GetUnsignedOfflineLogonTicket_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_offline_steamclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_steammessages_offline_steamclient_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x26, 0x43, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f,
	0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a,
	0x12, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01, 0x0a,
	0x27, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x30, 0x0a, 0x2e, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f,
	0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x6f, 0x0a, 0x1b, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x32,
	0x0a, 0x15, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x33, 0x32, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x13, 0x72,
	0x74, 0x69, 0x6d, 0x65, 0x33, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x67, 0x0a, 0x2f, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47,
	0x65, 0x74, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x32, 0xa3, 0x03, 0x0a, 0x07,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x27, 0x2e, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x43, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x82, 0xb5, 0x18, 0x45, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x6c, 0x6f,
	0x67, 0x6f, 0x6e, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12,
	0xc1, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x2f, 0x2e, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74,
	0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4c,
	0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x43, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65,
	0x74, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82, 0xb5, 0x18, 0x39, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6e,
	0x20, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x20, 0x6c, 0x6f, 0x67, 0x6f, 0x6e, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x1c, 0x82, 0xb5, 0x18, 0x18, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x20, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x03, 0x80, 0x01, 0x01,
}

var (
	file_steammessages_offline_steamclient_proto_rawDescOnce sync.Once
	file_steammessages_offline_steamclient_proto_rawDescData = file_steammessages_offline_steamclient_proto_rawDesc
)

func file_steammessages_offline_steamclient_proto_rawDescGZIP() []byte {
	file_steammessages_offline_steamclient_proto_rawDescOnce.Do(func() {
		file_steammessages_offline_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_offline_steamclient_proto_rawDescData)
	})
	return file_steammessages_offline_steamclient_proto_rawDescData
}

var file_steammessages_offline_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_steammessages_offline_steamclient_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_steammessages_offline_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COffline_GetOfflineLogonTicket_Request); i {
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
		file_steammessages_offline_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COffline_GetOfflineLogonTicket_Response); i {
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
		file_steammessages_offline_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COffline_GetUnsignedOfflineLogonTicket_Request); i {
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
		file_steammessages_offline_steamclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COffline_OfflineLogonTicket); i {
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
		file_steammessages_offline_steamclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COffline_GetUnsignedOfflineLogonTicket_Response); i {
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
			RawDescriptor: file_steammessages_offline_steamclient_proto_rawDesc,
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
	file_steammessages_offline_steamclient_proto_rawDesc = nil
	file_steammessages_offline_steamclient_proto_goTypes = nil
	file_steammessages_offline_steamclient_proto_depIdxs = nil
}
