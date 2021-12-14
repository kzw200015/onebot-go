package onebot

func (bot *Bot) On(postType string, handler EventHandler) {
	handlers, isLoaded := bot.handlerMap.LoadOrStore(postType, []EventHandler{handler})
	if isLoaded {
		handlers = append(handlers.([]EventHandler), handler)
	}
}

func (bot *Bot) OnMessage(prefix string, fn EventHandlerFunc) {
	bot.On(PostTypeMessage, EventHandler{
		MessagePrefix: prefix,
		MessageType:   "",
		Handle:        fn,
	})
}

func (bot *Bot) OnPrivateMessage(prefix string, fn EventHandlerFunc) {
	bot.On(PostTypeMessage, EventHandler{
		MessagePrefix: prefix,
		MessageType:   MessageTypePrivate,
		Handle:        fn,
	})
}
func (bot *Bot) OnGroupMessage(prefix string, fn EventHandlerFunc) {
	bot.On(PostTypeMessage, EventHandler{
		MessagePrefix: prefix,
		MessageType:   MessageTypeGroup,
		Handle:        fn,
	})
}

func (bot *Bot) OnNotice(fn EventHandlerFunc) {
	bot.On(PostTypeNotice, EventHandler{
		MessagePrefix: "",
		MessageType:   "",
		Handle:        fn,
	})
}

func (bot *Bot) OnRequest(fn EventHandlerFunc) {
	bot.On(PostTypeRequest, EventHandler{
		MessagePrefix: "",
		MessageType:   "",
		Handle:        fn,
	})
}

func (bot *Bot) OnMeta(fn EventHandlerFunc) {
	bot.On(PostTypeMeta, EventHandler{
		MessagePrefix: "",
		MessageType:   "",
		Handle:        fn,
	})
}
