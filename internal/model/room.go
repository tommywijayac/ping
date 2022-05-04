package model

type RawRoom struct {
	ID        int
	Timestamp int64
}

type Room struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	IconPath string `json:"icon_path"`

	State              string `json:"state"`
	ConsecutivePing    int
	FirstPingTimestamp int64
	LastPingTimestamp  int64
}
