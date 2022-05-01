package model

type RawRoom struct {
	Id        int
	Timestamp int64
}

type Room struct {
	Title string `json:"title"`
	State string `json:"state"`
}
