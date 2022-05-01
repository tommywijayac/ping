package display

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/tommywijayac/ping/internal/model"
	"github.com/tommywijayac/ping/internal/repo/serial"
)

type Usecase struct {
	repoSerial *serial.Repo
	roomState  [6]model.Room
}

func New(serial *serial.Repo) *Usecase {
	return &Usecase{
		repoSerial: serial,
	}
}

//SendRoomPing send a room ping to client
func (u *Usecase) SendRoomPing(conn *websocket.Conn) error {
	stream := u.repoSerial.Stream()

	for {
		select {
		case data := <-stream:
			conn.WriteJSON(data)
			log.Printf("writing %v", data)
		}
	}
}

//ReceiveRoomPingAck receives a room ping acknowledgement from client,
//and will update room state in memory.
func (u *Usecase) ReceiveRoomPingAck(room model.Room) error {
	fmt.Println(room)
	return nil
}

//[DEV] DevPush push data into serial repo, which should trigger
//display client update.
func (u *Usecase) DevPush(data int) {
	u.repoSerial.Push(data)
}
