package utils

import (
	"github.com/DustinMeyer1010/livechat/internal/db"
	"github.com/DustinMeyer1010/livechat/internal/types"
	"github.com/gorilla/websocket"
)

func ReadMessage(c *types.Client, room *types.Room) {
	defer func() {
		room.Clients[c] = false
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		room.Broadcast <- msg
		db.SaveMessage(msg)
	}

}

func WriteMessage(c *types.Client) {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
