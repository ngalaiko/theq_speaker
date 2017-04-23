package reader

import (
	"github.com/ngalayko/theq_ask/speaker"
)

type Reader interface {
	Read()
}

type reader struct {
	speaker speaker.Speaker

	queue chan *Question
	seen  map[int64]bool
}

func New(speaker speaker.Speaker) Reader {
	return &reader{
		speaker: speaker,

		queue: make(chan *Question),
		seen:  map[int64]bool{},
	}
}

func (t *reader) Read() {
	defer func() {
		recover()
	}()

	go t.fetchLoop()

	for question := range t.queue {
		if err := t.readQuestion(question); err != nil {
			panic(err)
		}
	}
}

func (t *reader) readQuestion(question *Question) error {
	gender := speaker.Gender(question.Account.Gender)

	return t.speaker.Say(question.Title, gender)
}
