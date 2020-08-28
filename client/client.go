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
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/issadarkthing/gochat/structure"
)


type Client struct {
	username   string
	url        string
	color      string
	connection *websocket.Conn
}

func (c *Client) connect() error {
	con, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.color = randomColor()
	c.connection = con
	return nil
}

func (c Client) send(message string) error {
	placeHolder := `
{
	"username": "%s",
	"color": "%s",
	"message": "%s"
}`
	data := fmt.Sprintf(placeHolder, c.username, c.color, message)
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
