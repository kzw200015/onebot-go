# onebot-go

onebot协议机器人库（开发中）

### 目前已实现

#### API

- [x] `send_private_msg`
- [x] `send_group_msg`

#### 消息段

- [x] `Text`
- [x] `Face`
- [x] `Image`
- [x] `Record`
- [x] `Video`
- [x] `At`
- [x] `Reply`

### 事件

- [x] `message`

### 快速开始

```go
package main

import (
	"errors"
	"fmt"

	"github.com/kzw200015/onebot-go"
	"github.com/sirupsen/logrus"
)

func main() {
	bot := onebot.New(onebot.BotConfig{
		SelfId: 12345678,
		ApiConfig: onebot.ApiConfig{
			Token:   "token",
			Address: "http://localhost:5700",
		},
		ServerConfig: onebot.ServerConfig{
			Secret:  "secret",
			Address: ":8080",
			Path:    "/event",
		},
	})

	bot.OnPrivateMessage("echo", func(bot *onebot.Bot, event onebot.Event) {
		bot.SendPrivateMsg(event.UserId, event.RawMessage, false)
	})

	bot.Start()
}
```