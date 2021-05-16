package info

import (
	pb "dgs/api/proto"
	"dgs/framework"
	"dgs/framework/event"
	"dgs/model"
)

type HeroRankInfo struct {
	framework.BaseEvent
	HeroID    int32
	HeroName  string
	HeroRank  int32
	HeroScore int32
}

func NewHeroRankInfo(hero *model.Hero) *HeroRankInfo {
	return &HeroRankInfo{
		HeroID:    hero.ID,
		HeroName:  hero.Name,
		HeroRank:  hero.Rank,
		HeroScore: hero.Score,
	}
}

func (h *HeroRankInfo) FromMessage(obj interface{}) {
	pbMsg := obj.(*pb.HeroRankMsg)
	h.HeroID = pbMsg.GetHeroId()
	h.HeroName = pbMsg.GetHeroName()
	h.HeroRank = pbMsg.GetHeroRank()
	h.HeroScore = pbMsg.GetHeroScore()
}

func (h *HeroRankInfo) CopyFromMessage(obj interface{}) event.Event {
	pbMsg := obj.(*pb.HeroRankMsg)
	return &HeroRankInfo{
		HeroID:    pbMsg.GetHeroId(),
		HeroName:  pbMsg.GetHeroName(),
		HeroRank:  pbMsg.GetHeroRank(),
		HeroScore: pbMsg.GetHeroScore(),
	}
}

func (h *HeroRankInfo) ToMessage() interface{} {
	pbMsg := &pb.HeroRankMsg{
		HeroId:    h.HeroID,
		HeroName:  h.HeroName,
		HeroRank:  h.HeroRank,
		HeroScore: h.HeroScore,
	}
	return pbMsg
}
