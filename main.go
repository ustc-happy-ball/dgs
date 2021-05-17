package main

import (
	"dgs/configs"
	"dgs/db"
	"dgs/internal/game"
	"dgs/match/matchGrpc"
	"dgs/match/matchService"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "net/http/pprof"
	"runtime"
)

func initDB() {
	var (
		dbProxyPort string
		dbProxyHost string
	)

	flag.StringVar(&dbProxyHost, "DBProxyHost", "", "Host addr of dbproxy")
	flag.StringVar(&dbProxyPort, "DBProxyPort", "", " Port of dbproxy")
	flag.Parse()
	configs.DBProxyAddr = dbProxyHost + ":" + dbProxyPort
}

func main() {
	initDB()
	if configs.DBProxyAddr == "" {
		log.Fatalln("DBProxy addr is nil")
	}
	log.Println("Initialize DBProxyAddr to", configs.DBProxyAddr)
	go db.InitConnection(configs.DBProxyAddr)
	go initGrpcServer()
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	log.Println("[USTC-Tencent]Game Server Started!")
	s := game.NewGameStarter(configs.ServerAddr)
	s.Boot()
}

func initGrpcServer() {
	lis, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		log.Println("[TcpListen] tcp listen failed, the error is ", err.Error())
	}
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)
	log.Println("[GrpcServer] grpc server start")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("[GrpcServer] server start failed, the error is ", err.Error())
	}
}

func registerServer(server *grpc.Server) {
	matchGrpc.RegisterRoomServiceServer(server, matchService.GetRoomServer())
}
