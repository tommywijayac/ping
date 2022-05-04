package room

import "github.com/tommywijayac/ping/internal/model"

func (r *Repo) Queue() <-chan model.Room {
	return r.queue
}

func (r *Repo) Push(room model.Room) {
	r.queue <- room
}
