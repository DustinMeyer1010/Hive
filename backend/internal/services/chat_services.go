package services

import (
	"sync"

	"github.com/DustinMeyer1010/livechat/internal/db"
	"github.com/DustinMeyer1010/livechat/internal/types"
	"github.com/gorilla/websocket"
)

var (
	rooms     = make(map[string]*types.Room) // room name -> room
	roomsLock sync.Mutex
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

func RoomConnection(name string) *types.Room {
	roomsLock.Lock()
	defer roomsLock.Unlock()

	var room *types.Room
	var exist bool

	if room, exist = rooms[name]; exist {
		return room
	}

	room = &types.Room{
		Clients:   make(map[*types.Client]bool),
		Broadcast: make(chan []byte),
	}

	rooms[name] = room
	go room.Run()

	return room
}
