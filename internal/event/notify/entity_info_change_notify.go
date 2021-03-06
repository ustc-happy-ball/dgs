package notify

import (
	//pb "dgs/api/proto"
	pb "dgs/api/proto"
	"dgs/framework"
	e "dgs/framework/event"
	"dgs/internal/event/info"
	"dgs/tools"
	"github.com/golang/protobuf/proto"
)

//message EntityInfoChangeNotify {
//ENTITY_TYPE entityType = 1; //实体的类型
//int32 entityId = 2; //实体的Id
//HeroMsg HeroMsg = 3; //玩家的信息
//ItemMsg itemMsg = 4; //物体的信息
//}

type EntityInfoChangeNotify struct {
	framework.BaseEvent
	EntityType int32
	EntityId   int32
	HeroMsg    *info.HeroInfo
	ItemMsg    *info.ItemInfo
}

func NewEntityInfoChangeNotify(entityType int32, entityId int32, heroInfo *info.HeroInfo, itemInfo *info.ItemInfo) *EntityInfoChangeNotify {
	return &EntityInfoChangeNotify{
		EntityType: entityType,
		EntityId:   entityId,
		HeroMsg:    heroInfo,
		ItemMsg:    itemInfo,
	}
}

func (e *EntityInfoChangeNotify) FromMessage(obj interface{}) {
	pbMsg := obj.(*pb.EntityInfoChangeNotify)
	e.SetCode(int32(pb.GAME_MSG_CODE_ENTITY_INFO_CHANGE_NOTIFY))
	e.EntityType = int32(pbMsg.EntityType)
	e.EntityId = pbMsg.EntityId
	//hero
	pbHeroMsg := pbMsg.HeroMsg
	if nil != pbHeroMsg {
		e.HeroMsg = &info.HeroInfo{}
		e.HeroMsg.FromMessage(pbHeroMsg)
	}
	//item
	pbItemMsg := pbMsg.ItemMsg
	if nil != pbItemMsg {
		e.ItemMsg = &info.ItemInfo{}
		e.ItemMsg.FromMessage(pbItemMsg)
	}
}

func (notify *EntityInfoChangeNotify) CopyFromMessage(obj interface{}) e.Event {
	pbMsg := obj.(*pb.Notify).EntityInfoChangeNotify
	n := &EntityInfoChangeNotify{
		EntityId:   pbMsg.GetEntityId(),
		EntityType: int32(pbMsg.GetEntityType()),
	}
	//hero
	pbHeroMsg := pbMsg.HeroMsg
	if nil != pbHeroMsg {
		n.HeroMsg = &info.HeroInfo{}
		n.HeroMsg.FromMessage(pbHeroMsg)
	}
	//item
	pbItemMsg := pbMsg.ItemMsg
	if nil != pbItemMsg {
		n.ItemMsg = &info.ItemInfo{}
		n.ItemMsg.FromMessage(pbItemMsg)
	}
	n.SetCode(int32(pb.GAME_MSG_CODE_ENTITY_INFO_CHANGE_NOTIFY))
	return n
}

func (notify *EntityInfoChangeNotify) ToMessage() interface{} {
	pbMsg := &pb.EntityInfoChangeNotify{
		EntityType: pb.ENTITY_TYPE(notify.EntityType),
		EntityId:   notify.EntityId,
		//HeroMsg:    notify.HeroMsg.ToMessage().(*pb.HeroMsg),
		//ItemMsg:    notify.ItemMsg.ToMessage().(*pb.ItemMsg),
	}
	if nil != notify.HeroMsg {
		pbMsg.HeroMsg = notify.HeroMsg.ToMessage().(*pb.HeroMsg)
	}
	if nil != notify.ItemMsg {
		pbMsg.ItemMsg = notify.ItemMsg.ToMessage().(*pb.ItemMsg)
	}
	return pbMsg
}

func (notify *EntityInfoChangeNotify) ToGMessageBytes() []byte {
	n := &pb.Notify{
		EntityInfoChangeNotify: notify.ToMessage().(*pb.EntityInfoChangeNotify),
	}
	msg := pb.GMessage{
		MsgType:  pb.MSG_TYPE_NOTIFY,
		MsgCode:  pb.GAME_MSG_CODE_ENTITY_INFO_CHANGE_NOTIFY,
		Notify:   n,
		SendTime: tools.TIME_UTIL.NowMillis(),
	}
	out, _ := proto.Marshal(&msg)
	return out
}
