package display

import (
	"fmt"
	"log"
	"time"
)

//ReceiveRoomPingAck receives a room ping acknowledgement from client
//and set server room state to inactive, keeping source-of-truth up to date.
func (u *Usecase) ReceiveRoomPingAck(roomID int) error {
	elapsed := int64(0)

	//use simplest approach O(n) since rooms would rarely grow
	for i := range u.rooms {
		if u.rooms[i].ID == roomID {
			//calculate stats
			elapsed = time.Now().Unix() - u.rooms[i].FirstPingTimestamp
			log.Printf("[stats] elapsed: %dsec. consecutive: %d", elapsed, u.rooms[i].ConsecutivePing)

			//reset state and attributes
			u.rooms[i].State = ""
			u.rooms[i].ConsecutivePing = 0
			u.rooms[i].FirstPingTimestamp = 0
			u.rooms[i].LastPingTimestamp = 0

			break
		}
		return fmt.Errorf("receive unknown room ID: %d\n", roomID)
	}

	log.Printf("[usecase] receive room ping ack: %d\n", roomID)
	return nil
}
