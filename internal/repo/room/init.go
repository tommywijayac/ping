package room

import (
	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/model"
)

type Repo struct {
	rooms []model.Room //source-of-truth. retain room order as stated in config.
	queue chan model.Room
}

func New(rcfg []config.RoomConfig) *Repo {
	rooms := []model.Room{}
	for _, r := range rcfg {
		rooms = append(rooms, model.Room{
			ID:       r.ID,
			Title:    r.Title,
			IconPath: r.IconPath,
		})
	}

	return &Repo{
		rooms: rooms,
		queue: make(chan model.Room, 100),
	}
}
