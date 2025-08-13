package utils

import (
	"sync"

	"github.com/DustinMeyer1010/livechat/internal/types"
)

var (
	rooms     = make(map[string]*types.Room) // room name -> room
	roomsLock sync.Mutex
)

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
