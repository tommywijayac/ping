package display

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tommywijayac/ping/internal/pkg/oto"
)

func (u *Usecase) SendAllRoomStates(conn *websocket.Conn) error {
	if err := conn.WriteJSON(u.rooms); err != nil {
		return fmt.Errorf("usecase: fail to write json to websocket: %v", err)
	}

	return nil
}

//SendRoomPing send a room ping to client
func (u *Usecase) SendRoomPing(conn *websocket.Conn) error {
	stream := u.repoSerial.Stream()

	for {
		select {
		case data := <-stream:
			roomID := data.ID

			for i := range u.rooms {
				if u.rooms[i].ID == roomID {
					u.rooms[i].State = "active"
					u.rooms[i].ConsecutivePing++
					if u.rooms[i].FirstPing == 0 {
						u.rooms[i].FirstPing = time.Now().Unix()
					}
				}
			}

			conn.WriteJSON(u.rooms)
			log.Printf("writing %v", u.rooms)

			if err := oto.PlayPingSound(); err != nil {
				log.Printf("error playing ping sound: %s\n", err)
			}
		}
	}
}
