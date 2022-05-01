package serial

import "github.com/tommywijayac/ping/internal/model"

func (r *Repo) Stream() <-chan model.RawRoom {
	return r.channel
}
