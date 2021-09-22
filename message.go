package onebot

const (
	PrivateMessage = "private"
	GroupMessage   = "group"
)

// MessageSegment 消息段 https://github.com/botuniverse/onebot/blob/master/v11/specs/message/array.md
type MessageSegment struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

// Text https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%BA%AF%E6%96%87%E6%9C%AC
func Text(text string) MessageSegment {
	return MessageSegment{
		Type: "text",
		Data: map[string]string{
			"text": text,
		},
	}
}

// Face https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#qq-%E8%A1%A8%E6%83%85
func Face(id string) MessageSegment {
	return MessageSegment{
		Type: "face",
		Data: map[string]string{
			"id": id,
		},
	}
}

// Image https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%9B%BE%E7%89%87
func Image(file string) MessageSegment {
	return MessageSegment{
		Type: "image",
		Data: map[string]string{
			"file": file,
		},
	}
}
