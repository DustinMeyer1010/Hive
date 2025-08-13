package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn  *websocket.Conn
	group string
}

var (
	clients   = make(map[*websocket.Conn]string) // conn => group
	broadcast = make(chan Message)               // channel for incoming messages
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // allow all origins
	}
	groups = make(map[string]map[*websocket.Conn]bool) // group => connections
)

// Message defines the structure of messages
type Message struct {
	Username string `json:"username"`
	Group    string `json:"group"` // group/chat room ID
	Message  string `json:"message"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	// First message from client should specify the group to join
	var joinMsg Message
	err = ws.ReadJSON(&joinMsg)
	if err != nil {
		log.Println("Failed to read join message:", err)
		return
	}

	group := joinMsg.Group
	clients[ws] = group

	// Add client to group map
	if groups[group] == nil {
		groups[group] = make(map[*websocket.Conn]bool)
	}
	groups[group][ws] = true

	log.Printf("Client joined group %s\n", group)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// Set group to client’s group for safety
		msg.Group = group
		broadcast <- msg
	}

	// Cleanup on disconnect
	delete(clients, ws)
	delete(groups[group], ws)
}

func HandleMessages() {
	for {
		msg := <-broadcast
		conns := groups[msg.Group]

		for client := range conns {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
				delete(groups[msg.Group], client)
			}
		}
	}
}
