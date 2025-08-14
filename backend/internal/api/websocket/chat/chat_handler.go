package chat

import (
	"fmt"
	"net/http"

	"github.com/DustinMeyer1010/livechat/internal/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleChatConnections(w http.ResponseWriter, r *http.Request) {
	roomName := r.URL.Query().Get("room")

	if roomName == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	conn, _ := upgrader.Upgrade(w, r, nil)
	client := &types.Client{
		Conn: conn,
		Send: make(chan []byte),
		Room: roomName,
	}

	room := RoomConnection(roomName)
	room.Clients[client] = true

	fmt.Println(len(room.Clients))

	go ReadMessage(client, room)
	go WriteMessage(client)
}
