package onebot

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Bot struct {
	config     *BotConfig
	httpServer *http.ServeMux
	Logger     *logrus.Logger
	handlerMap *HandlerMap
}

type BotConfig struct {
	SelfId     int64
	Secret     string
	HttpConfig HttpConfig
}

type HttpConfig struct {
	RemoteApiAddr string
	Port          int
	Path          string
}

func NewBot(config BotConfig) *Bot {
	return &Bot{
		config:     &config,
		httpServer: http.NewServeMux(),
		Logger:     DefaultLogger(logrus.DebugLevel),
	}
}

func (bot *Bot) Start() {
	bot.init()
	http.ListenAndServe(":"+strconv.Itoa(bot.config.HttpConfig.Port), bot.httpServer)
}

func (bot *Bot) init() {
	bot.httpServer.HandleFunc(bot.config.HttpConfig.Path, bot.eventDispatcher)
}

func (bot *Bot) eventDispatcher(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		bot.Logger.Error("读取 Body 错误: " + err.Error())
		return
	}

	if receivedSig := strings.TrimPrefix(r.Header.Get("X-Signature"), "sha1="); checkSignature(receivedSig, body, bot.config.Secret) {
		bot.Logger.Warnln("签名认证失败")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	selfId, _ := strconv.ParseInt(r.Header.Get("X-Self-ID"), 10, 64)

	parsedBody := gjson.ParseBytes(body)
	if bot.config.SelfId != selfId {
		bot.Logger.Warnln("未知 self_id 事件上报", selfId)
		return
	}

	postType := parsedBody.Get("post_type").String()
	switch postType {
	case PostTypeMessage:
		bot.dispatch(body, postType)
	case PostTypeNotice:
		bot.dispatch(body, postType)
	case PostTypeRequest:
		bot.dispatch(body, postType)
	case PostTypeMeta:
		bot.dispatch(body, postType)
	}
}

func (bot *Bot) dispatch(body []byte, postType string) {
	var e Event
	if err := json.Unmarshal(body, &e); err != nil {
		bot.Logger.Error("解析 MessageEvent 错误: " + err.Error())
		return
	}
	DefaultEventValidator().Validate(&e)
	handlers := bot.handlerMap.GetHandlers(postType)
	for _, handler := range handlers {
		go handler.Handle(bot, e)
	}
}

func checkSignature(sha string, body []byte, secret string) bool {
	hash := hmac.New(sha1.New, []byte(secret))
	hash.Write(body)
	expectedSig := hex.EncodeToString(hash.Sum(nil))
	return sha == expectedSig
}
