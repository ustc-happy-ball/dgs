package matchService

import (
	"context"
	"dgs/internal/game"
	"dgs/match/matchGrpc"
)

type roomGrpcServer struct {
	matchGrpc.UnimplementedRoomServiceServer
}


func GetRoomServer() *roomGrpcServer {
	return &roomGrpcServer{}
}

func (rs *roomGrpcServer) CreateRoom(context.Context, *matchGrpc.CreateRoomRequest) (*matchGrpc.CreateRoomResponse, error) {
	newRoomId := game.GAME_ROOM_MANAGER.NewGameRoom()
	response := &matchGrpc.CreateRoomResponse{RoomId: newRoomId}
	return response, nil
}
func (rs *roomGrpcServer) GetRoomList(context.Context, *matchGrpc.GetRoomListRequest) (*matchGrpc.GetRoomListResponse, error) {
	roomList := game.GAME_ROOM_MANAGER.GetRoomList()
	var roomMsgList []*matchGrpc.RoomMsg
	for _, room := range roomList {
		roomInfo := room.ToInfo()
		roomMsgList = append(roomMsgList, roomInfo.ToMessage().(*matchGrpc.RoomMsg))
	}

	response := &matchGrpc.GetRoomListResponse{RoomList: roomMsgList}
	return response, nil
}