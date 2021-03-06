package info

import (
	pb "dgs/api/proto"
	"dgs/framework"
	"dgs/framework/event"
	"dgs/model"
)

type HeroInfo struct {
	framework.BaseEvent
	ID            int32
	Status        int32
	Speed         float32
	Size          float32
	Score         int32
	Invincible    bool
	SpeedUp       bool
	SpeedDown     bool
	HeroPosition  *CoordinateXYInfo
	HeroDirection *CoordinateXYInfo
	HeroName      string
}

func NewHeroInfo(hero *model.Hero) *HeroInfo {
	return &HeroInfo{
		ID:            hero.ID,
		Speed:         hero.Speed,
		Size:          hero.Size,
		Status:        hero.Status,
		Score:         hero.Score,
		Invincible:    hero.Invincible,
		SpeedUp:       hero.SpeedUp,
		SpeedDown:     hero.SpeedDown,
		HeroPosition:  NewCoordinateInfo(hero.HeroPosition.X, hero.HeroPosition.Y),
		HeroDirection: NewCoordinateInfo(hero.HeroDirection.X, hero.HeroDirection.Y),
		HeroName:      hero.Name,
	}
}

func (h *HeroInfo) FromMessage(obj interface{}) {
	pbMsg := obj.(*pb.HeroMsg)
	h.ID = pbMsg.GetHeroId()
	h.Status = int32(pbMsg.GetHeroStatus())
	h.Speed = pbMsg.GetHeroSpeed()
	h.Size = pbMsg.HeroSize
	h.Score = pbMsg.HeroScore
	pos := CoordinateXYInfo{}
	pos.FromMessage(pbMsg.GetHeroPosition())
	h.HeroPosition = &pos
	dict := CoordinateXYInfo{}
	dict.FromMessage(pbMsg.GetHeroDirection())
	h.HeroDirection = &dict
	h.HeroName = pbMsg.HeroName
}

func (h *HeroInfo) CopyFromMessage(obj interface{}) event.Event {
	pbMsg := obj.(*pb.HeroMsg)
	pos := CoordinateXYInfo{}
	dict := CoordinateXYInfo{}
	dict.FromMessage(pbMsg.GetHeroDirection())
	pos.FromMessage(pbMsg.GetHeroPosition())
	return &HeroInfo{
		ID:            pbMsg.GetHeroId(),
		Status:        int32(pbMsg.GetHeroStatus()),
		Speed:         pbMsg.GetHeroSpeed(),
		Size:          pbMsg.GetHeroSize(),
		Score:         pbMsg.HeroScore,
		Invincible:    pbMsg.Invincible,
		SpeedUp:       pbMsg.SpeedUp,
		SpeedDown:     pbMsg.SpeedDown,
		HeroPosition:  &pos,
		HeroDirection: &dict,
		HeroName:      pbMsg.HeroName,
	}
}

func (h *HeroInfo) ToMessage() interface{} {
	pbMsg := &pb.HeroMsg{
		HeroId: h.ID,
		//HeroStatus:    h.Status,
		HeroSpeed:     h.Speed,
		HeroSize:      h.Size,
		HeroScore:     h.Score,
		Invincible:    h.Invincible,
		SpeedUp:       h.SpeedUp,
		SpeedDown:     h.SpeedDown,
		HeroPosition:  h.HeroPosition.ToMessage().(*pb.CoordinateXY),
		HeroDirection: h.HeroDirection.ToMessage().(*pb.CoordinateXY),
		HeroName:      h.HeroName,
	}
	switch h.Status {
	case int32(pb.HERO_STATUS_LIVE):
		pbMsg.HeroStatus = pb.HERO_STATUS_LIVE
		break

	case int32(pb.HERO_STATUS_DEAD):
		pbMsg.HeroStatus = pb.HERO_STATUS_DEAD
		break
	}
	return pbMsg
}
