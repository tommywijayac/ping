package display

import (
	"fmt"
	"time"
)

type Display struct {
	roomState [6]bool
}

func New() Display {
	return Display{}
}

//SetDisplayState is a callback function that receive data from serial to determine room state
func (d *Display) SetDisplayState(data int) {
	//TODO
	_ = translateDataIntoRoom(data)

	fmt.Printf("received %d\n", data)

	//test, simulate length operation longer than the data. oke2 aja bisa, channel nya jadi buffer
	//TODO: what should we do if channel is full
	time.Sleep(5 * time.Second)
}

func translateDataIntoRoom(data int) int {
	//TODO
	return 0
}
