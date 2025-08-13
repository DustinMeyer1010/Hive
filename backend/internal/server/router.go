package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/DustinMeyer1010/livechat/internal/api/websocket/chat"
)

//go:embed dist/*
var frontend embed.FS

//go:embed dist/index.html
var indexHTML []byte

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", reactHandler)
	mux.HandleFunc("/ws", chat.HandleChatConnections)

	return mux
}

func reactHandler(w http.ResponseWriter, r *http.Request) {

	distFS, err := fs.Sub(frontend, "dist")

	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(distFS))
	requestPath := path.Clean(r.URL.Path[1:])

	if requestPath == "/" {
		w.Write(indexHTML)
		return
	}

	file, err := distFS.Open(requestPath)

	if err != nil {
		w.Write(indexHTML)
	}

	file.Close()

	fileServer.ServeHTTP(w, r)

}
