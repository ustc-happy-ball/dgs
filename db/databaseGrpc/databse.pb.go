// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: databse.proto

package databaseGrpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccountFindByPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *AccountFindByPhoneRequest) Reset() {
	*x = AccountFindByPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountFindByPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountFindByPhoneRequest) ProtoMessage() {}

func (x *AccountFindByPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountFindByPhoneRequest.ProtoReflect.Descriptor instead.
func (*AccountFindByPhoneRequest) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{0}
}

func (x *AccountFindByPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type AccountFindByPhoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccountFindByPhoneResponse) Reset() {
	*x = AccountFindByPhoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountFindByPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountFindByPhoneResponse) ProtoMessage() {}

func (x *AccountFindByPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountFindByPhoneResponse.ProtoReflect.Descriptor instead.
func (*AccountFindByPhoneResponse) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{1}
}

func (x *AccountFindByPhoneResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type AccountAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccountAddRequest) Reset() {
	*x = AccountAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAddRequest) ProtoMessage() {}

func (x *AccountAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAddRequest.ProtoReflect.Descriptor instead.
func (*AccountAddRequest) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{2}
}

func (x *AccountAddRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type AccountAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountAddResponse) Reset() {
	*x = AccountAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAddResponse) ProtoMessage() {}

func (x *AccountAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAddResponse.ProtoReflect.Descriptor instead.
func (*AccountAddResponse) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{3}
}

type PlayerFindByPlayerIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int32 `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *PlayerFindByPlayerIdRequest) Reset() {
	*x = PlayerFindByPlayerIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerFindByPlayerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerFindByPlayerIdRequest) ProtoMessage() {}

func (x *PlayerFindByPlayerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerFindByPlayerIdRequest.ProtoReflect.Descriptor instead.
func (*PlayerFindByPlayerIdRequest) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerFindByPlayerIdRequest) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

type PlayerFindByPlayerIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PlayerFindByPlayerIdResponse) Reset() {
	*x = PlayerFindByPlayerIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerFindByPlayerIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerFindByPlayerIdResponse) ProtoMessage() {}

func (x *PlayerFindByPlayerIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerFindByPlayerIdResponse.ProtoReflect.Descriptor instead.
func (*PlayerFindByPlayerIdResponse) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerFindByPlayerIdResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type PlayerAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PlayerAddRequest) Reset() {
	*x = PlayerAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerAddRequest) ProtoMessage() {}

func (x *PlayerAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerAddRequest.ProtoReflect.Descriptor instead.
func (*PlayerAddRequest) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{6}
}

func (x *PlayerAddRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type PlayerAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayerAddResponse) Reset() {
	*x = PlayerAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerAddResponse) ProtoMessage() {}

func (x *PlayerAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerAddResponse.ProtoReflect.Descriptor instead.
func (*PlayerAddResponse) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{7}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId      string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"` //objectId
	PlayerId      int32  `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	LoginPassword string `protobuf:"bytes,3,opt,name=loginPassword,proto3" json:"loginPassword,omitempty"`
	Delete        bool   `protobuf:"varint,4,opt,name=delete,proto3" json:"delete,omitempty"`
	Phone         string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	RecentLogin   int64  `protobuf:"varint,6,opt,name=recentLogin,proto3" json:"recentLogin,omitempty"`
	CreateAt      int64  `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt      int64  `protobuf:"varint,8,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{8}
}

func (x *Account) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *Account) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *Account) GetLoginPassword() string {
	if x != nil {
		return x.LoginPassword
	}
	return ""
}

func (x *Account) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *Account) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Account) GetRecentLogin() int64 {
	if x != nil {
		return x.RecentLogin
	}
	return 0
}

func (x *Account) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Account) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId     int32  `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	AccountId    string `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	HighestScore int32  `protobuf:"varint,4,opt,name=highestScore,proto3" json:"highestScore,omitempty"`
	HighestRank  int32  `protobuf:"varint,5,opt,name=highestRank,proto3" json:"highestRank,omitempty"`
	CreateAt     int64  `protobuf:"varint,6,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt     int64  `protobuf:"varint,7,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databse_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_databse_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_databse_proto_rawDescGZIP(), []int{9}
}

func (x *Player) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *Player) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Player) GetHighestScore() int32 {
	if x != nil {
		return x.HighestScore
	}
	return 0
}

func (x *Player) GetHighestRank() int32 {
	if x != nil {
		return x.HighestRank
	}
	return 0
}

func (x *Player) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Player) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

var File_databse_proto protoreflect.FileDescriptor

var file_databse_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x22, 0x31, 0x0a,
	0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x22, 0x4d, 0x0a, 0x1a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x44, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x1b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xef, 0x01, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0xc0, 0x01,
	0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x68, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x32, 0xce, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0xd0, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_databse_proto_rawDescOnce sync.Once
	file_databse_proto_rawDescData = file_databse_proto_rawDesc
)

func file_databse_proto_rawDescGZIP() []byte {
	file_databse_proto_rawDescOnce.Do(func() {
		file_databse_proto_rawDescData = protoimpl.X.CompressGZIP(file_databse_proto_rawDescData)
	})
	return file_databse_proto_rawDescData
}

var file_databse_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_databse_proto_goTypes = []interface{}{
	(*AccountFindByPhoneRequest)(nil),    // 0: databaseGrpc.AccountFindByPhoneRequest
	(*AccountFindByPhoneResponse)(nil),   // 1: databaseGrpc.AccountFindByPhoneResponse
	(*AccountAddRequest)(nil),            // 2: databaseGrpc.AccountAddRequest
	(*AccountAddResponse)(nil),           // 3: databaseGrpc.AccountAddResponse
	(*PlayerFindByPlayerIdRequest)(nil),  // 4: databaseGrpc.PlayerFindByPlayerIdRequest
	(*PlayerFindByPlayerIdResponse)(nil), // 5: databaseGrpc.PlayerFindByPlayerIdResponse
	(*PlayerAddRequest)(nil),             // 6: databaseGrpc.PlayerAddRequest
	(*PlayerAddResponse)(nil),            // 7: databaseGrpc.PlayerAddResponse
	(*Account)(nil),                      // 8: databaseGrpc.Account
	(*Player)(nil),                       // 9: databaseGrpc.Player
}
var file_databse_proto_depIdxs = []int32{
	8, // 0: databaseGrpc.AccountFindByPhoneResponse.account:type_name -> databaseGrpc.Account
	8, // 1: databaseGrpc.AccountAddRequest.account:type_name -> databaseGrpc.Account
	9, // 2: databaseGrpc.PlayerFindByPlayerIdResponse.player:type_name -> databaseGrpc.Player
	9, // 3: databaseGrpc.PlayerAddRequest.player:type_name -> databaseGrpc.Player
	0, // 4: databaseGrpc.AccountService.AccountFindByPhone:input_type -> databaseGrpc.AccountFindByPhoneRequest
	2, // 5: databaseGrpc.AccountService.AccountAdd:input_type -> databaseGrpc.AccountAddRequest
	4, // 6: databaseGrpc.PlayerService.PlayerFindByPlayerId:input_type -> databaseGrpc.PlayerFindByPlayerIdRequest
	6, // 7: databaseGrpc.PlayerService.PlayerAdd:input_type -> databaseGrpc.PlayerAddRequest
	1, // 8: databaseGrpc.AccountService.AccountFindByPhone:output_type -> databaseGrpc.AccountFindByPhoneResponse
	3, // 9: databaseGrpc.AccountService.AccountAdd:output_type -> databaseGrpc.AccountAddResponse
	5, // 10: databaseGrpc.PlayerService.PlayerFindByPlayerId:output_type -> databaseGrpc.PlayerFindByPlayerIdResponse
	7, // 11: databaseGrpc.PlayerService.PlayerAdd:output_type -> databaseGrpc.PlayerAddResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_databse_proto_init() }
func file_databse_proto_init() {
	if File_databse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_databse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountFindByPhoneRequest); i {
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
		file_databse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountFindByPhoneResponse); i {
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
		file_databse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAddRequest); i {
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
		file_databse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAddResponse); i {
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
		file_databse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerFindByPlayerIdRequest); i {
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
		file_databse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerFindByPlayerIdResponse); i {
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
		file_databse_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerAddRequest); i {
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
		file_databse_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerAddResponse); i {
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
		file_databse_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_databse_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_databse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_databse_proto_goTypes,
		DependencyIndexes: file_databse_proto_depIdxs,
		MessageInfos:      file_databse_proto_msgTypes,
	}.Build()
	File_databse_proto = out.File
	file_databse_proto_rawDesc = nil
	file_databse_proto_goTypes = nil
	file_databse_proto_depIdxs = nil
}
