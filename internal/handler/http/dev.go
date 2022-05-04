package http

import (
	"net/http"
	"strconv"
)

func (h *Handler) HandlerDevPush(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.ParseInt(r.URL.Query().Get("room_id"), 10, 64)
	h.ucDisplay.DevPush(int(roomID))
}
