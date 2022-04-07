package serial

import (
	"fmt"
	"math/rand"
	"time"
)

//Listen is a blocking function that continuously listen to a serial port and push received data to channel
//TODO: must handle SIGTERM
func (s *Serial) Listen(port string) {
	for {
		//TODO: change into real serial listener
		data := rand.Intn(100)
		s.push(data)
		fmt.Printf("pushed %d into channel\n", data)

		time.Sleep(2 * time.Second)
	}
}

//Push will put data into channel
func (s *Serial) push(data int) {
	s.channel <- data
}
