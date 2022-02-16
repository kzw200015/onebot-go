# onebot-go

onebot协议机器人库（已停止开发）

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

服务端实现推荐使用 `go-cqhttp`，使用HTTP/HTTP POST方式连接

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
		SelfId: 12345678, //QQ号
		ApiConfig: onebot.ApiConfig{ //与HTTP API配置要一致
			Token:   "token", //API token
			Address: "http://localhost:5700", //API地址
		},
		ServerConfig: onebot.ServerConfig{ //与HTTP POST配置要一致
			Secret:  "secret", //验证用secret
			Address: ":8080", //监听地址
			Path:    "/event", //监听Path
		},
	})

	bot.OnPrivateMessage("", func(bot *onebot.Bot, event onebot.Event) {
		bot.SendPrivateMsg(event.UserId, onebot.MessageBuilder().Text("hello").Text("world").Build()...)
	})

	bot.Start()
}
```
