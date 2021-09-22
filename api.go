package main

import (
	"github.com/google/uuid"
	"time"
)

// APIReq API请求 https://github.com/botuniverse/onebot/blob/master/v11/specs/communication/ws.md
type APIReq struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

// APIResp API响应 https://github.com/botuniverse/onebot/blob/master/v11/specs/communication/ws.md
type APIResp struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    interface{} `json:"data"`
	Echo    string      `json:"echo"`
}

func (r APIResp) IsOK() bool {
	return r.Status == "ok"
}

// SendPrivateMessage 发送私聊消息 https://github.com/botuniverse/onebot/blob/master/v11/specs/api/public.md#send_private_msg-%E5%8F%91%E9%80%81%E7%A7%81%E8%81%8A%E6%B6%88%E6%81%AF
func (bot *Bot) SendPrivateMessage(userId int64, message ...MessageSegment) (APIResp, error) {
	bot.logger.Infof("发送私聊消息至 %v :%v", userId, message)
	params := map[string]interface{}{
		"user_id": userId,
		"message": message,
	}
	return bot.Request(APIReq{
		Action: "send_private_msg",
		Params: params,
	})
}

// SendGroupMessage 发送群消息 https://github.com/botuniverse/onebot/blob/master/v11/specs/api/public.md#send_group_msg-%E5%8F%91%E9%80%81%E7%BE%A4%E6%B6%88%E6%81%AF
func (bot *Bot) SendGroupMessage(groupId int64, message ...MessageSegment) (APIResp, error) {
	bot.logger.Infof("发送群消息至 %v :%v", groupId, message)
	params := map[string]interface{}{
		"group_id": groupId,
		"message":  message,
	}
	return bot.Request(APIReq{
		Action: "send_group_msg",
		Params: params,
	})
}

// Request 发送请求，一般情况不应直接使用此方法
func (bot *Bot) Request(req APIReq) (APIResp, error) {
	echo := uuid.NewString()
	req.Echo = echo
	respChan := make(chan APIResp)
	bot.respTemp.Store(echo, respChan)

	// 加锁以并发安全
	bot.client.mu.Lock()
	err := bot.client.conn.WriteJSON(req)
	bot.client.mu.Unlock()
	if err != nil {
		return APIResp{}, err
	}

	var resp APIResp

	select {
	case resp = <-respChan:
		close(respChan)
	case <-time.After(10 * time.Second):
		close(respChan) //关闭channel，通知本次请求已结束
		bot.logger.Warnf("API请求超时: %+v", req)
	}

	return resp, nil
}
