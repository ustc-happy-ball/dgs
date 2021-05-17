package info

import (
	"dgs/framework"
	"dgs/framework/event"
	"dgs/match/matchGrpc"
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
	return matchGrpc.RoomMsg{
		Id: r.ID,
		Status: r.Status,
		CreateTime: r.CreateTime,
		PlayerCount: r.PlayerCount,
		HighestScore: r.HighestScore,
	}
}

func (r RoomInfo) FromMessage(obj interface{}) {
	pbMsg := obj.(*matchGrpc.RoomMsg)
	r.ID = pbMsg.Id
	r.Status = pbMsg.Status
	r.CreateTime = pbMsg.CreateTime
	r.PlayerCount = pbMsg.PlayerCount
	r.HighestScore = pbMsg.HighestScore
}

func (r RoomInfo) CopyFromMessage(obj interface{}) event.Event {
	pbMsg := obj.(*matchGrpc.RoomMsg)
	return &RoomInfo{
		ID: pbMsg.Id,
		Status: pbMsg.Status,
		CreateTime: pbMsg.CreateTime,
		PlayerCount: pbMsg.PlayerCount,
		HighestScore: pbMsg.HighestScore,
	}
}
