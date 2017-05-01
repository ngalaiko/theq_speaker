package speaker

import (
	"github.com/gorilla/websocket"
	"github.com/ngalayko/theq_speaker/server/converter"
	"github.com/ngalayko/theq_speaker/server/fetcher"
	"github.com/ngalayko/theq_speaker/server/logger"
	"github.com/ngalayko/theq_speaker/server/sender"
	"github.com/ngalayko/theq_speaker/server/types"
	"net/http"
	"time"
)

const (
	maxMessageSize = 1024 * 1024
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  maxMessageSize,
		WriteBufferSize: maxMessageSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Speaker interface {
	Start()
	ServeWs(w http.ResponseWriter, r *http.Request)
}

type Config struct {
	ApiKey      string        `yaml:"ApiKey"`
	SendTimeout time.Duration `yaml:"SendTimeout"`
	Listen      string        `yaml:"Listen"`

	HelloBase64 string
}

type speaker struct {
	logger logger.Logger

	fetcher   fetcher.Fetcher
	converter converter.Converter
	sender    sender.Sender
}

func New(config Config) Speaker {
	queueToConvert := make(chan types.Text)
	queueToSend := make(chan *types.Message)

	return &speaker{
		logger: logger.New("speaker"),

		fetcher:   fetcher.New(queueToConvert),
		converter: converter.New(config.ApiKey, queueToConvert, queueToSend),
		sender:    sender.New(queueToSend, config.SendTimeout, config.HelloBase64),
	}
}

func (t *speaker) Start() {
	go t.fetcher.FetchLoop()
	go t.converter.ConvertLoop()
	go t.sender.SendLoop()
}

func (t *speaker) ServeWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		t.logger.Error("Upgrade error", logger.Fields{
			"error": err,
		})
		return
	}

	client := &types.Client{
		Send: make(chan *types.Message, maxMessageSize),
		Ws:   ws,
	}

	t.sender.Register(client)

	defer func() {
		t.sender.Unregister(client)
		client.Ws.Close()
	}()

	client.WritePump()
}
