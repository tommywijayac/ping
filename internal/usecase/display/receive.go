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

	r, err := u.repoRoom.Get(roomID)
	if err != nil {
		return fmt.Errorf("[usecase] fail to get room attributes: %w", err)
	}

	//calculate stats
	elapsed = time.Now().Unix() - r.FirstPingTimestamp
	log.Printf("[stats] elapsed: %dsec. consecutive: %d", elapsed, r.ConsecutivePing)

	//reset state and attributes
	if err := u.repoRoom.SetAttributes(r.ID, "", 0, 0); err != nil {
		return fmt.Errorf("[usecase] fail to reset room attributes: %w", err)
	}

	log.Printf("[usecase] receive room ping ack: %d\n", roomID)
	return nil
}
