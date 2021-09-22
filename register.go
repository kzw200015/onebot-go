package main

// OnMessage 注册消息 EventHandlerFunc
func (bot *Bot) OnMessage(handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	bot.OnMessageWithPrefix("", handler, middlewares...)
}

// OnMessageWithPrefix 匹配前缀注册消息 EventHandlerFunc
func (bot *Bot) OnMessageWithPrefix(prefix string, handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	bot.OnPrivateMessageWithPrefix(prefix, handler, middlewares...)
	bot.OnGroupMessageWithPrefix(prefix, handler, middlewares...)
}

// OnPrivateMessage 注册私聊消息 EventHandlerFunc
func (bot *Bot) OnPrivateMessage(handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	bot.OnPrivateMessageWithPrefix("", handler, middlewares...)
}

// OnPrivateMessageWithPrefix 匹配前缀注册私聊消息 EventHandlerFunc
func (bot *Bot) OnPrivateMessageWithPrefix(prefix string, handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	h := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	bot.handlerMap.privateMessage.Set(prefix, h)
}

// OnGroupMessage 注册群聊消息 EventHandlerFunc
func (bot *Bot) OnGroupMessage(handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	bot.OnGroupMessageWithPrefix("", handler, middlewares...)
}

// OnGroupMessageWithPrefix 匹配前缀注册群聊消息 EventHandlerFunc
func (bot *Bot) OnGroupMessageWithPrefix(prefix string, handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	h := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	bot.handlerMap.groupMessage.Set(prefix, h)
}
