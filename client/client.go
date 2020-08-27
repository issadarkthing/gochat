package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/issadarkthing/gochat/structure"
)


type Client struct {
	username   string
	email      string
	url        string
	connection *websocket.Conn
}

func (c *Client) connect() error {
	con, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}

	c.connection = con
	return nil
}

func (c Client) send(message string) error {
	data := fmt.Sprintf(`{ "username": "%s", "email": "%s", "message": "%s" }`, 
		c.username, c.email, message)
	return c.connection.WriteMessage(websocket.TextMessage, []byte(data))
}

func (c Client) receiveHandler(handler func(data structure.Message)) {
	go func() {
		for {
			var msg structure.Message
			c.connection.ReadJSON(&msg)
			handler(msg)
		}
	}()
}
