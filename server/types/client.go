package types

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
	pongWait   = 60 * time.Second
)

type Client struct {
	Ws   *websocket.Conn
	Send chan *Message
}

func (t *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		t.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-t.Send:
			if !ok {
				t.write(websocket.CloseMessage, nil)
				return
			}
			if err := t.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := t.write(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) write(mt int, message *Message) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return c.Ws.WriteMessage(mt, data)
}

func (t *Client) SayHello(msg *Message) {
	t.write(websocket.TextMessage, msg)
}
