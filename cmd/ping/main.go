package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tommywijayac/ping/internal/repo/serial"
	"github.com/tommywijayac/ping/internal/usecase/display"
)

func main() {
	serial := serial.New("test")
	display := display.New()

	//TODO: move to usecase
	go serial.Listen("test")
	go serial.Receive(display.SetDisplayState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
	fmt.Println("closing")
}
