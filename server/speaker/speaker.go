package speaker

import (
	"github.com/ngalayko/theq_speaker/server/converter"
	"github.com/ngalayko/theq_speaker/server/fetcher"
	"github.com/ngalayko/theq_speaker/server/sender"
	"github.com/ngalayko/theq_speaker/server/types"
	"golang.org/x/net/websocket"
)

type Speaker interface {
	Start(ws *websocket.Conn)
}

type speaker struct {
	fetcher   fetcher.Fetcher
	converter converter.Converter
	sender    sender.Sender
}

func New(apiKey string) Speaker {
	queueToConvert := make(chan types.Text)
	queueToSend := make(chan types.Message)

	return &speaker{
		fetcher:   fetcher.New(queueToConvert),
		converter: converter.New(apiKey, queueToConvert, queueToSend),
		sender:    sender.New(queueToSend),
	}
}

func (t *speaker) Start(ws *websocket.Conn) {
	go t.fetcher.FetchLoop()
	go t.converter.ConvertLoop()

	t.sender.SendLoop(ws)
}
