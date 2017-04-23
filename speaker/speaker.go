package speaker

import (
	"github.com/ngalayko/theq_ask/fetcher"
	"github.com/ngalayko/theq_ask/reader"
	"github.com/ngalayko/theq_ask/types"
)

type Speaker interface {
	Start()
}

type speaker struct {
	fetcher fetcher.Fetcher
	reader  reader.Reader
}

func New(apiKey string) Speaker {
	queue := make(chan types.Text)

	return &speaker{
		fetcher: fetcher.New(queue),
		reader:  reader.New(apiKey, queue),
	}
}

func (t *speaker) Start() {
	go t.fetcher.FetchLoop()

	t.reader.ReadLoop()
}
