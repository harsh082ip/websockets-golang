package manager

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) ServeWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection...")

	// upgrade the http into websocket
	_, err := WebSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {

		log.Println("Error: ", err)
		return
	}

	// conn.Close()
}
