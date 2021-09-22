package onebot

import (
	"strings"
	"sync"
)

// EventHandlerFunc 消息处理函数
type EventHandlerFunc func(ctx *Context) error

// EventMiddlewareFunc 中间件
type EventMiddlewareFunc func(handlerFunc EventHandlerFunc) EventHandlerFunc

// EventHandlerMap 前缀到 EventHandlerFunc 映射结构体
type EventHandlerMap struct {
	prefixHandlerMap sync.Map
}

// Set 设置前缀到 EventHandlerFunc 的映射
func (e *EventHandlerMap) Set(prefix string, handlers ...EventHandlerFunc) {
	hs, loaded := e.prefixHandlerMap.LoadOrStore(prefix, handlers)
	if loaded {
		hs = append(hs.([]EventHandlerFunc), handlers...)
		e.prefixHandlerMap.Store(prefix, hs)
	}
}

// Get 判断传入字符串是否满足前缀，返回所有满足的 EventHandlerFunc
func (e *EventHandlerMap) Get(message string) []EventHandlerFunc {
	var handlers []EventHandlerFunc
	e.prefixHandlerMap.Range(func(prefix, handler interface{}) bool {
		if strings.HasPrefix(message, prefix.(string)) {
			handlers = append(handlers, handler.([]EventHandlerFunc)...)
		}
		return true
	})

	return handlers
}

// handleEvent 处理事件
func (bot *Bot) handleEvent(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			bot.logger.Error("recover from panic:", err)
		}
	}()

	if event.SelfId != bot.config.SelfId {
		return
	}
	// 根据上报事件类型分发
	switch event.PostType {
	case MessageEvent:
		switch event.MessageType {
		case PrivateMessage:
			bot.logger.Infof("收到私聊消息: %v", event.RawMessage)
			handlers := bot.handlerMap.privateMessage.Get(event.RawMessage)
			bot.executeEventHandler(event, handlers)
		case GroupMessage:
			bot.logger.Infof("收到群聊消息: %v", event.RawMessage)
			handlers := bot.handlerMap.groupMessage.Get(event.RawMessage)
			bot.executeEventHandler(event, handlers)
		}
	case MetaEvent:
	case NoticeEvent:
	case RequestEvent:
	}
}

// executeEventHandler 执行 EventHandlerFunc
func (bot *Bot) executeEventHandler(event Event, handlers []EventHandlerFunc) {
	for _, handler := range handlers {
		err := handler(&Context{
			bot:   bot,
			Event: event,
		})
		if err != nil {
			bot.logger.Warn(err)
		}
	}
}
