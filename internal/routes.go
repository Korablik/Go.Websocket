package internal

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/Korablik/Go.Websocket/internal/handlers"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
