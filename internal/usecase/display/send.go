package display

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//SendAllRoomAttributes sends all registered rooms attributes to client
//via supplied websocket connection.
func (u *Usecase) SendAllRoomAttributes(conn *websocket.Conn) error {
	if err := conn.WriteJSON(u.rooms); err != nil {
		return fmt.Errorf("[usecase] fail to send all room attributes: %v", err)
	}

	log.Printf("[usecase] send all room attributes: %+v\n", u.rooms)
	return nil
}

//TODO: function to read queue channel and send to client

//TODO: this should receives data from serial channel, and put into display channel with rules:
//1. sequential display call spanning less than 5 seconds is merged into one (the first)
//2. ...
//this means we need to keep track of data put in display channel
//display channel acts as a queue ONLY, to trigger action (send to client)

//SendRoomPing is a blocking function that sends a room state to client
//every time new data is received in display repo channel, via supplied websocket connection.
func (u *Usecase) SendRoomPing(conn *websocket.Conn) error {
	stream := u.repoSerial.Stream()

	for {
		select {
		case raw := <-stream:
			id := raw.ID
			ts := raw.Timestamp

			for i := range u.rooms {
				if u.rooms[i].ID == id {
					u.rooms[i].State = "active"

					//TODO: if fulfill the requirement, put into queue channel
					u.rooms[i].ConsecutivePing++
					if u.rooms[i].FirstPingTimestamp == 0 {
						u.rooms[i].FirstPingTimestamp = ts
					}
					u.rooms[i].LastPingTimestamp = ts

					break
				}
				return fmt.Errorf("receive unknown room ID: %d\n", id)
			}

			if err := conn.WriteJSON(u.rooms); err != nil {
				return fmt.Errorf("[usecase] fail to send all room attributes: %v", err)
			}
			log.Printf("[usecase] send room states %+v", u.rooms)

			//TODO: debug oto cleanup first
			// if err := oto.PlayPingSound(); err != nil {
			// 	log.Printf("error playing ping sound: %s\n", err)
			// }
		}
	}
}
