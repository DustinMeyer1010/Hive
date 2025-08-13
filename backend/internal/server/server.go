package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

var server *http.Server

func Run(address string, port string) {
	router := createRouter()

	server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", address, port),
		Handler: router,
	}

	fmt.Printf("Started server at \nhttp://%s:%s\n", address, port)
	log.Fatal(server.ListenAndServe())
}

func Shutdown(ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	fmt.Println("Server has been shutdown")
}
