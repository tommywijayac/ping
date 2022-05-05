package http

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
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
		log.Fatalf("[http][handler] fail to upgrade to websocket: %s\n", err)
	}

	ctx, cancel := context.WithCancel(r.Context())
	wait := sync.WaitGroup{}
	waitCh := make(chan struct{})

	//send default room config to client
	if err := h.ucDisplay.SendAllRoomAttributes(ctx, h.conn); err != nil {
		log.Printf("[http][handler] fail to send default room: %s\n", err)
	}

	go func() {
		go func() {
			h.ucDisplay.SendRoomPing(ctx, h.conn)

			//since SendRoomPing is a blocking func, reaching here means for loop has ended (by ctx cancel)
			wait.Done()
		}()
		wait.Add(1)

		go func() {
			for {
				select {
				case <-ctx.Done():
					wait.Done()

					//must use 'return' instead of 'break'
					//somehow 'break' makes the for-loop execute one more time,
					//hence execute wait.Done() one extra time and cause panic at wait.Wait()
					return
				default:
					//ReadMessage is a blocking function that waits for new message.
					_, message, err := h.conn.ReadMessage()
					if err != nil {
						//if tab is refreshed, then will receive 'close 1001'.
						//cancel the context so both processing goroutine will exit
						cancel()

						log.Printf("[http][handler] read websocket error message: %s\n", err)
						continue
					}

					//client is expected to send room ID
					roomID, err := strconv.Atoi(string(message))
					if err != nil {
						log.Printf("[http][handler] fail to unmarshal websocket message: %s\n", err)
						continue
					}

					err = h.ucDisplay.ReceiveRoomPingAck(roomID)
					if err != nil {
						log.Printf("[http][handler] fail to receive room ping ack: %s\n", err)
					}
				}
			}
		}()
		wait.Add(1)

		wait.Wait()
		close(waitCh)
	}()

	//signal app that there's one bg process here. app should wait for this clean up before closing
	h.appWg.Add(1)

	select {
	case <-waitCh:
		log.Println("[http][handler] close because websocket closes")
	case <-h.close:
		log.Println("[http][handler] close because app terminates")

		//write close message. this will automatically trigger the rest of sequence
		cm := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "interrupt/termination signal")
		if err := h.conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
			log.Printf("[http][handler] fail to write close message: %s\n", err)
		}

		//wait until cleanup finish
		<-waitCh
	}

	//cleanup
	h.conn.Close()
	h.appWg.Done() //signal app we're done with the clean up
	log.Println("[http][handler] done closing websockets.")
}
