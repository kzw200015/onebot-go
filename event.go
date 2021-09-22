package onebot

const (
	MessageEvent = "message"
	NoticeEvent  = "notice"
	RequestEvent = "Request"
	MetaEvent    = "meta_event"
)

// Event 事件 https://github.com/botuniverse/onebot/tree/master/v11/specs/event
type Event struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`

	// Private
	MessageType string           `json:"message_type"`
	SubType     string           `json:"sub_type"`
	MessageId   int32            `json:"message_id,"`
	UserId      int64            `json:"user_id"`
	Message     []MessageSegment `json:"message"`
	RawMessage  string           `json:"raw_message"`
	Font        int32            `json:"font"`

	// Group
	GroupId   int64     `json:"group_id"`
	Anonymous Anonymous `json:"anonymous"`
	Sender    Sender    `json:"sender"`

	// Notice
	File File `json:"file"`

	//Meta
	MetaEventType string `json:"meta_event_type"`
}

type File struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}

type Sender struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`

	// Group
	Card  string `json:"card"`
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
