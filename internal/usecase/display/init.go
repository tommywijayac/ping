package display

import (
	"fmt"

	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/model"
	"github.com/tommywijayac/ping/internal/repo/serial"
)

type Usecase struct {
	repoSerial *serial.Repo
	rooms      []model.Room //state source-of-truth. retain room order
}

func New(rcfg []config.RoomConfig, serial *serial.Repo) *Usecase {
	rooms := []model.Room{}
	for _, r := range rcfg {
		rooms = append(rooms, model.Room{
			ID:       r.ID,
			Title:    r.Title,
			IconPath: r.IconPath,
			State:    "", //all begin with inactive (empty)
		})
	}

	return &Usecase{
		repoSerial: serial,
		rooms:      rooms,
	}
}

//ReceiveRoomPingAck receives a room ping acknowledgement from client,
//and will set room state to inactive.
func (u *Usecase) ReceiveRoomPingAck(roomID int) error {
	fmt.Println(roomID)

	for i := range u.rooms {
		if u.rooms[i].ID == roomID {
			u.rooms[i].State = ""
			u.rooms[i].ConsecutivePing = 0
			u.rooms[i].FirstPing = 0
		}
	}

	fmt.Println(u.rooms)

	return nil
}

//[DEV] DevPush push data into serial repo, which should trigger
//display client update.
func (u *Usecase) DevPush(data int) {
	u.repoSerial.Push(data)
}
