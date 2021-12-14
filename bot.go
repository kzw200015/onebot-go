package onebot

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Bot struct {
	BotConfig  *BotConfig
	Logger     *logrus.Logger
	httpServer *http.ServeMux
	handlerMap *HandlerMap
}

type BotConfig struct {
	SelfId       int64
	ApiConfig    ApiConfig
	ServerConfig ServerConfig
}

type ApiConfig struct {
	Token   string
	Address string
}

type ServerConfig struct {
	Secret  string
	Address string
	Path    string
}

func New(botConfig BotConfig) *Bot {
	return &Bot{
		BotConfig:  &botConfig,
		Logger:     DefaultLogger(logrus.DebugLevel),
		httpServer: http.NewServeMux(),
		handlerMap: &HandlerMap{
			Map: &sync.Map{},
		},
	}
}

func (bot *Bot) Start() {
	bot.init()
	bot.Logger.Infoln("Bot tarted")
	bot.Logger.Infoln("Listening event on " + path.Join(bot.BotConfig.ServerConfig.Address, bot.BotConfig.ServerConfig.Path))
	http.ListenAndServe(bot.BotConfig.ServerConfig.Address, bot.httpServer)
}

func (bot *Bot) init() {
	bot.httpServer.HandleFunc(bot.BotConfig.ServerConfig.Path, bot.eventDispatcher)
}

func (bot *Bot) eventDispatcher(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		bot.Logger.Error("读取 Body 错误: " + err.Error())
		return
	}

	sig := r.Header.Get("X-Signature")
	if sig != "" {
		if receivedSig := strings.TrimPrefix(r.Header.Get("X-Signature"), "sha1="); !checkSignature(receivedSig, body, bot.BotConfig.ServerConfig.Secret) {
			bot.Logger.Warnln("签名认证失败")
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)

	selfId, _ := strconv.ParseInt(r.Header.Get("X-Self-ID"), 10, 64)

	if bot.BotConfig.SelfId != selfId {
		bot.Logger.Warnln("未知 self_id 事件上报", selfId)
		return
	}

	postType := gjson.ParseBytes(body).Get("post_type").String()
	bot.dispatch(body, postType)
}

func (bot *Bot) dispatch(body []byte, postType string) {
	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		bot.Logger.Errorln("解析 MessageEvent 错误: " + err.Error())
		return
	}

	if err := DefaultEventValidator().Validate(&event); err != nil {
		bot.Logger.Errorln("解析 MessageEvent 错误: " + err.Error())
		return
	}

	handlers := bot.handlerMap.GetHandlers(postType)
	for _, handler := range handlers {

		if event.PostType == PostTypeMessage {
			if !strings.HasPrefix(event.RawMessage, handler.MessagePrefix) {
				continue
			}

			if handler.MessageType != "" && handler.MessageType != event.MessageType {
				continue
			}
		}

		go handler.Handle(bot, event)
	}
}

func checkSignature(sha string, body []byte, secret string) bool {
	hash := hmac.New(sha1.New, []byte(secret))
	hash.Write(body)
	expectedSig := hex.EncodeToString(hash.Sum(nil))
	return sha == expectedSig
}
