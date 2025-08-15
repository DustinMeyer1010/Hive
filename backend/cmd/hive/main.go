package main

import (
	"fmt"
	"log"

	"github.com/DustinMeyer1010/livechat/internal/db"
	"github.com/DustinMeyer1010/livechat/internal/server"
)

func main() {

	db.Init()
	fmt.Println(db.ReadAllChat())
	server.Run("localhost", "5000")

	log.Println("HTTP server started on :5000")
}
