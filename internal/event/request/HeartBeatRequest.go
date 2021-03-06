package request

import (
	pb "dgs/api/proto"
	"dgs/framework"
	"dgs/framework/event"
)

type HeartBeatRequest struct {
	framework.BaseEvent
	SendTime int64
}

func (e *HeartBeatRequest) ToMessage() interface{} {
	return &pb.HeartBeatRequest{
		SendTime: e.SendTime,
	}
}

func (e *HeartBeatRequest) FromMessage(obj interface{}) {
	pbMsg := obj.(*pb.HeartBeatRequest)
	e.SetCode(int32(pb.GAME_MSG_CODE_HEART_BEAT_REQUEST))
	e.SendTime = pbMsg.GetSendTime()
}

func (e *HeartBeatRequest) CopyFromMessage(obj interface{}) event.Event {
	pbMsg := obj.(*pb.Request).HeartBeatRequest
	req := &HeartBeatRequest{
		SendTime: pbMsg.GetSendTime(),
	}
	req.SetCode(int32(pb.GAME_MSG_CODE_HEART_BEAT_REQUEST))
	return req
}
