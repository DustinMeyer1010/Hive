package main

import (
	"log"

	"github.com/DustinMeyer1010/livechat/internal/api/websocket/chat"
	"github.com/DustinMeyer1010/livechat/internal/server"
)

func main() {

	server.Run("localhost", "8080")

	go chat.HandleMessages()

	log.Println("HTTP server started on :8080")
}
