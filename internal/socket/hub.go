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
	uis    map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		agents: make(map[string]*Client),
		uis:    make(map[*Client]bool),
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
	delete(h.agents, id)
}

func (h *Hub) RegisterUI(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.uis[c] = true
}

func (h *Hub) UnregisterUI(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.uis, c)
}

func (h *Hub) BroadcastToUI(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for c := range h.uis {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(h.uis, c)
		}
	}
}
