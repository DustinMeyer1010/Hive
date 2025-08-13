package main

import (
	"log"

	"github.com/DustinMeyer1010/livechat/internal/server"
)

func main() {

	server.Run("localhost", "5000")

	log.Println("HTTP server started on :5000")
}
