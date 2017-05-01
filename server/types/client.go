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

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, nil)
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, nil); err != nil {
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
