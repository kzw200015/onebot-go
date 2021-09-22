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

	qqBot.OnPrivateMessage(func(ctx *onebot.Context) error {
		resp, err := ctx.Bot.SendPrivateMessage(ctx.UserId, onebot.Text(ctx.RawMessage))
		if err != nil {
			return err
		}

		if resp.IsOK() {
			return nil
		}
		return errors.New("something wrong")
	})

	// 使用中间件
	qqBot.OnMessageWithPrefix("", func(ctx *onebot.Context) error {
		resp, err := ctx.Bot.SendPrivateMessage(ctx.UserId, onebot.Text(ctx.RawMessage))
		if err != nil {
			return err
		}

		if resp.IsOK() {
			return nil
		}
		return errors.New("something wrong")
	}, func(next onebot.EventHandlerFunc) onebot.EventHandlerFunc {
		return func(ctx *onebot.Context) error {
			fmt.Println("before handler")
			if err := next(ctx); err != nil {
				// process error
			}
			fmt.Println("after handler")
			return nil
		}
	})

	qqBot.Run()
}
```