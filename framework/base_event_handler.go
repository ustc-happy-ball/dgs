package framework

import (
	"dgs/framework/event"
	//event2 "dgs/internal/event"
)

var EVENT_HANDLER = &BaseEventHandler{}

type BaseEventHandler struct {
}

func (b BaseEventHandler) OnEvent(e event.Event) {
	if nil == e {
		return
	}
	handler := event.Manager.FetchHandler(e.GetCode())
	if nil != handler {
		handler.OnEvent(e)
	}
}

func (b BaseEventHandler) OnEventToSession(e event.Event, s event.Session) {
	if nil == e {
		return
	}
	handler := event.Manager.FetchHandler(e.GetCode())

	if nil != handler {
		handler.OnEvent(e)
	}
}
