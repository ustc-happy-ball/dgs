package game

import (
	"dgs/configs"
	"dgs/internal/prop"
	"dgs/model"
	"dgs/tools"
	"time"
)

func (g *GameRoom) AdjustPropsIntoTower(props []*model.Prop) {
	towers := g.GetTowers()
	for _, prop := range props {
		if prop.Status == configs.PropStatusDead {
			continue
		}
		towerId := tools.CalTowerId(prop.Pos.X, prop.Pos.Y)
		prop.TowerId = towerId
		towers[towerId].PropEnter(prop)
		//fmt.Printf("把编号为%d的道具放入%d号灯塔中\n, 该灯塔的坐标为X:%f, Y:%f \n", prop.ID(), towerId, prop.GetX(), prop.GetY())
	}
	// todo 通知灯塔中的玩家道具信息
	go g.NotifyHeroPropMsg()
}

func (g *GameRoom) InitNewProps() {
	if g.props.GetPropsCount() >= configs.MaxPropsCountInMap {
		return
	}
	newProps := prop.NewProps(configs.PeriodicPropsInitCount)
	g.props.AddProps(newProps)
	g.AdjustPropsIntoTower(newProps)
}

func (g *GameRoom) PeriodicalInitProps() {
	for {
		select {
		case <-g.die:
			return
		default:
			g.InitNewProps()
			time.Sleep(5 * time.Second) //睡15s
		}
	}
}
