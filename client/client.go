// Copyright (C) 2020 Raziman Mahathir

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/issadarkthing/gochat/structure"
)

// Struct to hold data about our client and to abstract the websocket
type Client struct {
	username   string
	passphrase string
	url        string
	color      string
	connection *websocket.Conn
}

// Connects client to the server
func (c *Client) connect() error {
	con, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.color = randomColor()
	c.connection = con
	return nil
}

// Send message to the server
func (c Client) send(message string) error {
	placeHolder := `
{
	"username": "%s",
	"color":    "%s",
	"message":  "%s"
}`
	data := fmt.Sprintf(placeHolder, c.username, c.color, message)
	encrypted := encrypt([]byte(data), hashKey(c.passphrase))
	return c.connection.WriteMessage(websocket.TextMessage, encrypted)
}

// Handles the incoming data and unmarshalles (if i spelt it correctly) the
// json into Message struct
// This part is very tricky, make sure the receiveHandler is a pointer receiver,
// otherwise, it cannot reflect the changes made to the struct. 
// This is due to the for loop that only receives the copy
func (c *Client) receiveHandler(handler func(data structure.Message)) {
	go func() {
		for {

			var data structure.Message
			_, msg, err := c.connection.ReadMessage()
			if err != nil {
				panic(err)
			}

			decrypted, err := decrypt(msg, hashKey(c.passphrase))
			if err != nil {
				continue
			}

			json.Unmarshal(decrypted, &data)
			handler(data)
		}
	}()
}
