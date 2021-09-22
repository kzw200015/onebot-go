package main

type Context struct {
	bot *Bot
	Event
}

func (ctx *Context) Reply() error {
	return nil
}
