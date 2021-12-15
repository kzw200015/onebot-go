package onebot

import "errors"

type Context struct {
	Event Event
	Bot   *Bot
}

func (ctx *Context) SendMessage(message ...MessageSegment) (APIResp, error) {
	switch ctx.Event.MessageType {
	case MessageTypePrivate:
		return ctx.Bot.SendPrivateMsg(ctx.Event.UserId, message...)
	case MessageTypeGroup:
		return ctx.Bot.SendGroupMsg(ctx.Event.GroupId, message...)
	default:
		return APIResp{}, errors.New("message type not supported")
	}
}
