package ws_ui

import (
	"net/http"
	"watchtower/internal/socket"

	"github.com/gorilla/websocket"
)

type Handler struct {
	service *Service
	hub *socket.Hub
}

func NewHandler(service *Service, hub *socket.Hub) *Handler {
	return &Handler{
		service: service,
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade failed", http.StatusInternalServerError)
		return
	}

	client := &socket.Client{
		ID:   "ui-client",
		Type: socket.ClientUI,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	h.hub.RegisterUI(client)

	go h.writePump(client)
	go h.readPump(client)
}

func (h *Handler) readPump(c *socket.Client) {
	defer func() {
		h.hub.UnregisterUI(c)
		c.Conn.Close()
	}()
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		// UI usually won’t send anything
	}
}

func (h *Handler) writePump(c *socket.Client) {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
