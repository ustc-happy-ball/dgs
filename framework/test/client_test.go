package test

import (
	pb "dgs/api/proto"
	"dgs/configs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

// deprecated
// TestConnect check whether the client could connect to server correctly
//func TestConnect(t *testing.T) {
//	// wait for server to become ready
//	time.Sleep(time.Second)
//
//	msg := event.GMessage{
//		MsgType:     configs.MsgTypeRequest,
//		GameMsgCode: configs.EntityInfoChangeRequest,
//		SessionId:   0,
//		SeqId:       0,
//		Data: &request.EntityInfoChangeRequest{
//			EventType:  0,
//			HeroId:     0,
//			LinkedId:   0,
//			LinkedType: "",
//			HeroMsg:    info.HeroInfo{},
//		},
//	}
//
//	pbMsg := msg.ToMessage().(*pb.GMessage)
//	out, err := proto.Marshal(pbMsg)
//	if err != nil {
//		t.Fatalf("get %s", err)
//	}
//	var buf []byte
//	// dial to the echo server
//	if sess, err := kcp.DialWithOptions(addr, nil, 0, 0); err == nil {
//		for {
//			data := time.Now().String()
//			log.Println("[Client SENT]: ", data)
//			if _, err := sess.Write(out); err != nil {
//				log.Println(err.Error())
//			}
//			if _, err := io.ReadFull(sess, buf); err == nil {
//				log.Println("get from server", string(buf))
//			} else {
//				log.Println(err.Error())
//			}
//			time.Sleep(time.Second)
//		}
//	} else {
//		log.Fatal(err)
//	}
//}

// getEnterGameReq generate an enterGameRequest and return to caller
func getEnterGameReq(sessionID int32, client *pb.ConnectMsg) []byte {
	playerID := sessionID
	enterGameReq := &pb.GMessage{
		MsgType: pb.MSG_TYPE_REQUEST,
		MsgCode: pb.GAME_MSG_CODE_ENTER_GAME_REQUEST,
		Request: &pb.Request{
			EnterGameRequest: &pb.EnterGameRequest{
				PlayerId:         playerID,
				ClientConnectMsg: client,
			},
		},
	}
	enterGameReqData, err := proto.Marshal(enterGameReq)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	return enterGameReqData
}

// getEntityInfoChangeReq generate an entityInfoChangeRequest to its caller.
// It generate
func getEntityInfoChangeReq(sessionID int32, heroID int32) []byte {
	var entityChangeReq = &pb.EntityInfoChangeRequest{
		HeroId:    heroID,
		EventType: pb.EVENT_TYPE_HERO_MOVE,
	}
	hMsg := &pb.HeroMsg{
		HeroId:        heroID,
		HeroPosition:  &pb.CoordinateXY{},
		HeroDirection: &pb.CoordinateXY{},
		HeroSpeed:     1000,
	}

	randX := configs.MapMaxX - configs.MapMinX
	randY := configs.MapMaxY - configs.MapMinY
	randDirX := float32(rand.Int31n(2) - 1)
	randDirY := float32(rand.Int31n(2) - 1)
	var moveSpeed float32 = 5

	hMsg.HeroPosition.CoordinateX = randX + moveSpeed
	hMsg.HeroPosition.CoordinateY = randY + moveSpeed
	hMsg.HeroDirection.CoordinateX = randDirX
	hMsg.HeroDirection.CoordinateY = randDirY
	entityChangeReq.HeroMsg = hMsg

	msg1 := &pb.GMessage{}
	msg1.SessionId = sessionID
	msg1.MsgCode = pb.GAME_MSG_CODE_ENTITY_INFO_CHANGE_REQUEST
	msg1.MsgType = pb.MSG_TYPE_REQUEST
	r := &pb.Request{}
	r.EntityChangeRequest = entityChangeReq
	msg1.Request = r
	out1, _ := proto.Marshal(msg1)
	return out1
}

func receive(sess *kcp.UDPSession) {
	for {
		buf := make([]byte, 4096)
		n, err := sess.Read(buf)
		if err != nil {
			log.Printf("%v", err.Error())
		}
		msg := &pb.GMessage{}
		err = proto.Unmarshal(buf[:n], msg)
		if err != nil {
			log.Printf("%v", err.Error())
		}

		switch msg.MsgCode {
		case pb.GAME_MSG_CODE_ENTER_GAME_RESPONSE:
			log.Println("Receive enter game response")
			log.Printf("%+v", msg.Response.EnterGameResponse)
			break
		case pb.GAME_MSG_CODE_GAME_GLOBAL_INFO_NOTIFY:
			items := msg.Notify.GameGlobalInfoNotify.ItemMsg
			log.Println("Receive game global info notify")
			log.Printf("Item info: type %v, id %v,status %v", items[0].ItemType, items[0].ItemId, items[0].ItemStatus)
			break
		case pb.GAME_MSG_CODE_ENTITY_INFO_NOTIFY:
			log.Println("Receive entity info change notify")
			log.Printf("")
			break
		case pb.GAME_MSG_CODE_ENTITY_INFO_CHANGE_RESPONSE:
			log.Println("Receive entity info change response")
			log.Printf("%+v", msg.Response.EntityChangeResponse)
		default:
			log.Println(msg.MsgType, " ", msg.MsgType)
			log.Println("New packet received")
		}
	}
}

// ?????????????????????????????????????????????????????????????????????
func TestGlobalPropInfoNotify(t *testing.T) {
	skipCI(t)
	time.Sleep(time.Second)
	var sessionID int32 = 1234
	buf := make([]byte, 4096)

	remoteServ := "127.0.0.1" + ":" + "8888"
	_ = remoteServ
	sess, err := kcp.DialWithOptions(remoteServ, nil, 0, 0)
	if err == nil {
		enterGameReq := getEnterGameReq(sessionID, &pb.ConnectMsg{
			Ip:   "127.0.0.1",
			Port: 4567,
		})

		sess.Write(enterGameReq)
		n, err := sess.Read(buf)
		if err != nil {
			log.Printf("%v", err.Error())
		}
		msg := &pb.GMessage{}
		err = proto.Unmarshal(buf[:n], msg)
		if err != nil {
			log.Println("fail to unmarshal data to GMessage")
		}

		log.Printf("receive %v", msg.MsgCode.String())
		heroID := msg.Response.EnterGameResponse.HeroId
		data := getEntityInfoChangeReq(sessionID, heroID)

		go receive(sess)
		for {
			time.Sleep(200 * time.Millisecond)
			sess.Write(data)
		}
	} else {
		log.Fatal(err)
	}
}

func TestHeroMove1(t *testing.T) {
	skipCI(t)
	var sessionID int32 = 1234
	buf := make([]byte, 4096)

	remoteServ := "127.0.0.1" + ":" + "8888"
	_ = remoteServ
	sess, err := kcp.DialWithOptions(configs.RemoteCLB, nil, 0, 0)
	if err == nil {
		enterGameReq := getEnterGameReq(sessionID, &pb.ConnectMsg{
			Ip:   "127.0.0.1",
			Port: 4567,
		})

		sess.Write(enterGameReq)
		n, err := sess.Read(buf)
		if err != nil {
			log.Printf("%v", err.Error())
		}
		msg := &pb.GMessage{}
		err = proto.Unmarshal(buf[:n], msg)
		if err != nil {
			log.Println("fail to unmarshal data to GMessage")
		}

		log.Printf("receive %v", msg.MsgCode.String())
		heroID := msg.Response.EnterGameResponse.HeroId
		data := getEntityInfoChangeReq(sessionID, heroID)
		sess.Write(data)
	} else {
		log.Fatal(err)
	}
}

func TestHeroMove2(t *testing.T) {
	skipCI(t)
	var sessionID int32 = 4567
	buf := make([]byte, 4096)

	remoteServ := "127.0.0.1" + ":" + "8888"
	_ = remoteServ
	sess, err := kcp.DialWithOptions(configs.RemoteCLB, nil, 0, 0)
	if err == nil {
		enterGameReq := getEnterGameReq(sessionID, &pb.ConnectMsg{
			Ip:   "127.0.0.1",
			Port: 4567,
		})

		sess.Write(enterGameReq)
		n, err := sess.Read(buf)
		if err != nil {
			log.Printf("%v", err.Error())
		}
		msg := &pb.GMessage{}
		err = proto.Unmarshal(buf[:n], msg)
		if err != nil {
			log.Println("fail to unmarshal data to GMessage")
		}

		log.Printf("receive %v", msg.MsgCode.String())
		heroID := msg.Response.EnterGameResponse.HeroId
		fmt.Println("hedo???", heroID)
		data := getEntityInfoChangeReq(sessionID, heroID)
		sess.Write(data)
	} else {
		log.Fatal(err)
	}
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}
