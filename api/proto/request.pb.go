// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: request.proto

package __

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityChangeRequest *EntityInfoChangeRequest `protobuf:"bytes,1,opt,name=entityChangeRequest,proto3" json:"entityChangeRequest,omitempty"` //实体改变请求，例如两条蛇相撞的时候、蛇吃道具的时候要发此类请求
	HeroQuitRequest     *HeroQuitRequest         `protobuf:"bytes,2,opt,name=heroQuitRequest,proto3" json:"heroQuitRequest,omitempty"`         //玩家退出请求
	EnterGameRequest    *EnterGameRequest        `protobuf:"bytes,3,opt,name=enterGameRequest,proto3" json:"enterGameRequest,omitempty"`
	HeartBeatRequest    *HeartBeatRequest        `protobuf:"bytes,4,opt,name=heartBeatRequest,proto3" json:"heartBeatRequest,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetEntityChangeRequest() *EntityInfoChangeRequest {
	if x != nil {
		return x.EntityChangeRequest
	}
	return nil
}

func (x *Request) GetHeroQuitRequest() *HeroQuitRequest {
	if x != nil {
		return x.HeroQuitRequest
	}
	return nil
}

func (x *Request) GetEnterGameRequest() *EnterGameRequest {
	if x != nil {
		return x.EnterGameRequest
	}
	return nil
}

func (x *Request) GetHeartBeatRequest() *HeartBeatRequest {
	if x != nil {
		return x.HeartBeatRequest
	}
	return nil
}

type EnterGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId         int32       `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`                //玩家ID
	ClientConnectMsg *ConnectMsg `protobuf:"bytes,2,opt,name=clientConnectMsg,proto3" json:"clientConnectMsg,omitempty"` //网络信息
	PlayerName       string      `protobuf:"bytes,3,opt,name=playerName,proto3" json:"playerName,omitempty"`             //玩家名称
}

func (x *EnterGameRequest) Reset() {
	*x = EnterGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterGameRequest) ProtoMessage() {}

func (x *EnterGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterGameRequest.ProtoReflect.Descriptor instead.
func (*EnterGameRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *EnterGameRequest) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *EnterGameRequest) GetClientConnectMsg() *ConnectMsg {
	if x != nil {
		return x.ClientConnectMsg
	}
	return nil
}

func (x *EnterGameRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type EntityInfoChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType  EVENT_TYPE  `protobuf:"varint,1,opt,name=eventType,proto3,enum=proto.EVENT_TYPE" json:"eventType,omitempty"`    // 事件类型 enum
	HeroId     int32       `protobuf:"varint,2,opt,name=heroId,proto3" json:"heroId,omitempty"`                                // 发送请求的实体Id
	LinkedId   int32       `protobuf:"varint,3,opt,name=linkedId,proto3" json:"linkedId,omitempty"`                            // 交互对象
	LinkedType ENTITY_TYPE `protobuf:"varint,4,opt,name=linkedType,proto3,enum=proto.ENTITY_TYPE" json:"linkedType,omitempty"` //交互的对象类型
	HeroMsg    *HeroMsg    `protobuf:"bytes,5,opt,name=heroMsg,proto3" json:"heroMsg,omitempty"`                               //玩家的信息
}

func (x *EntityInfoChangeRequest) Reset() {
	*x = EntityInfoChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityInfoChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityInfoChangeRequest) ProtoMessage() {}

func (x *EntityInfoChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityInfoChangeRequest.ProtoReflect.Descriptor instead.
func (*EntityInfoChangeRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *EntityInfoChangeRequest) GetEventType() EVENT_TYPE {
	if x != nil {
		return x.EventType
	}
	return EVENT_TYPE_HERO_COLLISION
}

func (x *EntityInfoChangeRequest) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *EntityInfoChangeRequest) GetLinkedId() int32 {
	if x != nil {
		return x.LinkedId
	}
	return 0
}

func (x *EntityInfoChangeRequest) GetLinkedType() ENTITY_TYPE {
	if x != nil {
		return x.LinkedType
	}
	return ENTITY_TYPE_HERO_TYPE
}

func (x *EntityInfoChangeRequest) GetHeroMsg() *HeroMsg {
	if x != nil {
		return x.HeroMsg
	}
	return nil
}

type HeroQuitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int32 `protobuf:"varint,1,opt,name=heroId,proto3" json:"heroId,omitempty"` //玩家Id
	Time   int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`     //离开时间
}

func (x *HeroQuitRequest) Reset() {
	*x = HeroQuitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroQuitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroQuitRequest) ProtoMessage() {}

func (x *HeroQuitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroQuitRequest.ProtoReflect.Descriptor instead.
func (*HeroQuitRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *HeroQuitRequest) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *HeroQuitRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendTime int64 `protobuf:"varint,1,opt,name=sendTime,proto3" json:"sendTime,omitempty"` //消息发送时间，用于测试网络时延
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{4}
}

func (x *HeartBeatRequest) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x73, 0x67, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x50, 0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x13, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x0f, 0x68, 0x65, 0x72, 0x6f, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0f, 0x68, 0x65, 0x72, 0x6f, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x10, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x10, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8d, 0x01,
	0x0a, 0x10, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d,
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x10, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdc, 0x01,
	0x0a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x72, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x4e, 0x54, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x68, 0x65, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x4d, 0x73, 0x67, 0x52, 0x07, 0x68, 0x65, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x22, 0x3d, 0x0a, 0x0f,
	0x48, 0x65, 0x72, 0x6f, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_request_proto_goTypes = []interface{}{
	(*Request)(nil),                 // 0: proto.Request
	(*EnterGameRequest)(nil),        // 1: proto.EnterGameRequest
	(*EntityInfoChangeRequest)(nil), // 2: proto.EntityInfoChangeRequest
	(*HeroQuitRequest)(nil),         // 3: proto.HeroQuitRequest
	(*HeartBeatRequest)(nil),        // 4: proto.HeartBeatRequest
	(*ConnectMsg)(nil),              // 5: proto.ConnectMsg
	(EVENT_TYPE)(0),                 // 6: proto.EVENT_TYPE
	(ENTITY_TYPE)(0),                // 7: proto.ENTITY_TYPE
	(*HeroMsg)(nil),                 // 8: proto.HeroMsg
}
var file_request_proto_depIdxs = []int32{
	2, // 0: proto.Request.entityChangeRequest:type_name -> proto.EntityInfoChangeRequest
	3, // 1: proto.Request.heroQuitRequest:type_name -> proto.HeroQuitRequest
	1, // 2: proto.Request.enterGameRequest:type_name -> proto.EnterGameRequest
	4, // 3: proto.Request.heartBeatRequest:type_name -> proto.HeartBeatRequest
	5, // 4: proto.EnterGameRequest.clientConnectMsg:type_name -> proto.ConnectMsg
	6, // 5: proto.EntityInfoChangeRequest.eventType:type_name -> proto.EVENT_TYPE
	7, // 6: proto.EntityInfoChangeRequest.linkedType:type_name -> proto.ENTITY_TYPE
	8, // 7: proto.EntityInfoChangeRequest.heroMsg:type_name -> proto.HeroMsg
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	file_msgenum_proto_init()
	file_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterGameRequest); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityInfoChangeRequest); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroQuitRequest); i {
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
		file_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatRequest); i {
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
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
