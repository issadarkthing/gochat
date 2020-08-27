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
