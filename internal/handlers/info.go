package handlers

import (
	"fmt"
	"net/http"
)

func InfoHandler (w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Info handler");
}