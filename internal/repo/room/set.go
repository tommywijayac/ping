package room

import (
	"fmt"
)

func (r *Repo) SetAttributes(id int, state string, count int, ts int64) error {
	for i := range r.rooms {
		if r.rooms[i].ID == id {
			r.rooms[i].State = state

			r.rooms[i].ConsecutivePing = count

			if r.rooms[i].FirstPingTimestamp == 0 {
				r.rooms[i].FirstPingTimestamp = ts
			}
			r.rooms[i].LastPingTimestamp = ts

			return nil
		}
	}
	return fmt.Errorf("room ID not found: id [%d] in %v", id, r.rooms)
}
