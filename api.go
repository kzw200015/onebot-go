package onebot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"
)

const MimeType = "application/json"

// APIResp API响应 https://github.com/botuniverse/onebot-11/blob/master/communication/http.md#%E5%93%8D%E5%BA%94
type APIResp struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    interface{} `json:"data"`
}

// SendPrivateMsg 发送私聊消息 https://github.com/botuniverse/onebot-11/blob/master/api/public.md#send_private_msg-%E5%8F%91%E9%80%81%E7%A7%81%E8%81%8A%E6%B6%88%E6%81%AF
func (bot *Bot) SendPrivateMsg(userId int64, msg interface{}, autoEscape bool) (APIResp, error) {
	return bot.Request("send_private_msg", map[string]interface{}{
		"user_id":     userId,
		"message":     msg,
		"auto_escape": autoEscape,
	})
}

// SendGroupMsg 发送群消息 https://github.com/botuniverse/onebot-11/blob/master/api/public.md#send_group_msg-%E5%8F%91%E9%80%81%E7%BE%A4%E6%B6%88%E6%81%AF
func (bot *Bot) SendGroupMsg(groupId int64, msg interface{}, autoEscape bool) (APIResp, error) {
	return bot.Request("send_group_msg", map[string]interface{}{
		"group_id":    groupId,
		"message":     msg,
		"auto_escape": autoEscape,
	})
}

func (bot *Bot) Request(action string, requestBody interface{}) (APIResp, error) {
	bs, err := json.Marshal(requestBody)
	if err != nil {
		return APIResp{}, err
	}

	resp, err := http.Post(path.Join(bot.config.HttpConfig.RemoteApiAddr, action),
		MimeType,
		bytes.NewReader(bs))
	if err != nil {
		return APIResp{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResp{}, err
	}

	var apiResp APIResp
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return APIResp{}, err
	}

	return apiResp, nil
}
