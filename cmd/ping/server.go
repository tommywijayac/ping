package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const server string = "localhost:3000"
const client string = "http://localhost:8080" //the only IP we expect requests are coming from

var upgrader = websocket.Upgrader{}

func initServer() {

	http.HandleFunc("/ping", handlerClientConn)

	fmt.Printf("Server starting at %s\n", server)
	go http.ListenAndServe(server, nil)
}

func handlerClientConn(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("origin")
		if origin == client {
			return true
		}
		return false
	}
	c, err := upgrader.Upgrade(w, r, w.Header())

	if err != nil {
		log.Fatal("[main] upgrade err: ", err)
	}
	defer c.Close()

	for {
		//TODO: this is INF loop to receive something from client
		// hence, this should:
		// 1. translate into set member
		// 2. remove that set
		// 3. send request to update the client with current set (using websocket)
		// (this way source of truth is in BE)
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("[main] read err: %s\n", err)
		}

		//TODO this code should go to usecase

		//TODO: 1. process message

		//3. send whole client data (refresh all: BE source of truth)
		now := time.Now()

		type foo struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		}
		data := []foo{
			{
				ID:    1,
				Title: fmt.Sprintf("carla %d", now.Unix()),
			},
			{
				ID:    999,
				Title: fmt.Sprintf("ruby %d", now.Unix()),
			},
		}
		databy, _ := json.Marshal(data)

		//TODO: temp to check websocket server
		err = c.WriteMessage(mt, databy)
		if err != nil {
			log.Println("write:", err)
			break
		}

		//TODO: change this?
		log.Printf("receive: %s\n", message)
	}
}
