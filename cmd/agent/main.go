package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	nodeID := "node-4"
	url := "ws://localhost:3905/api/v1/ws/agent"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// First message = register
	conn.WriteMessage(websocket.TextMessage, []byte(nodeID))

	// Send metrics every 5s
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		metrics := `{"cpu":44,"mem":444}`
		conn.WriteMessage(websocket.TextMessage, []byte(metrics))
	}
}
