package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	// to avoid concurrent writes
	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {

	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

func (c *Client) readMessages() {

	for {

		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Error Reading Messages, %v", err)

			}
			break
		}

		for wsclient := range c.manager.clients {
			wsclient.egress <- payload
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}

func (c *Client) writeMessage() {

	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection Closed: ", err)
				}
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Failed to Send Message: ", err)
			}
			log.Println("Message Sent")
		}
	}
}
