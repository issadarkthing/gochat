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
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/issadarkthing/gochat/structure"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan structure.Message)
	upgrader  = websocket.Upgrader{}
)


const (
	PORT = ":8080"
)

func main() {
	
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("http server started on "+PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	clients[ws] = true

	fmt.Println("One connection created")
	for {
		var msg structure.Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			// connection closed
			// remove client
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {

		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
