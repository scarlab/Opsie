package ws_agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"opsie/core/socket"
	"opsie/pkg/system"

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

	// ...
	
	// payload.data example: {"hostname":"agent1","ip_address":"xx.x.x.xx","os":"linux","arch":"amd64","cpu_cores":4,"memory_gb":8}
	// payload.type example: "register"
	// payload.token:

	// if token is  available, validate it here
	// if token in expired or not available, reject the connection
	// and check for data. if data is not available, reject the connection
	// if data - register.

	log.Printf("Agent Connected: %s", conn.LocalAddr())

	// Read register message first
	_, msg, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return
	}


	// Wrap struct to match what agent sends
	envelope, err := socket.UnmarshalEnvelope(msg)
	if err != nil { log.Fatal(err) }


	switch envelope.Type {
		case "register":
			metrics, err := socket.DecodePayload[socket.TRegisterAgentPayload](envelope)
			if err != nil { log.Fatal(err) }
			
			log.Printf("Register request from %s (%s)", metrics.Hostname, metrics.IPAddress)
			fmt.Println(metrics)

		case "connect":
			var payload socket.TConnectPayload
			if err := json.Unmarshal(envelope.Payload, &payload); err != nil {
				log.Println("invalid connect payload:", err)
				return
			}
			log.Printf("Reconnect with token: %s", payload.Token)

		default:
			log.Println("unknown message type:", envelope.Type)
	}

	// TODO: Prform database operation
	// client ID should be the database ID of the agent

	client := &socket.Client{
		ID:   "ID_333",  // it should be database ID
		Type: socket.ClientAgent,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	h.hub.RegisterAgent("ID_333", client)

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
			fmt.Println("read error:", err)
			return
		}

		// unmarshal envelope
		var envelope socket.TEnvelope
		if err := json.Unmarshal(msg, &envelope); err != nil {
			fmt.Println("invalid message envelope:", err)
			continue
		}

		// handle by type
		switch envelope.Type {
		case "metrics":
			var metrics system.SystemMetrics
			if err := json.Unmarshal(envelope.Payload, &metrics); err != nil {
				fmt.Println("invalid metrics payload:", err)
				continue
			}
			// Store / process metrics
			// h.service.SaveMetrics(c.ID, metrics)

			// Optionally broadcast to UI
			h.hub.BroadcastToUI(msg)

			log.Println(metrics)

		case "command_response":
			// handle command response from agent
			// h.service.HandleCommandResponse(c.ID, envelope.Payload)
			h.hub.BroadcastToUI(msg)

		default:
			fmt.Printf("unknown message type from agent %s: %s\n", c.ID, envelope.Type)
		}
	}
}


func (h *Handler) writePump(c *socket.Client) {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
