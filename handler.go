package onebot

import (
	"sync"
)

type EventHandlerFunc func(*Context)

type EventHandler struct {
	MessagePrefix string
	MessageType   string
	Handle        EventHandlerFunc
}

type HandlerMap struct {
	*sync.Map
}

func (hm *HandlerMap) GetHandlers(postType string) []EventHandler {
	handlers, ok := hm.Load(postType)
	if !ok {
		return []EventHandler{}
	}
	return handlers.([]EventHandler)
}
