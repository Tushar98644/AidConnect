package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

const portNum string = ":8080";

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn){	
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error Upgrading Connection" , err);
	}

	log.Println("Client Connected");

	err = ws.WriteMessage(1, []byte("Hi Client!"));
	if err != nil {
		log.Println("There was an error writing the message",err)
	}

	reader(ws);
}

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Home Page!");
}

func Info(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Info");
}



func setupRoutes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)
	http.HandleFunc("/ws", wsEndpoint)
}	


func main() {
    log.SetFlags(0);
	log.Println("Server is running");
    
	setupRoutes();

	log.Println("Started on port", portNum);
    fmt.Println("To close connection CTRL+C :-)");

	err := http.ListenAndServe(portNum, nil)
    if err != nil {
        log.Fatal(err)
    }
}
