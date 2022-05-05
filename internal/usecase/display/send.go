package display

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/tommywijayac/ping/internal/pkg/oto"
)

//SendAllRoomAttributes sends all registered rooms attributes to client
//via supplied websocket connection.
func (u *Usecase) SendAllRoomAttributes(ctx context.Context, conn *websocket.Conn) error {
	rs := u.repoRoom.GetAll()
	if err := conn.WriteJSON(rs); err != nil {
		return fmt.Errorf("fail to write json to websocket: %v", err)
	}

	log.Printf("[SendAllRoomAttributes] send all room attributes: %+v\n", rs)
	return nil
}

//SendRoomPing is a blocking function that sends a room state to client
//via supplied websocket connection, every time new data is received in serial stream
func (u *Usecase) SendRoomPing(ctx context.Context, conn *websocket.Conn) {
	stream := u.repoSerial.Stream()

	for {
		select {
		case <-ctx.Done():
			return
		case raw := <-stream:
			id := raw.ID
			ts := raw.Timestamp

			r, err := u.repoRoom.Get(id)
			if err != nil {
				log.Printf("[usecase] fail to send room ping: %s\n", err)
				continue
			}

			if r.LastPingTimestamp != 0 && ts-r.LastPingTimestamp < u.pingDelay {
				//ignore repeated consecutive ping in short timespan
				continue
			}

			//set room to active, update attributes
			if err := u.repoRoom.SetAttributes(id, "active", r.ConsecutivePing+1, ts); err != nil {
				log.Printf("[usecase] fail to set room attributes: %s\n", err)
				continue
			}

			//send all room attributes (overwrite) instead of individual rooms. reason:
			//1. easier. FE no need to look to their array first
			//2. FE & BE always in sync in all time. no need to build separate mechanism to sync every x min
			//3. vue reactivity is pretty smart. even though array changes, ongoing click animation isn't interrupted
			//   by array overwrites
			if err := u.SendAllRoomAttributes(ctx, conn); err != nil {
				log.Printf("[usecase] fail to send all room attributes: %s\n", err)
				continue
			}

			if err := oto.PlayPingSound(); err != nil {
				log.Printf("error playing ping sound: %s\n", err)
			}
		}
	}
}
