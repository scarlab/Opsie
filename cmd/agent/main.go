package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	nodeID := "node-1"
	url := "ws://192.168.0.202:3905/api/v1/ws/agent"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// First message = register
	regMsg, _ := json.Marshal(map[string]string{
		"node_id": nodeID,
	})
	conn.WriteMessage(websocket.TextMessage, regMsg)

	// Send metrics every 5s
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		metrics := map[string]interface{}{
			"cpu": rand.Intn(100),          // 0–99%
			"mem": rand.Intn(16000) + 500,  // 500–16500 MB
		}
		msg, _ := json.Marshal(metrics)
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
