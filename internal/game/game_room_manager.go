package game

import (
	pb "dgs/api/proto"
	"dgs/configs"
	"dgs/framework"
	"dgs/framework/kcpnet"
	event2 "dgs/internal/event"
	"dgs/internal/event/info"
	notify2 "dgs/internal/event/notify"
	"dgs/internal/event/request"
	"dgs/internal/scheduler"
	"dgs/model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type GameRoomManager struct {
	roomMap          sync.Map //游戏房间集合
	waitedSessionMap sync.Map //等待匹配会话集合
	timer            *scheduler.Timer
	server           *kcpnet.KcpServer
	waitedSessionNum int32
}

//func init() {
//	GAME_ROOM_MANAGER = NewGameRoomManager()
//	GAME_ROOM_MANAGER.timer = scheduler.NewTimer(configs.GlobalInfoNotifyInterval, GlobalInfoNotify)
//}

func NewGameRoomManager() *GameRoomManager {
	s, err := kcpnet.NewKcpServer(configs.ServerAddr)
	if err != nil {
		return nil
	}
	manager := &GameRoomManager{
		//roomMap: make(map[int64]*GameRoom),
		server:           s,
		waitedSessionNum: 0,
	}
	//manager.timer = scheduler.NewTimer(configs.GlobalInfoNotifyInterval, GlobalInfoNotify)
	return manager
}

func (manager *GameRoomManager) Serv() {
	log.Println("[GameRoomManager]游戏房间管理器开始监听新连接！")
	//go manager.waitForMatching()
	go manager.waitForEnterGame()
	for {
		conn, err := manager.server.Listen.AcceptKCP()
		if err != nil {
			log.Println("[GameRoomManager]kcp接收会话出错！")
		}
		//接收新的kcp连接
		conn.SetWindowSize(4800, 4800)
		session := framework.NewBaseSession(conn)
		log.Printf("[GameRoomManager]监听到新连接！%v \n", session)
		atomic.AddInt32(&manager.waitedSessionNum, 1)
		manager.waitedSessionMap.Store(session.Id, session)
	}
}

func (manager *GameRoomManager) NewGameRoom() int64 {
	newRoom := NewGameRoom()
	manager.roomMap.Store(newRoom.ID, newRoom)
	return newRoom.ID
}

func (manager *GameRoomManager) GetRoomList() []*GameRoom {
	roomList := []*GameRoom{}
	manager.roomMap.Range(func(key, value interface{}) bool {
		room := value.(*GameRoom)
		roomList = append(roomList, room)
		return true
	})
	return roomList
}

func (manager *GameRoomManager) waitForEnterGame() {
	log.Println("[GameRoomManager]游戏房间管理器开始轮询对等待匹配会话进行进入房间处理!")
	buf := make([]byte, 4096)
	for {
		manager.waitedSessionMap.Range(func(_, value interface{}) bool {
			session := value.(*framework.BaseSession)
			//给session加一个读超时函数，及时释放锁
			err := session.Sess.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(2)))
			if err != nil {
				panic("setDeadLine出错")
			}
			num, _ := session.Sess.Read(buf)
			if num == 0 {
				return true
			}
			session.UpdateTime()
			pbMsg := &pb.GMessage{}
			proto.Unmarshal(buf[:num], pbMsg)
			msg := event2.GMessage{}
			m := msg.CopyFromMessage(pbMsg)
			if m.GetCode() != int32(pb.GAME_MSG_CODE_ENTER_GAME_REQUEST) {
				fmt.Printf("[GameRoomManager]进入对局时收到非法请求！\n")
				return true
			}
			//开始进行进入世界处理
			enterGameReq := m.(*event2.GMessage).Data.(*request.EnterGameRequest)
			//绑定会话至玩家
			roomId := enterGameReq.RoomID
			room := manager.FetchGameRoom(roomId)
			if nil == room {
				fmt.Printf("[GameRoomManager]进入对局时未找到该房间！检查玩家上送roomID是否合法！roomID:%v\n", roomId)
				return true
			}
			//下放玩家至该房间
			fmt.Printf("[GameRoomManager]将玩家分派到房间成功！playerID:%v, room:%+v\n", enterGameReq.PlayerID, room)
			manager.waitedSessionMap.Delete(session.Id)
			room.OnEnterGame(m.(*event2.GMessage), session)
			atomic.AddInt32(&manager.waitedSessionNum, -1)
			return true
		})
	}
}

// TODO 优化： CPU
// sync.Map 是线程安全的 map，可以一边遍历，一边删除
//func (manager *GameRoomManager) waitForMatching() {
//	log.Println("[GameRoomManager]游戏房间管理器开始轮询对等待匹配会话进行匹配!")
//	for {
//		manager.waitedSessionMap.Range(func(_, value interface{}) bool {
//			session := value.(*framework.BaseSession)
//			//1.达到匹配人数 或者 2.等待超时
//			if manager.waitedSessionNum >= configs.MinMatchingBatchSessionNum || manager.isWaitedSessionOvertime(session) {
//				//选择一个适合的房间，将新会话放入房间的未注册会话集合中
//				room, err := manager.matchingSessionToRoom(session)
//				if nil != err {
//					log.Println("[GameRoomManager]匹配新会话至房间时出错！")
//				}
//				room.AcceptConnector(session)
//				manager.waitedSessionMap.Delete(session.Id)
//				atomic.AddInt32(&manager.waitedSessionNum, -1)
//			}
//			return true
//		})
//	}
//}
//
//func (manager *GameRoomManager) isWaitedSessionOvertime(session *framework.BaseSession) bool {
//	nowTime := time.Now().UnixNano()
//	res := (nowTime-session.CreationTime)/1e9 >= configs.MatchingWaitOverTime
//	return res
//}
//
//func (manager *GameRoomManager) matchingSessionToRoom(session *framework.BaseSession) (*GameRoom, error) {
//	//1.满足房间人数限制要求
//	//2.满足房间中的最高分限制要求（最高分不超过游戏胜利分数的30%）
//	var targetRoom *GameRoom
//	manager.roomMap.Range(func(_, value interface{}) bool {
//		room := value.(*GameRoom)
//		if room.AliveHeroNum < configs.GameAliveHeroLimit {
//			targetRoom = room
//			return true
//		}
//		return true
//	})
//	//3.若没有找到上述的房间，则新创建一个房间
//	if nil == targetRoom {
//		log.Printf("[GameRoomManager]匹配会话时未找到合适的对局，新建一个新的对局！")
//		targetRoom = NewGameRoom()
//		go targetRoom.Serv()
//		manager.roomMap.Store(targetRoom.ID, targetRoom)
//	}
//	log.Printf("[GameRoomManager]会话匹配成功！session：%v, room: %v \n", session, targetRoom)
//
//	return targetRoom, nil
//}

// Cron initialize timer for GameRoomManager. It will broadcast props/food info of each room to its clients.
func GlobalInfoNotify() {
	// set cron function for room
	//log.Printf("Len of roomMap: %d", len(GAME_ROOM_MANAGER.roomMap))
	GAME_ROOM_MANAGER.roomMap.Range(func(_, value interface{}) bool {
		room := value.(*GameRoom)
		var pbMsg *pb.GMessage
		roomHeros := room.GetHeros()
		// 该房间内所有的小球
		var heroNeedToNotify []*model.Hero
		roomHeros.Range(func(k, v interface{}) bool {
			heroNeedToNotify = append(heroNeedToNotify, v.(*model.Hero))
			return true
		})
		for _, hero := range heroNeedToNotify {
			var heroInfos []info.HeroInfo
			players, props := room.GetItemsNearby(hero)
			var items []info.ItemInfo
			for _, v := range props {
				itemInfo := info.ItemInfo{
					ID:     v.Id,
					Type:   v.PropType,
					Status: v.Status,
					ItemPosition: &info.CoordinateXYInfo{
						CoordinateX: v.Pos.X,
						CoordinateY: v.Pos.Y,
					},
				}
				items = append(items, itemInfo)
			}
			for _, player := range players {
				//heroInfo := player.ToEvent()
				heroInfo := *info.NewHeroInfo(player)
				heroInfos = append(heroInfos, heroInfo)
			}
			heroNum := int32(len(players))
			notify := notify2.GameGlobalInfoNotify{
				HeroNumber: heroNum,
				//Time:       0,
				HeroInfos: heroInfos,
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
			GAME_ROOM_MANAGER.Unicast(room.GetRoomID(), hero.Session.Id, data)
		}
		return true
	})
	//if GAME_ROOM_MANAGER != nil && len(GAME_ROOM_MANAGER.roomMap) != 0 {
	//	for _, room := range GAME_ROOM_MANAGER.roomMap {
	//		var pbMsg *pb.GMessage
	//
	//		roomHeros := room.GetHeros()
	//		// 该房间内所有的小球
	//		var heroNeedToNotify []*model.Hero
	//		roomHeros.Range(func(k, v interface{}) bool {
	//			heroNeedToNotify = append(heroNeedToNotify, v.(*model.Hero))
	//			return true
	//		})
	//		for _, hero := range heroNeedToNotify {
	//			var heroInfos []info.HeroInfo
	//			players, props := room.GetItemsNearby(hero)
	//			var items []info.ItemInfo
	//			for _, v := range props {
	//				itemInfo := info.ItemInfo{
	//					ID:     v.Id,
	//					Type:   v.PropType,
	//					Status: v.Status,
	//					ItemPosition: &info.CoordinateXYInfo{
	//						CoordinateX: v.Pos.X,
	//						CoordinateY: v.Pos.Y,
	//					},
	//				}
	//				items = append(items, itemInfo)
	//			}
	//			for _, player := range players {
	//				//heroInfo := player.ToEvent()
	//				heroInfo := *info.NewHeroInfo(player)
	//				heroInfos = append(heroInfos, heroInfo)
	//			}
	//			heroNum := int32(len(players))
	//			notify := notify2.GameGlobalInfoNotify{
	//				HeroNumber: heroNum,
	//				//Time:       0,
	//				HeroInfos: heroInfos,
	//				ItemInfos: items,
	//				//MapInfo:    info.MapInfo{},
	//			}
	//			msg := event2.GMessage{
	//				MsgType:     configs.MsgTypeNotify,
	//				GameMsgCode: configs.GameGlobalInfoNotify,
	//				//SessionId:   this.room.,
	//				Data: &notify,
	//			}
	//
	//			pbMsg = msg.ToMessage().(*pb.GMessage)
	//			data, err := proto.Marshal(pbMsg)
	//			if err != nil {
	//				log.Printf("fail to marshal: %s", err.Error())
	//			}
	//			GAME_ROOM_MANAGER.Unicast(room.GetRoomID(), hero.Session.Id, data)
	//		}
	//
	//	}
	//}
}

func (m *GameRoomManager) FetchGameRoom(id int64) *GameRoom {
	room, _ := m.roomMap.Load(id)
	if nil != room {
		return room.(*GameRoom)
	}
	return nil
}

func (m *GameRoomManager) RegisterGameRoom(room *GameRoom) {
	m.roomMap.Store(room.ID, room)
}

func (m *GameRoomManager) DeleteGameRoom(roomId int64) {
	log.Printf("[GameRoomManager]删除房间对局！roomID：%v \n", roomId)
	m.roomMap.Delete(roomId)
}

func (m *GameRoomManager) Unicast(roomId int64, sessionId int32, buff []byte) {
	defer func() {
		e := recover()
		if e != nil {
			//fmt.Println("在unicast的时候出错了，错误为：", e)
		}
	}()

	r := m.FetchGameRoom(roomId)
	s := r.FetchConnector(sessionId)
	if s == nil {
		panic("没有该玩家")
	}
	err := s.SendMessage(buff)
	if err != nil {
		panic(err)
	}
	//m.FetchGameRoom(roomId).FetchConnector(sessionId).SendMessage(buff)
}

func (m *GameRoomManager) Braodcast(roomId int64, buff []byte) {
	r := m.FetchGameRoom(roomId)
	if nil != r {
		r.BroadcastAll(buff)
	}
	//m.FetchGameRoom(roomId).FetchConnector(sessionId).SendMessage(buff)
}

func (m *GameRoomManager) MutiplecastToNearBy(roomId int64, buf []byte, hero *model.Hero) {
	r := m.FetchGameRoom(roomId)
	var sessionToSend []*framework.BaseSession
	heroToSend, _ := r.GetItemsNearby(hero)
	for _, hero := range heroToSend {
		if hero.Session != nil && hero.Session.Status == configs.SessionStatusCreated {
			sessionToSend = append(sessionToSend, hero.Session)
		} else {
			fmt.Println("hero的session为null")
		}
	}
	r.Multiplecast(buf, sessionToSend)
}

func (m *GameRoomManager) DeleteUnavailableSession() {
	m.roomMap.Range(func(_, value interface{}) bool {
		room := value.(*GameRoom)
		err := room.DeleteUnavailableSession()
		if err != nil {
			fmt.Println("清理不可用session的时候发生了error: ", err.Error())
		}
		return true
	})
	//for _, room := range m.roomMap {
	//	err := room.DeleteUnavailableSession()
	//	if err != nil {
	//		fmt.Println("清理不可用session的时候发生了error: ", err.Error())
	//	}
	//}
}

func (m *GameRoomManager) DeleteDeprecatedHero() {
	m.roomMap.Range(func(_, value interface{}) bool {
		room := value.(*GameRoom)
		err := room.DeleteOfflinePlayer()
		if err != nil {
			fmt.Println("清理废弃玩家的小球时发生了error: ", err.Error())
		}
		return true
	})
	//for _, room := range m.roomMap {
	//	err := room.DeleteOfflinePlayer()
	//	if err != nil {
	//		fmt.Println("清理废弃玩家的小球时发生了error: ", err.Error())
	//	}
	//}
}

func (m *GameRoomManager) UpdateHeroPositionAndStatus() {
	m.roomMap.Range(func(_, value interface{}) bool {
		room := value.(*GameRoom)
		room.UpdateHeroPosAndStatus()
		return true
	})
	//for _, room := range m.roomMap {
	//	room.UpdateHeroPosAndStatus()
	//}
}
