package sender

import (
	"github.com/ngalayko/theq_speaker/server/logger"
	"github.com/ngalayko/theq_speaker/server/types"
	"time"
)

type Sender interface {
	SendLoop()

	Register(client *types.Client)
	Unregister(client *types.Client)
}

type sender struct {
	logger logger.Logger

	clients    map[*types.Client]bool
	register   chan *types.Client
	unregister chan *types.Client

	queueToSend chan *types.Message

	timeout time.Duration
}

func New(queueToSend chan *types.Message, timeout time.Duration) Sender {
	return &sender{
		logger: logger.New("sender"),

		clients:    map[*types.Client]bool{},
		register:   make(chan *types.Client),
		unregister: make(chan *types.Client),

		queueToSend: queueToSend,

		timeout: timeout,
	}
}

func (t *sender) SendLoop() {
	ticker := time.NewTicker(t.timeout)

	for {
		select {
		case client := <-t.register:
			t.addClient(client)

		case client := <-t.unregister:
			_, ok := t.clients[client]
			if ok {
				t.deleteClient(client)
			}

		case <-ticker.C:
			message := <-t.queueToSend
			t.broadcast(message)
		}
	}
}

func (t *sender) Register(client *types.Client) {
	t.register <- client
}

func (t *sender) Unregister(client *types.Client) {
	t.unregister <- client
}

func (t *sender) broadcast(message *types.Message) {
	success := 0
	fail := 0
	for client := range t.clients {
		select {
		case client.Send <- message:
			success++
		default:
			fail++
			t.deleteClient(client)
		}
	}

	t.logger.Info("Mesage sended", logger.Fields{
		"message": message.Text,
		"success": success,
		"fail":    fail,
	})
}

func (t *sender) addClient(client *types.Client) {
	t.clients[client] = true

	t.logger.Info("Client registred", logger.Fields{
		"clients": len(t.clients),
	})
}

func (t *sender) deleteClient(client *types.Client) {
	delete(t.clients, client)
	close(client.Send)

	t.logger.Info("Client unregistred", logger.Fields{
		"clients": len(t.clients),
	})
}
