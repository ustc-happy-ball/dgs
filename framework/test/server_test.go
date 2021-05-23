package test

import (
	"testing"
)

var (
	addr = "127.0.0.1:8888"
)

// TestNewKcpServer create a server for test
func TestNewKcpServer(t *testing.T) {
	skipCI(t)
	//b, _ := game.NewGameRoom(addr)
	//b.Serv()
}
