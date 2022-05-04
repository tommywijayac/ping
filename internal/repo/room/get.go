package room

import (
	"fmt"

	"github.com/tommywijayac/ping/internal/model"
)

func (r *Repo) GetAll() []model.Room {
	return r.rooms
}

func (r *Repo) Get(id int) (model.Room, error) {
	for i := range r.rooms {
		if r.rooms[i].ID == id {
			return r.rooms[i], nil
		}
	}
	return model.Room{}, fmt.Errorf("room ID not found: id [%d] in %v", id, r.rooms)
}
