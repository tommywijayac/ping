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

// //TODO: temporary endpoint to test component refresh
// http.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
// 	mode := r.URL.Query().Get("mode")

// 	type foo struct {
// 		ID    int    `json:"id"`
// 		Title string `json:"title"`
// 		State string `json:"state"`
// 	}

// 	data := []foo{}
// 	if mode == "1" {
// 		data = []foo{
// 			{
// 				ID:    1,
// 				Title: fmt.Sprintf("lazule"),
// 			},
// 			{
// 				ID:    500,
// 				Title: fmt.Sprintf("sapphire"),
// 			},
// 			{
// 				ID:    998,
// 				Title: fmt.Sprintf("ruby"),
// 			},
// 			{
// 				ID:    999,
// 				Title: fmt.Sprintf("emerald"),
// 			},
// 		}
// 	} else if mode == "2" {
// 		data = []foo{
// 			{
// 				ID:    1,
// 				Title: fmt.Sprintf("lazule"),
// 				State: "warning",
// 			},
// 			{
// 				ID:    500,
// 				Title: fmt.Sprintf("sapphire"),
// 				State: "active",
// 			},
// 			{
// 				ID:    998,
// 				Title: fmt.Sprintf("ruby"),
// 			},
// 			{
// 				ID:    999,
// 				Title: fmt.Sprintf("emerald"),
// 			},
// 		}
// 	}

// 	if len(data) == 0 {
// 		return
// 	}

// 	err := conn.WriteJSON(data)
// 	if err != nil {
// 		log.Printf("write: %s\n", err)
// 	}
// })
