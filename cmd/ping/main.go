package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tommywijayac/ping/internal/config"
	"github.com/tommywijayac/ping/internal/pkg/oto"
	"github.com/tommywijayac/ping/internal/repo/serial"

	"github.com/tommywijayac/ping/internal/usecase/display"

	"github.com/tommywijayac/ping/internal/handler/http"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	closeSig := make(chan bool, 1)

	//component close signal
	wsClose := make(chan bool, 1)

	appWg := sync.WaitGroup{}
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		closeSig <- true
		wsClose <- true
	}()

	//init config
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("main: fail to init config: %v\n", err)
	}

	//init repo
	repoSerial := serial.New("test")

	//init usecase
	usecaseDisplay := display.New(cfg.RoomConfig, &repoSerial)

	//init handler
	handlerHttp := http.New(usecaseDisplay, wsClose, &appWg)
	//TODO:
	//handlerSerial.. initSerialListener?

	initServer(handlerHttp)

	//wait for interrupt or terminate signal
	<-closeSig

	fmt.Println("main: doing cleanup..")
	oto.Close()
	appWg.Wait()

	fmt.Println("main: terminated")
}
