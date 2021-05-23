package db

import (
	"log"
	"os"
	"testing"
)

func TestPlayerUpdateHighestScoreByPlayerId(t *testing.T) {
	skipCI(t)
	InitConnection("localhost:8890")
	err := PlayerUpdateHighestScoreByPlayerId(560, 2333)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("成功了")
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}
