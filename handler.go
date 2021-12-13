package onebot

import (
	"strings"
	"sync"
)

type EventHandlerFunc func(*Bot, Event)

type EventHandler struct {
	prefix string
	fn     EventHandlerFunc
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

func (eh EventHandler) Handle(bot *Bot, event Event) {
	if strings.HasPrefix(event.RawMessage, eh.prefix) {
		eh.fn(bot, event)
	}
}
