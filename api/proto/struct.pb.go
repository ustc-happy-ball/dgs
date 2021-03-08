// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: struct.proto

package proto

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

type HeroMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId        int32         `protobuf:"varint,1,opt,name=heroId,proto3" json:"heroId,omitempty"`                                //玩家的ID
	HeroPosition  *CoordinateXY `protobuf:"bytes,2,opt,name=heroPosition,proto3" json:"heroPosition,omitempty"`                     //玩家的坐标
	HeroSpeed     float32       `protobuf:"fixed32,3,opt,name=heroSpeed,proto3" json:"heroSpeed,omitempty"`                         //玩家的速度
	HeroSize      float32       `protobuf:"fixed32,4,opt,name=heroSize,proto3" json:"heroSize,omitempty"`                           //玩家的大小
	HeroDirection *CoordinateXY `protobuf:"bytes,5,opt,name=heroDirection,proto3" json:"heroDirection,omitempty"`                   //玩家的行进方向
	HeroStatus    HERO_STATUS   `protobuf:"varint,6,opt,name=heroStatus,proto3,enum=proto.HERO_STATUS" json:"heroStatus,omitempty"` //玩家的状态
}

func (x *HeroMsg) Reset() {
	*x = HeroMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroMsg) ProtoMessage() {}

func (x *HeroMsg) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroMsg.ProtoReflect.Descriptor instead.
func (*HeroMsg) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{0}
}

func (x *HeroMsg) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *HeroMsg) GetHeroPosition() *CoordinateXY {
	if x != nil {
		return x.HeroPosition
	}
	return nil
}

func (x *HeroMsg) GetHeroSpeed() float32 {
	if x != nil {
		return x.HeroSpeed
	}
	return 0
}

func (x *HeroMsg) GetHeroSize() float32 {
	if x != nil {
		return x.HeroSize
	}
	return 0
}

func (x *HeroMsg) GetHeroDirection() *CoordinateXY {
	if x != nil {
		return x.HeroDirection
	}
	return nil
}

func (x *HeroMsg) GetHeroStatus() HERO_STATUS {
	if x != nil {
		return x.HeroStatus
	}
	return HERO_STATUS_LIVE
}

type ItemMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       int32         `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`                                //Item的Id
	ItemType     ENTITY_TYPE   `protobuf:"varint,2,opt,name=ItemType,proto3,enum=proto.ENTITY_TYPE" json:"ItemType,omitempty"`     //Item的类型
	ItemPosition *CoordinateXY `protobuf:"bytes,3,opt,name=ItemPosition,proto3" json:"ItemPosition,omitempty"`                     //Item的初始坐标
	ItemStatus   ITEM_STATUS   `protobuf:"varint,4,opt,name=ItemStatus,proto3,enum=proto.ITEM_STATUS" json:"ItemStatus,omitempty"` //Item的状态
	ItemRadius   float32       `protobuf:"fixed32,5,opt,name=ItemRadius,proto3" json:"ItemRadius,omitempty"`
}

func (x *ItemMsg) Reset() {
	*x = ItemMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemMsg) ProtoMessage() {}

func (x *ItemMsg) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemMsg.ProtoReflect.Descriptor instead.
func (*ItemMsg) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{1}
}

func (x *ItemMsg) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemMsg) GetItemType() ENTITY_TYPE {
	if x != nil {
		return x.ItemType
	}
	return ENTITY_TYPE_HERO_TYPE
}

func (x *ItemMsg) GetItemPosition() *CoordinateXY {
	if x != nil {
		return x.ItemPosition
	}
	return nil
}

func (x *ItemMsg) GetItemStatus() ITEM_STATUS {
	if x != nil {
		return x.ItemStatus
	}
	return ITEM_STATUS_ITEM_LIVE
}

func (x *ItemMsg) GetItemRadius() float32 {
	if x != nil {
		return x.ItemRadius
	}
	return 0
}

type MapMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XMin float32 `protobuf:"fixed32,1,opt,name=XMin,proto3" json:"XMin,omitempty"`
	XMax float32 `protobuf:"fixed32,2,opt,name=XMax,proto3" json:"XMax,omitempty"`
	YMin float32 `protobuf:"fixed32,3,opt,name=YMin,proto3" json:"YMin,omitempty"`
	YMax float32 `protobuf:"fixed32,4,opt,name=YMax,proto3" json:"YMax,omitempty"` //四个边界
}

func (x *MapMsg) Reset() {
	*x = MapMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMsg) ProtoMessage() {}

func (x *MapMsg) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMsg.ProtoReflect.Descriptor instead.
func (*MapMsg) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{2}
}

func (x *MapMsg) GetXMin() float32 {
	if x != nil {
		return x.XMin
	}
	return 0
}

func (x *MapMsg) GetXMax() float32 {
	if x != nil {
		return x.XMax
	}
	return 0
}

func (x *MapMsg) GetYMin() float32 {
	if x != nil {
		return x.YMin
	}
	return 0
}

func (x *MapMsg) GetYMax() float32 {
	if x != nil {
		return x.YMax
	}
	return 0
}

type CoordinateXY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoordinateX float32 `protobuf:"fixed32,1,opt,name=coordinateX,proto3" json:"coordinateX,omitempty"` //横坐标
	CoordinateY float32 `protobuf:"fixed32,2,opt,name=coordinateY,proto3" json:"coordinateY,omitempty"` //纵坐标
}

func (x *CoordinateXY) Reset() {
	*x = CoordinateXY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinateXY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinateXY) ProtoMessage() {}

func (x *CoordinateXY) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinateXY.ProtoReflect.Descriptor instead.
func (*CoordinateXY) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{3}
}

func (x *CoordinateXY) GetCoordinateX() float32 {
	if x != nil {
		return x.CoordinateX
	}
	return 0
}

func (x *CoordinateXY) GetCoordinateY() float32 {
	if x != nil {
		return x.CoordinateY
	}
	return 0
}

type ConnectMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectMsg) Reset() {
	*x = ConnectMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectMsg) ProtoMessage() {}

func (x *ConnectMsg) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectMsg.ProtoReflect.Descriptor instead.
func (*ConnectMsg) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectMsg) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ConnectMsg) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_struct_proto protoreflect.FileDescriptor

var file_struct_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x73, 0x67, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x48, 0x65, 0x72, 0x6f, 0x4d, 0x73, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x68, 0x65, 0x72, 0x6f,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x58, 0x59, 0x52, 0x0c, 0x68, 0x65, 0x72, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x72, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x68, 0x65, 0x72, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x65, 0x72, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x68, 0x65, 0x72, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x68,
	0x65, 0x72, 0x6f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x58, 0x59, 0x52, 0x0d, 0x68, 0x65, 0x72, 0x6f, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0a, 0x68, 0x65, 0x72, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x45, 0x52, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x0a,
	0x68, 0x65, 0x72, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x07, 0x49,
	0x74, 0x65, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37,
	0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x58, 0x59, 0x52, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52,
	0x0a, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x06, 0x4d,
	0x61, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x58, 0x4d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x58, 0x4d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x58, 0x4d, 0x61,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x58, 0x4d, 0x61, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x59, 0x4d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x59, 0x4d, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x4d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x59, 0x4d, 0x61, 0x78, 0x22, 0x52, 0x0a, 0x0c, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x58, 0x59, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x58, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x59, 0x22, 0x30, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_struct_proto_rawDescOnce sync.Once
	file_struct_proto_rawDescData = file_struct_proto_rawDesc
)

func file_struct_proto_rawDescGZIP() []byte {
	file_struct_proto_rawDescOnce.Do(func() {
		file_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_struct_proto_rawDescData)
	})
	return file_struct_proto_rawDescData
}

var file_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_struct_proto_goTypes = []interface{}{
	(*HeroMsg)(nil),      // 0: proto.HeroMsg
	(*ItemMsg)(nil),      // 1: proto.ItemMsg
	(*MapMsg)(nil),       // 2: proto.MapMsg
	(*CoordinateXY)(nil), // 3: proto.CoordinateXY
	(*ConnectMsg)(nil),   // 4: proto.ConnectMsg
	(HERO_STATUS)(0),     // 5: proto.HERO_STATUS
	(ENTITY_TYPE)(0),     // 6: proto.ENTITY_TYPE
	(ITEM_STATUS)(0),     // 7: proto.ITEM_STATUS
}
var file_struct_proto_depIdxs = []int32{
	3, // 0: proto.HeroMsg.heroPosition:type_name -> proto.CoordinateXY
	3, // 1: proto.HeroMsg.heroDirection:type_name -> proto.CoordinateXY
	5, // 2: proto.HeroMsg.heroStatus:type_name -> proto.HERO_STATUS
	6, // 3: proto.ItemMsg.ItemType:type_name -> proto.ENTITY_TYPE
	3, // 4: proto.ItemMsg.ItemPosition:type_name -> proto.CoordinateXY
	7, // 5: proto.ItemMsg.ItemStatus:type_name -> proto.ITEM_STATUS
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_struct_proto_init() }
func file_struct_proto_init() {
	if File_struct_proto != nil {
		return
	}
	file_msgenum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroMsg); i {
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
		file_struct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemMsg); i {
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
		file_struct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMsg); i {
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
		file_struct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinateXY); i {
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
		file_struct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectMsg); i {
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
			RawDescriptor: file_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_struct_proto_goTypes,
		DependencyIndexes: file_struct_proto_depIdxs,
		MessageInfos:      file_struct_proto_msgTypes,
	}.Build()
	File_struct_proto = out.File
	file_struct_proto_rawDesc = nil
	file_struct_proto_goTypes = nil
	file_struct_proto_depIdxs = nil
}
