package display

import "fmt"

//ReceiveRoomPingAck receives a room ping acknowledgement from client
//and set server room state to inactive, keeping source-of-truth up to date.
func (u *Usecase) ReceiveRoomPingAck(roomID int) error {
	//use simplest approach O(n) since rooms would rarely grow
	for i := range u.rooms {
		if u.rooms[i].ID == roomID {
			u.rooms[i].State = ""
			u.rooms[i].ConsecutivePing = 0
			u.rooms[i].FirstPingTimestamp = 0
			u.rooms[i].LastPingTimestamp = 0
		}
	}

	fmt.Println(u.rooms)

	return nil
}
