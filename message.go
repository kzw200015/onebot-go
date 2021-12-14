package onebot

// MessageSegment 消息段 https://github.com/botuniverse/onebot-11/blob/master/message/array.md
type MessageSegment struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type messageBuilder struct {
	message []MessageSegment
}

func MessageBuilder() *messageBuilder {
	return &messageBuilder{}
}

func (m *messageBuilder) Build() []MessageSegment {
	return m.message
}

// Text https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%BA%AF%E6%96%87%E6%9C%AC
func (m *messageBuilder) Text(text string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "text",
		Data: map[string]string{
			"text": text,
		},
	})
	return m
}

// Face https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#qq-%E8%A1%A8%E6%83%85
func (m *messageBuilder) Face(id string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "face",
		Data: map[string]string{
			"id": id,
		},
	})
	return m
}

// Image https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%9B%BE%E7%89%87
func (m *messageBuilder) Image(file string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "image",
		Data: map[string]string{
			"file": file,
		},
	})
	return m
}

// Record https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E8%AF%AD%E9%9F%B3
func (m *messageBuilder) Record(file string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "record",
		Data: map[string]string{
			"file": file,
		},
	})
	return m
}

// Video https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%9F%AD%E8%A7%86%E9%A2%91
func (m *messageBuilder) Video(file string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "video",
		Data: map[string]string{
			"file": file,
		},
	})
	return m
}

// At https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E6%9F%90%E4%BA%BA
func (m *messageBuilder) At(qq string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "at",
		Data: map[string]string{
			"qq": qq,
		},
	})
	return m
}

// Reply https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%9B%9E%E5%A4%8D
func (m *messageBuilder) Reply(id string) *messageBuilder {
	m.message = append(m.message, MessageSegment{
		Type: "reply",
		Data: map[string]string{
			"id": id,
		},
	})
	return m
}
