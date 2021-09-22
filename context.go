package onebot

type Context struct {
	Bot *Bot
	Event
}

func (ctx *Context) Reply() error {
	return nil
}
