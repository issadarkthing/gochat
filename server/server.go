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
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
	upgrader  = websocket.Upgrader{}
)



func main() {

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "8080"
	}
	
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("http server started on "+PORT)
	err := http.ListenAndServe("localhost:"+PORT, nil)
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

	log.Println("One connection created")
	log.Printf("Total number of connections %d\n", len(clients))
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// connection closed
			// remove client
			delete(clients, ws)
			log.Println("One connection closed")
			log.Printf("Total number of connections %d\n", len(clients))
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				delete(clients, client)
			}
		}
	}
}
