package serial

import (
	"fmt"
	"time"
)

//Pop is a blocking function that invokes callback function when channel receives data
//TODO: add callback fn as parameter
//TODO: must handle SIGTERM
func (s *Serial) Receive() {
	for {
		select {
		case data := <-s.channel:
			fmt.Printf("received %d\n", data)

			//TODO: test, simulate length operation longer than the data. oke2 aja bisa, channel nya jadi buffer
			//TODO: what should we do if channel is full
			time.Sleep(5 * time.Second)
		}
	}
}
