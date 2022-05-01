package serial

import (
	"github.com/tommywijayac/ping/internal/model"
)

type Repo struct {
	channel chan model.RawRoom
}

func New(port string) Repo {
	repo := Repo{
		channel: make(chan model.RawRoom, 100), //TODO: 100 move into config?
	}

	//TODO: start listening
	//go repoSerial.Listen("test")

	return repo
}
