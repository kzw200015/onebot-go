package onebot

import "errors"

type Context struct {
	Bot *Bot
	Event
}

func (ctx *Context) Send(message ...MessageSegment) (APIResp, error) {
	switch ctx.PostType {
	case MessageEvent:
		switch ctx.MessageType {
		case PrivateMessage:
			return ctx.Bot.SendPrivateMessage(ctx.UserId, message...)
		case GroupMessage:
			return ctx.Bot.SendGroupMessage(ctx.GroupId, message...)
		}
	}

	return APIResp{}, errors.New("不支持的事件类型，仅支持直接回复消息事件")
}
