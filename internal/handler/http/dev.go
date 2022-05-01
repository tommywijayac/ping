package http

import (
	"net/http"
	"strconv"
)

func (h *Handler) HandlerDevPush(w http.ResponseWriter, r *http.Request) {
	data, _ := strconv.ParseInt(r.URL.Query().Get("data"), 10, 64)
	h.ucDisplay.DevPush(int(data))
}
