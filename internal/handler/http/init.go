package http

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tommywijayac/ping/internal/usecase/display"
)

type Handler struct {
	upgrader  websocket.Upgrader
	conn      *websocket.Conn
	ucDisplay *display.Usecase

	//internal cleanup
	close <-chan bool
	appWg *sync.WaitGroup
}

func New(display *display.Usecase, c <-chan bool, wg *sync.WaitGroup) *Handler {
	return &Handler{
		ucDisplay: display,
		close:     c,
		appWg:     wg,
	}
}
