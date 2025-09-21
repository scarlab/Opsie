package ws_agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"watchtower/internal/socket"

	"github.com/gorilla/websocket"
)

type Handler struct {
	service *Service
	hub *socket.Hub
}

func NewHandler(service *Service,hub *socket.Hub) *Handler {
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

	// 
	log.Printf("Agent Connected: %s", conn.LocalAddr())

	// Read register message first
	_, msg, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return
	}

	var reg struct {
		NodeID string `json:"node_id"`
	}
	if err := json.Unmarshal(msg, &reg); err != nil || reg.NodeID == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("invalid register"))
		conn.Close()
		return
	}

	client := &socket.Client{
		ID:   reg.NodeID,
		Type: socket.ClientAgent,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	h.hub.RegisterAgent(reg.NodeID, client)

	// Start goroutines
	go h.writePump(client)
	go h.readPump(client)
}

func (h *Handler) readPump(c *socket.Client) {
	defer func() {
		h.hub.UnregisterAgent(c.ID)
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("Agent %s sent: %s\n", c.ID, msg)
		h.hub.BroadcastToUI(msg)
	}
}

func (h *Handler) writePump(c *socket.Client) {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
