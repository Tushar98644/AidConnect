package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error Upgrading Connection", err)
        return
    }
    
    AddClient(ws)
    log.Println("Client Connected")

    for {
        var msg Message
        // Read in a new message as JSON and map it to a Message object
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            RemoveClient(ws)
            break
        }
        // Send the new message to the broadcast channel
        BroadcastMessage(msg)
    }
}



