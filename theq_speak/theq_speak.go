package theq_speak

import (
	"github.com/ngalayko/theq_ask/speaker"
)

type TheqSpeak interface {
	Start()
}

type theqSpeak struct {
	speaker speaker.Speaker

	queue chan *Question
	seen  map[int64]bool
}

func New(speaker speaker.Speaker) TheqSpeak {
	return &theqSpeak{
		speaker: speaker,

		queue: make(chan *Question),
		seen:  map[int64]bool{},
	}
}

func (t *theqSpeak) Start() {
	defer func() {
		recover()
	}()

	go t.questionsLoop()

	for question := range t.queue {
		if err := t.readQuestion(question); err != nil {
			panic(err)
		}
	}
}

func (t *theqSpeak) readQuestion(question *Question) error {
	gender := speaker.Gender(question.Account.Gender)

	return t.speaker.Say(question.Title, gender)
}
