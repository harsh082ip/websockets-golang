package main

import (
	"log"
	"net/http"

	manager "github.com/harsh082ip/websockets-golang/Manager"
)

const (
	WEBPORT = ":8084"
)

func main() {
	manager := manager.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", manager.ServeWs)

	log.Fatal(http.ListenAndServe(WEBPORT, nil))
}
