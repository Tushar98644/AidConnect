package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tushar98644/CoDev/internal/handlers"
	"github.com/Tushar98644/CoDev/pkg/websockets"
)

const portNum string = ":8080" 

func SetupRoutes(){
	http.HandleFunc("/",handlers.HomeHandler)
	http.HandleFunc("/info",handlers.InfoHandler)
	http.HandleFunc("/ws", websocket.WsEndpoint);
}

func Serve(){
	log.SetFlags(0);
	log.Println("Server is running");

	SetupRoutes();

	go websocket.HandleMessages();

	log.Println("Started on port", portNum);
    fmt.Println("To close connection CTRL+C :-)");

    err := http.ListenAndServe(portNum,nil)
	if err!= nil{
		log.Fatal("There was an error listening and serving to the port",err)
	}
}