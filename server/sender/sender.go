package sender

import (
	"github.com/ngalayko/theq_speaker/server/logger"
	"github.com/ngalayko/theq_speaker/server/types"
	"golang.org/x/net/websocket"
	"os"
)

type Sender interface {
	SendLoop(ws *websocket.Conn)
}

type sender struct {
	logger logger.Logger

	queueToSend chan types.Message
}

func New(queueToSend chan types.Message) Sender {
	return &sender{
		logger: logger.New("sender"),

		queueToSend: queueToSend,
	}
}

func (t *sender) SendLoop(ws *websocket.Conn) {
	defer func() {
		recover()
	}()

	for message := range t.queueToSend {
		if err := t.sendMessage(ws, message); err != nil {
			t.logger.Info("SendLoop error", logger.Fields{
				"error": err,
			})

			panic(err)
		}
	}
}

func (t *sender) sendMessage(ws *websocket.Conn, message types.Message) error {
	if err := websocket.JSON.Send(ws, message); err != nil {
		return err
	}

	t.logger.Info("Message sended", logger.Fields{
		"message": message,
	})

	return os.Remove(message.Text)
}
