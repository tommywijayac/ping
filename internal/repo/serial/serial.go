package serial

type Serial struct {
	channel chan int
}

func New(port string) Serial {
	serial := Serial{
		channel: make(chan int, 100), //TODO: 100 move into config?
	}

	return serial
}
