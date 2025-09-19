// internal/server/ws.go
package server

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)
type SocketHub struct {
	mu        sync.Mutex
	agents    map[string]*websocket.Conn
	uiClients map[*websocket.Conn]bool
}

func NewSocketHub() *SocketHub {
	return &SocketHub{
		agents:    make(map[string]*websocket.Conn),
		uiClients: make(map[*websocket.Conn]bool),
	}
}

func (h *SocketHub) RegisterAgent(nodeID string, conn *websocket.Conn) {
	h.mu.Lock()
	h.agents[nodeID] = conn
	h.mu.Unlock()
	log.Println("Agent registered:", nodeID)
}

func (h *SocketHub) UnregisterAgent(nodeID string) {
	h.mu.Lock()
	delete(h.agents, nodeID)
	h.mu.Unlock()
	log.Println("Agent disconnected:", nodeID)
}

func (h *SocketHub) RegisterUI(conn *websocket.Conn) {
	h.mu.Lock()
	h.uiClients[conn] = true
	h.mu.Unlock()
	log.Println("UI connected")
}

func (h *SocketHub) UnregisterUI(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.uiClients, conn)
	h.mu.Unlock()
	log.Println("UI disconnected")
}

func (h *SocketHub) BroadcastToUI(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for conn := range h.uiClients {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			conn.Close()
			delete(h.uiClients, conn)
		}
	}
}
