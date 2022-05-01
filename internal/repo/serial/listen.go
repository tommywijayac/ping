package serial

import (
	"bufio"
	"fmt"
	"time"

	tarmserial "github.com/tarm/serial"
	"github.com/tommywijayac/ping/internal/model"
)

//Listen continuously listen to a serial port and push received data to channel.
//TODO: must handle SIGTERM
func (r *Repo) Listen(port string) error {
	stream, err := tarmserial.OpenPort(&tarmserial.Config{
		Name: "", //COMxx in Windows or /dev/tty in Linux
		Baud: 9600,
		Size: 8,
	})
	if err != nil {
		//TODO: return err
		fmt.Println(err)
		//return fmt.Errorf("fail to open serial port: %w", err)
	}

	scanner := bufio.NewScanner(stream)
	//TODO: listen to actual scanner
	fmt.Println(scanner)

	for {
		//TODO: change into real serial listener
		//data := rand.Intn(100)
		// s.push(data)
		//fmt.Printf("pushed %d into channel\n", data)

		time.Sleep(2 * time.Second)
	}
}

//Push will put data into channel
func (r *Repo) push(data int) {
	r.channel <- model.RawRoom{
		Id:        data,
		Timestamp: time.Now().Unix(),
	}
}

//[DEV] Push put an integer value into display channel
func (r *Repo) Push(data int) {
	r.push(data)
}
