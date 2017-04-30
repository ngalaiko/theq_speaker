package types

type Message struct {
	Text string `json:"text"`
}

func NewMessage(text string) Message {
	return Message{
		Text: text,
	}
}
