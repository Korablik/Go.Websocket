package main

import (
	"log"
	"net/http"

	"github.com/Korablik/Go.Websocket/internal"
)

func main() {
	mux := internal.Routes()

	log.Println("Starting web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
