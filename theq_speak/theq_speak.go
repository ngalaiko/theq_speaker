package theq_speak

import (
	"../speaker"
)

type TheqAsk interface {
	Start()
}

type theqAsk struct {
	speaker speaker.Speaker

	queue chan *Question
	seen  map[int64]bool
}

func New(speaker speaker.Speaker) TheqAsk {
	return &theqAsk{
		speaker: speaker,

		queue: make(chan *Question),
		seen:  map[int64]bool{},
	}
}

func (t *theqAsk) Start() {
	go t.questionsLoop()

	for {
		select {
		case question := <-t.queue:
			if err := t.readQuestion(question); err != nil {
				panic(err)
			}
		}
	}
}

func (t *theqAsk) readQuestion(question *Question) error {
	gender := speaker.Gender(question.Account.Gender)

	return t.speaker.Say(question.Title, gender)
}
