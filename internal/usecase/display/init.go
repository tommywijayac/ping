package display

import (
	"github.com/tommywijayac/ping/internal/repo/room"
	"github.com/tommywijayac/ping/internal/repo/serial"
)

type Usecase struct {
	repoSerial *serial.Repo
	repoRoom   *room.Repo
}

func New(serial *serial.Repo, room *room.Repo) *Usecase {
	return &Usecase{
		repoSerial: serial,
		repoRoom:   room,
	}
}

//[DEV] DevPush push room ID into serial repo, which should trigger
//server to update client room state
func (u *Usecase) DevPush(roomID int) {
	u.repoSerial.Push(roomID)
}
