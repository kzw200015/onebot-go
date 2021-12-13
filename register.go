package onebot

func (bot *Bot) On(postType string, handler EventHandler) {
	handlers, isLoaded := bot.handlerMap.LoadOrStore(postType, []EventHandler{handler})
	if isLoaded {
		handlers = append(handlers.([]EventHandler), handler)
	}
}

func (bot *Bot) OnMessage(fn EventHandlerFunc) {
	bot.On(PostTypeMessage, EventHandler{fn: fn})
}

func (bot *Bot) OnNotice(fn EventHandlerFunc) {
	bot.On(PostTypeNotice, EventHandler{fn: fn})
}

func (bot *Bot) OnRequest(fn EventHandlerFunc) {
	bot.On(PostTypeRequest, EventHandler{fn: fn})
}

func (bot *Bot) OnMeta(fn EventHandlerFunc) {
	bot.On(PostTypeMeta, EventHandler{fn: fn})
}
