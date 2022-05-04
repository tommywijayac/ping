package http

import (
	"log"
	"net/http"
	"strconv"
)

const client string = "http://localhost:8080" //the only IP we expect requests are coming from

func (h *Handler) HandlerClientWebsocket(w http.ResponseWriter, r *http.Request) {
	h.upgrader.CheckOrigin = func(r *http.Request) bool {
		//whitelist client
		origin := r.Header.Get("origin")
		if origin == client {
			return true
		}
		log.Printf("[http][handler] handshake attempt by unknown client: %s\n", origin)
		return false
	}

	var err error
	h.conn, err = h.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		log.Fatalf("[http][handler] upgrade to websocket err: %s", err)
	}

	go func() {
		for {
			//ReadMessage is a blocking function that waits for new message.
			_, message, err := h.conn.ReadMessage()
			if err != nil {
				// if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				// 	return
				// }

				log.Printf("[http][handler] read websocket message err: %s\n", err)
				return
			}

			//client is expected to send room ID
			roomID, err := strconv.Atoi(string(message))
			if err != nil {
				log.Printf("[http][handler] fail unmarshal websocket message err: %s\n", err)
				continue
			}

			h.ucDisplay.ReceiveRoomPingAck(roomID)
		}
	}()
	//if websocket conn isn't created, then will not increment waitgroup
	h.appWg.Add(1)

	//cleanup by closing data channel, which should be on repo
	go func() {
		err := h.ucDisplay.SendRoomPing(h.conn)
		if err != nil {
			log.Printf("[http][handler] write websocket message err: %s\n", err)
		}
	}()

	//send default room config
	if err := h.ucDisplay.SendAllRoomAttributes(h.conn); err != nil {
		log.Printf("[http][handler] send default room err: %s\n", err)
	}

	//wait for close signal from app
	<-h.close

	//cleanup
	h.conn.Close()
	log.Println("http: handler: done closing websockets")

	//tell main http module has finished cleanup
	h.appWg.Done()
}
