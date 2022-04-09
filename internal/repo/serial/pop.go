package serial

type ReceiveCallback func(data int)

//Pop is a blocking function that invokes callback function when channel receives data
//TODO: must handle SIGTERM
func (s *Serial) Receive(callback ReceiveCallback) {
	for {
		select {
		case data := <-s.channel:
			callback(data)
		}
	}
}
