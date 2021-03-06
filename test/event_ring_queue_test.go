package test

import (
	"dgs/framework/event"
	"dgs/internal/event/notify"
	"fmt"
	"testing"
)

func TestEventRingQueue(t *testing.T) {
	e1 := notify.NewGameFinishNotify(nil, 0)
	queue := event.NewEventRingQueue(10)
	for i := 0; i < 15; i++ {
		err := queue.Push(e1)
		if err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < 15; i++ {
		_, err := queue.Pop()
		if err != nil {
			fmt.Println(err)
		}
	}
}
