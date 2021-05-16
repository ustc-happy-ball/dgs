package info

import (
	pb "dgs/api/proto"
	"dgs/framework"
	"dgs/framework/event"
)

type RoomInfo struct {
	framework.BaseEvent
	ID           int64
	Status       int32
	CreateTime   int64
	PlayerCount  int32
	HighestScore int32
}

func (r RoomInfo) ToMessage() interface{} {
	return &pb.RoomMsg{
		Id: r.ID,
		Status: r.Status,
		CreateTime: r.CreateTime,
		PlayerCount: r.PlayerCount,
		HighestScore: r.HighestScore,
	}
}

func (r RoomInfo) FromMessage(obj interface{}) {
	pbMsg := obj.(*pb.RoomMsg)
	r.ID = pbMsg.Id
	r.Status = pbMsg.Status
	r.CreateTime = pbMsg.CreateTime
	r.PlayerCount = pbMsg.PlayerCount
	r.HighestScore = pbMsg.HighestScore
}

func (r RoomInfo) CopyFromMessage(obj interface{}) event.Event {
	pbMsg := obj.(*pb.RoomMsg)
	return &RoomInfo{
		ID: pbMsg.Id,
		Status: pbMsg.Status,
		CreateTime: pbMsg.CreateTime,
		PlayerCount: pbMsg.PlayerCount,
		HighestScore: pbMsg.HighestScore,
	}
}
