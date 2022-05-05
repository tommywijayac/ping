package main

import (
	"log"
	"net/http"

	hHttp "github.com/tommywijayac/ping/internal/handler/http"
)

const server string = "localhost:3000"

func initServer(handler *hHttp.Handler) {
	http.HandleFunc("/ping", handler.HandlerClientWebsocket)
	http.HandleFunc("/dev", handler.HandlerDevPush)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//start in separate goroutine so we can do graceful shutdown
	go http.ListenAndServe(server, nil)
	log.Printf("main: server starting at %s\n", server)
}
