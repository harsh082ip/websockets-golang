package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	WEBPORT = ":8084"
)

func main() {
	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", manager.ServeWs)

	fmt.Println("Server Started on ", WEBPORT)
	log.Fatal(http.ListenAndServe(WEBPORT, nil))
}
