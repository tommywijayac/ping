package display

import (
	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/repo/room"
	"github.com/tommywijayac/ping/internal/repo/serial"
)

type Usecase struct {
	pingDelay  int64
	repoSerial *serial.Repo
	repoRoom   *room.Repo
}

func New(cfg *config.Config, serial *serial.Repo, room *room.Repo) *Usecase {
	return &Usecase{
		pingDelay:  cfg.PingDelay,
		repoSerial: serial,
		repoRoom:   room,
	}
}

//[DEV] DevPush push room ID into serial repo, which should trigger
//server to update client room state
func (u *Usecase) DevPush(roomID int) {
	u.repoSerial.Push(roomID)
}
