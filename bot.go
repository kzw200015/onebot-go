package onebot

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Bot 机器人
type Bot struct {
	client     *wsClient
	config     Config
	respTemp   sync.Map
	logger     *logrus.Logger
	handlerMap struct {
		privateMessage EventHandlerMap
		groupMessage   EventHandlerMap
		notice         EventHandlerMap
		request        EventHandlerMap
		meta           EventHandlerMap
	}
}

// Config 配置
type Config struct {
	SelfId           int64
	AdminIds         []int64
	URL              string
	Token            string
	Logger           *logrus.Logger
	ReceiveHeartBeat bool
}

// NewBot 构造函数，返回Bot实例指针
func NewBot(config Config) *Bot {
	return &Bot{
		config: config,
	}
}

// Run 启动机器人
func (bot *Bot) Run() {
	bot.setLogger(bot.config.Logger)

	err := bot.connect()
	if err != nil {
		bot.logger.Panic(err)
	}
	// 监听系统中断信号，断开连接
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	updates := bot.getUpdateChan()
	for {
		select {
		case update := <-updates:
			updateResult := gjson.ParseBytes(update)

			if updateResult.Get("post_type").Exists() { // 响应类型为事件
				var e Event
				err := json.Unmarshal(update, &e)
				if err != nil {
					bot.logger.Error(err)
				}
				bot.logger.Debugf("收到事件: %+v", e)
				go bot.handleEvent(e)

			} else if echo := updateResult.Get("echo"); echo.Exists() { // 响应类型为API调用
				respChan, ok := bot.respTemp.LoadAndDelete(echo.String())

				if ok {
					select {
					case <-respChan.(chan APIResp):
					default:
						var resp APIResp
						err := json.Unmarshal(update, &resp)
						if err != nil {
							bot.logger.Error(err)
						}
						bot.logger.Debugf("收到API请求响应: %+v", resp)
						// 响应传入 channel
						respChan.(chan APIResp) <- resp
					}
				}
			}
		case <-interrupt:
			bot.close()
			return
		}
	}
}

func (bot *Bot) setLogger(logger *logrus.Logger) {
	bot.logger = logger
}

// connect 连接服务器
func (bot *Bot) connect() error {
	bot.logger.Info("正在连接")
	client, err := newClient(
		bot.config.URL,
		http.Header{"Authorization": []string{"Bearer " + bot.config.Token}})

	if err != nil {
		return err
	}

	bot.client = client
	bot.logger.Info("已连接到:" + bot.config.URL)
	return nil
}

// read 读取数据
func (bot *Bot) read() (int, []byte) {
	t, message, err := bot.client.conn.ReadMessage()
	if err != nil {
		bot.logger.Warnf("Websocket读取失败: %v", err)
	}

	return t, message
}

// getUpdateChan 获取更新channel
func (bot *Bot) getUpdateChan() <-chan []byte {
	messageChan := make(chan []byte)
	go func() {
		for {
			t, message := bot.read()
			if t == websocket.TextMessage {
				messageChan <- message
			}
		}
	}()

	return messageChan
}

// close 断开服务器连接
func (bot *Bot) close() error {
	bot.logger.Info("正在断开")
	return bot.client.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
