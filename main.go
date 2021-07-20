package main

import (
	"log"
	"net/http"

	"github.com/Korablik/Go.Websocket/internal"
	"github.com/Korablik/Go.Websocket/internal/handlers"
)

func main() {
	mux := internal.Routes()

	log.Println("Starting channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
	//_ = http.ListenAndServeTLS(":8083", mux)
}
