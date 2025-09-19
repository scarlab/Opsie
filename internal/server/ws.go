package server

import (
	"database/sql"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Client represents a connected agent
type Client struct {
	Conn   *websocket.Conn
	NodeID string
	Send   chan []byte // outbound messages to this client
}

// Manager handles connected clients
type Manager struct {
	mu      sync.RWMutex
	clients map[string]*Client
	db      *sql.DB
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: tighten this for security (origin check)
	},
}

func NewManager(db *sql.DB) *Manager {
	return &Manager{
		clients: make(map[string]*Client),
		db:      db,
	}
}

// HandleAgentWS upgrades HTTP to WebSocket and registers agent
func (m *Manager) HandleAgentWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// Example: read auth token from query
	token := r.URL.Query().Get("token")
	if token == "" || !m.validateToken(token) {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "unauthorized"))
		conn.Close()
		return
	}

	nodeID := "node-" + token // just for demo; real system generates/registers node
	client := &Client{
		Conn:   conn,
		NodeID: nodeID,
		Send:   make(chan []byte, 256),
	}

	// Save client
	m.mu.Lock()
	m.clients[nodeID] = client
	m.mu.Unlock()

	log.Printf("Agent %s connected", nodeID)

	// Start goroutines for read & write
	go m.readPump(client)
	go m.writePump(client)
}

func (m *Manager) readPump(c *Client) {
	defer func() {
		m.disconnect(c)
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		log.Printf("Metrics from %s: %s", c.NodeID, string(msg))
		// TODO: save metrics to DB
	}
}

func (m *Manager) writePump(c *Client) {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

func (m *Manager) disconnect(c *Client) {
	m.mu.Lock()
	delete(m.clients, c.NodeID)
	m.mu.Unlock()
	c.Conn.Close()
	log.Printf("Agent %s disconnected", c.NodeID)
}

// SendInstruction sends a task to an agent
func (m *Manager) SendInstruction(nodeID string, instruction []byte) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if client, ok := m.clients[nodeID]; ok {
		client.Send <- instruction
	}
}

// dummy token validation
func (m *Manager) validateToken(t string) bool {
	return len(t) > 5
}
