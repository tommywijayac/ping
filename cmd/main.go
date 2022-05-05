package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/repo/room"
	"github.com/tommywijayac/ping/internal/repo/serial"

	"github.com/tommywijayac/ping/internal/usecase/display"

	"github.com/tommywijayac/ping/internal/handler/http"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//component close signals
	appClose := make(chan bool, 1)
	wsClose := make(chan bool, 1)
	//wait group to wait components finish their clean up
	appWg := sync.WaitGroup{}
	go func() {
		sig := <-sigs
		fmt.Println(sig)

		//trigger components clean up
		appClose <- true
		wsClose <- true
	}()

	//init config
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("main: fail to init config: %v\n", err)
	}
	if len(cfg.RoomConfig) == 0 {
		log.Fatalf("[main] no room defined")
	}

	//init repo
	repoSerial := serial.New("test")
	repoRoom := room.New(cfg.RoomConfig)

	//init usecase
	usecaseDisplay := display.New(cfg, repoSerial, repoRoom)

	//init handler
	handlerHttp := http.New(usecaseDisplay, wsClose, &appWg)
	//TODO:
	//handlerSerial.. initSerialListener?

	initServer(handlerHttp)

	//app wait here until termination/interrupt signal
	<-appClose
	//TODO: gracefully oto cleanup without hang
	//oto.Close()

	//wait here for components clean up (if any)
	appWg.Wait()

	fmt.Println("main: terminated")
}
