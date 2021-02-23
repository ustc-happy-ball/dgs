package game

import (
	"fmt"
	pb "github.com/LILILIhuahuahua/ustc_tencent_game/api/proto"
	"github.com/LILILIhuahuahua/ustc_tencent_game/configs"
	event2 "github.com/LILILIhuahuahua/ustc_tencent_game/internal/event"
	"github.com/LILILIhuahuahua/ustc_tencent_game/internal/event/info"
	notify2 "github.com/LILILIhuahuahua/ustc_tencent_game/internal/event/notify"
	"github.com/LILILIhuahuahua/ustc_tencent_game/internal/scheduler"
	"github.com/golang/protobuf/proto"
	"log"
)

type GameRoomManager struct {
	roomMap map[int64]*GameRoom
	timer   *scheduler.Timer
}

var GAME_ROOM_MANAGER *GameRoomManager

func init() {
	GAME_ROOM_MANAGER = NewGameRoomManager()
	GAME_ROOM_MANAGER.timer = scheduler.NewTimer(configs.GlobalInfoNotifyInterval, GlobalInfoNotify)
}

func NewGameRoomManager() *GameRoomManager {
	return &GameRoomManager{
		roomMap: make(map[int64]*GameRoom),
	}
}

// Cron initialize timer for GameRoomManager. It will broadcast props/food info of each room to its clients.
func GlobalInfoNotify() {
	// set cron function for room
	log.Printf("Len of roomMap: %d", len(GAME_ROOM_MANAGER.roomMap))
	if GAME_ROOM_MANAGER != nil && len(GAME_ROOM_MANAGER.roomMap) != 0 {
		for _, room := range GAME_ROOM_MANAGER.roomMap {
			var pbMsg *pb.GMessage
			ps, err := room.props.GetProps()
			if err != nil {
				log.Printf("GlobalInfoNotify: %s", err.Error())
			}

			items := make([]info.ItemInfo, len(ps))
			for k, v := range ps {
				items[k] = info.ItemInfo{
					ID:     v.ID(),
					Type:   int32(pb.ENTITY_TYPE_FOOD_TYPE),
					Status: v.Status(),
					ItemPosition: info.CoordinateXYInfo{
						CoordinateX: v.GetX(),
						CoordinateY: v.GetY(),
					},
				}
			}

			var heroNum int32
			room.sessions.Range(func(_, _ interface{}) bool {
				heroNum++
				return true
			})
			notify := notify2.GameGlobalInfoNotify{
				HeroNumber: heroNum,
				//Time:       0,
				//HeroInfos:  nil,
				ItemInfos: items,
				//MapInfo:    info.MapInfo{},
			}
			msg := event2.GMessage{
				MsgType:     configs.MsgTypeNotify,
				GameMsgCode: configs.GameGlobalInfoNotify,
				//SessionId:   this.room.,
				Data: &notify,
			}

			pbMsg = msg.ToMessage().(*pb.GMessage)
			data, err := proto.Marshal(pbMsg)
			if err != nil {
				log.Printf("fail to marshal: %s", err.Error())
			}

			GAME_ROOM_MANAGER.Braodcast(room.GetRoomID(), data)
		}
	}
}

func (m *GameRoomManager) FetchGameRoom(id int64) *GameRoom {
	return m.roomMap[id]
}

func (m *GameRoomManager) RegisterGameRoom(room *GameRoom) {
	if nil == m.roomMap[room.GetRoomID()] {
		m.roomMap[room.GetRoomID()] = room
	}
}

func (m *GameRoomManager) Unicast(roomId int64, sessionId int32, buff []byte) {
	r := m.FetchGameRoom(roomId)
	s := r.FetchConnector(sessionId)
	s.SendMessage(buff)
	//m.FetchGameRoom(roomId).FetchConnector(sessionId).SendMessage(buff)
	println("hello")
}

func (m *GameRoomManager) Braodcast(roomId int64, buff []byte) {
	r := m.FetchGameRoom(roomId)
	r.BroadcastAll(buff)
	//m.FetchGameRoom(roomId).FetchConnector(sessionId).SendMessage(buff)
}

func (m *GameRoomManager) DeleteUnavailableSession() {
	for _, room := range m.roomMap {
		err := room.DeleteUnavailableSession()
		if err != nil {
			fmt.Println("清理不可用session的时候发生了error: ", err.Error())
		}
	}
}

func (m *GameRoomManager) DeleteDeprecatedHero() {
	for _, room := range m.roomMap {
		err := room.DeleteOfflinePlayer()
		if err != nil {
			fmt.Println("清理废弃玩家的小球时发生了error: ", err.Error())
		}
	}
}