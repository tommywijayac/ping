package display

import (
	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/model"
	"github.com/tommywijayac/ping/internal/repo/serial"
)

type Usecase struct {
	repoSerial *serial.Repo
	rooms      []model.Room //state source-of-truth. retain room order as stated in config
}

func New(rcfg []config.RoomConfig, serial *serial.Repo) *Usecase {
	rooms := []model.Room{}
	for _, r := range rcfg {
		rooms = append(rooms, model.Room{
			ID:       r.ID,
			Title:    r.Title,
			IconPath: r.IconPath,
		})
	}

	return &Usecase{
		repoSerial: serial,
		rooms:      rooms,
	}
}

//[DEV] DevPush push room ID into serial repo, which should trigger
//server to update client room state
func (u *Usecase) DevPush(roomID int) {
	u.repoSerial.Push(roomID)
}
