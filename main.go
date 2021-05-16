package main

import (
	"dgs/configs"
	"dgs/db"
	"dgs/internal/game"
	"flag"
	"log"
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

	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	log.Println("[USTC-Tencent]Game Server Started!")
	s := game.NewGameStarter(configs.ServerAddr)
	s.Boot()
}
