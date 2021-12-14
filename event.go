package onebot

const (
	PostTypeMessage = "message"
	PostTypeNotice  = "notice"
	PostTypeRequest = "request"
	PostTypeMeta    = "meta_event"

	MessageTypePrivate = "private"
	MessageTypeGroup   = "group"
)

// Event 事件 https://github.com/botuniverse/onebot-11/tree/master/event
type Event struct {
	Time     int64  `json:"time" validate:"required"`
	SelfId   int64  `json:"self_id" validate:"required"`
	PostType string `json:"post_type" validate:"oneof=message notice request meta_event"`

	SubType string `json:"sub_type"`

	// message
	MessageType string           `json:"message_type"`
	MessageId   int32            `json:"message_id,"`
	UserId      int64            `json:"user_id"`
	Message     []MessageSegment `json:"message"`
	RawMessage  string           `json:"raw_message"`
	Font        int32            `json:"font"`
	Sender      Sender           `json:"sender"`

	// group
	GroupId   int64     `json:"group_id"`
	Anonymous Anonymous `json:"anonymous"`
}

type Sender struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Card     string `json:"card"`

	// group extras
	Area  string `json:"area"`
	Level string `json:"level"`
	Role  string `json:"role"`
	Title string `json:"title"`
}

type Anonymous struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
