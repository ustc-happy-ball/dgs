// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: msgenum.proto

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

type MSG_TYPE int32

const (
	MSG_TYPE_NOTIFY   MSG_TYPE = 0 //通知类型
	MSG_TYPE_REQUEST  MSG_TYPE = 1 //请求类型
	MSG_TYPE_RESPONSE MSG_TYPE = 2 //答复类型
)

// Enum value maps for MSG_TYPE.
var (
	MSG_TYPE_name = map[int32]string{
		0: "NOTIFY",
		1: "REQUEST",
		2: "RESPONSE",
	}
	MSG_TYPE_value = map[string]int32{
		"NOTIFY":   0,
		"REQUEST":  1,
		"RESPONSE": 2,
	}
)

func (x MSG_TYPE) Enum() *MSG_TYPE {
	p := new(MSG_TYPE)
	*p = x
	return p
}

func (x MSG_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[0].Descriptor()
}

func (MSG_TYPE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[0]
}

func (x MSG_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_TYPE.Descriptor instead.
func (MSG_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{0}
}

type GAME_MSG_CODE int32

const (
	GAME_MSG_CODE_ENTITY_INFO_CHANGE_REQUEST  GAME_MSG_CODE = 0 //EntityInfoChangeRequest
	GAME_MSG_CODE_ENTITY_INFO_CHANGE_RESPONSE GAME_MSG_CODE = 1 //EntityInfoChangeResponse
	GAME_MSG_CODE_ENTITY_INFO_NOTIFY          GAME_MSG_CODE = 2 //EntityInfoChangeNotify
	GAME_MSG_CODE_HERO_QUIT_REQUEST           GAME_MSG_CODE = 3 //HeroQuitRequest
	GAME_MSG_CODE_GAME_GLOBAL_INFO_NOTIFY     GAME_MSG_CODE = 4 //GameGlobalInfoNotify
	GAME_MSG_CODE_TIME_NOTIFY                 GAME_MSG_CODE = 5 //TimeNotify
	GAME_MSG_CODE_ENTER_GAME_NOTIFY           GAME_MSG_CODE = 6 //EnterGameNotify
	GAME_MSG_CODE_ENTER_GAME_REQUEST          GAME_MSG_CODE = 7
	GAME_MSG_CODE_ENTER_GAME_RESPONSE         GAME_MSG_CODE = 8
	GAME_MSG_CODE_HERO_VIEW_NOTIFY            GAME_MSG_CODE = 9
)

// Enum value maps for GAME_MSG_CODE.
var (
	GAME_MSG_CODE_name = map[int32]string{
		0: "ENTITY_INFO_CHANGE_REQUEST",
		1: "ENTITY_INFO_CHANGE_RESPONSE",
		2: "ENTITY_INFO_NOTIFY",
		3: "HERO_QUIT_REQUEST",
		4: "GAME_GLOBAL_INFO_NOTIFY",
		5: "TIME_NOTIFY",
		6: "ENTER_GAME_NOTIFY",
		7: "ENTER_GAME_REQUEST",
		8: "ENTER_GAME_RESPONSE",
		9: "HERO_VIEW_NOTIFY",
	}
	GAME_MSG_CODE_value = map[string]int32{
		"ENTITY_INFO_CHANGE_REQUEST":  0,
		"ENTITY_INFO_CHANGE_RESPONSE": 1,
		"ENTITY_INFO_NOTIFY":          2,
		"HERO_QUIT_REQUEST":           3,
		"GAME_GLOBAL_INFO_NOTIFY":     4,
		"TIME_NOTIFY":                 5,
		"ENTER_GAME_NOTIFY":           6,
		"ENTER_GAME_REQUEST":          7,
		"ENTER_GAME_RESPONSE":         8,
		"HERO_VIEW_NOTIFY":            9,
	}
)

func (x GAME_MSG_CODE) Enum() *GAME_MSG_CODE {
	p := new(GAME_MSG_CODE)
	*p = x
	return p
}

func (x GAME_MSG_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GAME_MSG_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[1].Descriptor()
}

func (GAME_MSG_CODE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[1]
}

func (x GAME_MSG_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GAME_MSG_CODE.Descriptor instead.
func (GAME_MSG_CODE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{1}
}

type HERO_STATUS int32

const (
	HERO_STATUS_LIVE       HERO_STATUS = 0 //存活
	HERO_STATUS_DEAD       HERO_STATUS = 1 //死亡
	HERO_STATUS_INVINCIBLE HERO_STATUS = 2 //暂时无敌
)

// Enum value maps for HERO_STATUS.
var (
	HERO_STATUS_name = map[int32]string{
		0: "LIVE",
		1: "DEAD",
		2: "INVINCIBLE",
	}
	HERO_STATUS_value = map[string]int32{
		"LIVE":       0,
		"DEAD":       1,
		"INVINCIBLE": 2,
	}
)

func (x HERO_STATUS) Enum() *HERO_STATUS {
	p := new(HERO_STATUS)
	*p = x
	return p
}

func (x HERO_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HERO_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[2].Descriptor()
}

func (HERO_STATUS) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[2]
}

func (x HERO_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HERO_STATUS.Descriptor instead.
func (HERO_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{2}
}

type ITEM_STATUS int32

const (
	ITEM_STATUS_ITEM_LIVE ITEM_STATUS = 0 //存活
	ITEM_STATUS_ITEM_DEAD ITEM_STATUS = 1 //死亡
)

// Enum value maps for ITEM_STATUS.
var (
	ITEM_STATUS_name = map[int32]string{
		0: "ITEM_LIVE",
		1: "ITEM_DEAD",
	}
	ITEM_STATUS_value = map[string]int32{
		"ITEM_LIVE": 0,
		"ITEM_DEAD": 1,
	}
)

func (x ITEM_STATUS) Enum() *ITEM_STATUS {
	p := new(ITEM_STATUS)
	*p = x
	return p
}

func (x ITEM_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ITEM_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[3].Descriptor()
}

func (ITEM_STATUS) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[3]
}

func (x ITEM_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ITEM_STATUS.Descriptor instead.
func (ITEM_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{3}
}

// eventType的枚举
type EVENT_TYPE int32

const (
	EVENT_TYPE_HERO_COLLISION EVENT_TYPE = 0 //玩家碰撞事件
	EVENT_TYPE_ITEM_COLLISION EVENT_TYPE = 1 //道具碰撞事件
	EVENT_TYPE_HERO_MOVE      EVENT_TYPE = 2 //玩家移动事件
	EVENT_TYPE_HERO_GROW      EVENT_TYPE = 3 //玩家变大事件
)

// Enum value maps for EVENT_TYPE.
var (
	EVENT_TYPE_name = map[int32]string{
		0: "HERO_COLLISION",
		1: "ITEM_COLLISION",
		2: "HERO_MOVE",
		3: "HERO_GROW",
	}
	EVENT_TYPE_value = map[string]int32{
		"HERO_COLLISION": 0,
		"ITEM_COLLISION": 1,
		"HERO_MOVE":      2,
		"HERO_GROW":      3,
	}
)

func (x EVENT_TYPE) Enum() *EVENT_TYPE {
	p := new(EVENT_TYPE)
	*p = x
	return p
}

func (x EVENT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVENT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[4].Descriptor()
}

func (EVENT_TYPE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[4]
}

func (x EVENT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVENT_TYPE.Descriptor instead.
func (EVENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{4}
}

// entityType的枚举
type ENTITY_TYPE int32

const (
	ENTITY_TYPE_HERO_TYPE ENTITY_TYPE = 0 //Hero类型
	ENTITY_TYPE_PROP_TYPE ENTITY_TYPE = 1 //Prop类型--道具
	ENTITY_TYPE_FOOD_TYPE ENTITY_TYPE = 2 //Food类型--食物
)

// Enum value maps for ENTITY_TYPE.
var (
	ENTITY_TYPE_name = map[int32]string{
		0: "HERO_TYPE",
		1: "PROP_TYPE",
		2: "FOOD_TYPE",
	}
	ENTITY_TYPE_value = map[string]int32{
		"HERO_TYPE": 0,
		"PROP_TYPE": 1,
		"FOOD_TYPE": 2,
	}
)

func (x ENTITY_TYPE) Enum() *ENTITY_TYPE {
	p := new(ENTITY_TYPE)
	*p = x
	return p
}

func (x ENTITY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ENTITY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[5].Descriptor()
}

func (ENTITY_TYPE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[5]
}

func (x ENTITY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ENTITY_TYPE.Descriptor instead.
func (ENTITY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{5}
}

// response的result枚举
type RESULT_TYPE int32

const (
	RESULT_TYPE_SUCCESS RESULT_TYPE = 0 //成功
	RESULT_TYPE_FAIL    RESULT_TYPE = 1 //失败
)

// Enum value maps for RESULT_TYPE.
var (
	RESULT_TYPE_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	RESULT_TYPE_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x RESULT_TYPE) Enum() *RESULT_TYPE {
	p := new(RESULT_TYPE)
	*p = x
	return p
}

func (x RESULT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RESULT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[6].Descriptor()
}

func (RESULT_TYPE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[6]
}

func (x RESULT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RESULT_TYPE.Descriptor instead.
func (RESULT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{6}
}

// 视野通知的类型
type VIEW_TYPE int32

const (
	VIEW_TYPE_ENTER VIEW_TYPE = 0 // 进入视野范围
	VIEW_TYPE_LEAVE VIEW_TYPE = 1 // 离开视野范围
)

// Enum value maps for VIEW_TYPE.
var (
	VIEW_TYPE_name = map[int32]string{
		0: "ENTER",
		1: "LEAVE",
	}
	VIEW_TYPE_value = map[string]int32{
		"ENTER": 0,
		"LEAVE": 1,
	}
)

func (x VIEW_TYPE) Enum() *VIEW_TYPE {
	p := new(VIEW_TYPE)
	*p = x
	return p
}

func (x VIEW_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VIEW_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msgenum_proto_enumTypes[7].Descriptor()
}

func (VIEW_TYPE) Type() protoreflect.EnumType {
	return &file_msgenum_proto_enumTypes[7]
}

func (x VIEW_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VIEW_TYPE.Descriptor instead.
func (VIEW_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msgenum_proto_rawDescGZIP(), []int{7}
}

var File_msgenum_proto protoreflect.FileDescriptor

var file_msgenum_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x31, 0x0a, 0x08, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x2a, 0x8b, 0x02, 0x0a, 0x0d, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x1e, 0x0a, 0x1a, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x45, 0x52, 0x4f, 0x5f, 0x51, 0x55, 0x49,
	0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x47,
	0x41, 0x4d, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4e, 0x54,
	0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x06,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10,
	0x08, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x45, 0x52, 0x4f, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x09, 0x2a, 0x31, 0x0a, 0x0b, 0x48, 0x45, 0x52, 0x4f, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e,
	0x56, 0x49, 0x4e, 0x43, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x0b, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x01, 0x2a, 0x52, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x52, 0x4f, 0x5f, 0x43, 0x4f,
	0x4c, 0x4c, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x48, 0x45, 0x52, 0x4f, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x48, 0x45, 0x52, 0x4f, 0x5f, 0x47, 0x52, 0x4f, 0x57, 0x10, 0x03, 0x2a, 0x3a, 0x0a, 0x0b, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45,
	0x52, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f,
	0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x4f, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x2a, 0x24, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x2a, 0x21, 0x0a,
	0x09, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgenum_proto_rawDescOnce sync.Once
	file_msgenum_proto_rawDescData = file_msgenum_proto_rawDesc
)

func file_msgenum_proto_rawDescGZIP() []byte {
	file_msgenum_proto_rawDescOnce.Do(func() {
		file_msgenum_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgenum_proto_rawDescData)
	})
	return file_msgenum_proto_rawDescData
}

var file_msgenum_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_msgenum_proto_goTypes = []interface{}{
	(MSG_TYPE)(0),      // 0: proto.MSG_TYPE
	(GAME_MSG_CODE)(0), // 1: proto.GAME_MSG_CODE
	(HERO_STATUS)(0),   // 2: proto.HERO_STATUS
	(ITEM_STATUS)(0),   // 3: proto.ITEM_STATUS
	(EVENT_TYPE)(0),    // 4: proto.EVENT_TYPE
	(ENTITY_TYPE)(0),   // 5: proto.ENTITY_TYPE
	(RESULT_TYPE)(0),   // 6: proto.RESULT_TYPE
	(VIEW_TYPE)(0),     // 7: proto.VIEW_TYPE
}
var file_msgenum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgenum_proto_init() }
func file_msgenum_proto_init() {
	if File_msgenum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgenum_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgenum_proto_goTypes,
		DependencyIndexes: file_msgenum_proto_depIdxs,
		EnumInfos:         file_msgenum_proto_enumTypes,
	}.Build()
	File_msgenum_proto = out.File
	file_msgenum_proto_rawDesc = nil
	file_msgenum_proto_goTypes = nil
	file_msgenum_proto_depIdxs = nil
}
