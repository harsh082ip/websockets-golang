package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	// cl "github.com/harsh082ip/websockets-golang/client"
	// "github.com/harsh082ip/websockets-golang/client"
)

var (
	WebSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Manager struct {
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) ServeWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection...")

	// upgrade the http into websocket
	conn, err := WebSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {

		log.Println("Error: ", err)
		return
	}

	// when we have a new connection, we'll make a new client
	client := NewClient(conn, m)

	// now add this client
	m.addClient(client)

	// start process
	go client.readMessages()
	go client.writeMessage()
}

func (m *Manager) addClient(client *Client) {

	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}
