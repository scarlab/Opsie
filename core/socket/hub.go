package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientType int

const (
	ClientAgent ClientType = iota
	ClientUI
)

type Client struct {
	ID   string
	Type ClientType
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	mu     sync.Mutex
	agents map[string]*Client
	uis    map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		agents: make(map[string]*Client),
		uis:    make(map[string]*Client),
	}
}

func (h *Hub) RegisterAgent(id string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.agents[id] = c
}

func (h *Hub) UnregisterAgent(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c, ok := h.agents[id]; ok {
		close(c.Send)
		delete(h.agents, id)
	}
}

func (h *Hub) RegisterUI(id string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.uis[id] = c
}

func (h *Hub) UnregisterUI(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c, ok := h.uis[id]; ok {
		close(c.Send)
		delete(h.uis, id)
	}
}

func (h *Hub) BroadcastToUI(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for id, c := range h.uis {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(h.uis, id)
		}
	}
}

func (h *Hub) BroadcastToAgent(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for id, c := range h.agents {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(h.agents, id)
		}
	}
}

func (h *Hub) SendToUI(id string, msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c, ok := h.uis[id]; ok {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(h.uis, id)
		}
	}
}

func (h *Hub) SendToAgent(id string, msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c, ok := h.agents[id]; ok {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(h.agents, id)
		}
	}
}
