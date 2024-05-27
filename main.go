package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNum string = ":8080";

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Home Page!");
}

func Info(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Info");
}

func main() {
    log.SetFlags(0);
	log.Println("Server is running");
    
	http.HandleFunc("/",Home);
	http.HandleFunc("/info",Info);

	log.Println("Started on port", portNum);
    fmt.Println("To close connection CTRL+C :-)");

	err := http.ListenAndServe(portNum, nil)
    if err != nil {
        log.Fatal(err)
    }
}
