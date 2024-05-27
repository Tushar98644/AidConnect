package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func HandleMessages() {
    for {
        // Grab the next message from the broadcast channel
        msg := <-broadcast

        // Send it out to every client that is currently connected
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func AddClient(client *websocket.Conn) {
    clients[client] = true
}

func RemoveClient(client *websocket.Conn) {
    delete(clients, client)
}

func BroadcastMessage(msg Message) {
    broadcast <- msg
}