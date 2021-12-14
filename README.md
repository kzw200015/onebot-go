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
	qqBot := onebot.NewBot(onebot.Config{
		SelfId:   12345678,
		AdminIds: []int64{12345678},
		URL:      "wss://example.com/ws", //目前支持正向WS连接
		Token:    "password",
		Logger:   onebot.DefaultLogger(logrus.DebugLevel),
	})

	qqBot.OnPrivateMessage("echo",func(bot *Bot, event Event){
		bot.SendPrivateMsg(event.UserId, event.RawMessage, false)
	})

	qqBot.Start()
}
```