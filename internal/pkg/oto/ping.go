package oto

import (
	"os"
)

func PlayPingSound() error {
	filepath := os.Getenv("GOPATH") + "\\src\\ping\\files\\ping.mp3"
	return Play(filepath)
}
