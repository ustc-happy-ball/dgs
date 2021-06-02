package db

import (
	"log"
	"testing"
)

func TestAccountGetAccountInfoByPlayerId(t *testing.T) {
	skipCI(t)
	InitConnection("localhost:8890")
	account, err := AccountGetAccountInfoByPlayerId(5500)
	if err != nil {
		log.Fatalln(err)
	}
	t.Log(account)
	log.Println("成功了")
}
