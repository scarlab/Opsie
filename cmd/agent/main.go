package main

import (
	"log"
	"opsie/config"
	"opsie/internal/socket"
	"opsie/pkg/system"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	url := "ws://"+config.ENV.ServerHost+"/api/v1/ws/agent"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Get the system info
	systemInfo := system.Info()
	payload := socket.TRegisterAgentPayload{
		Hostname:  systemInfo.Hostname,
		OS:        systemInfo.OS,
		Kernel:    systemInfo.Kernel,
		Arch:      systemInfo.Arch,
		IPAddress: systemInfo.IPAddress,
		Cores:     systemInfo.Cores,
		Threads:   systemInfo.Threads,
		Memory:    systemInfo.Memory,
	}

	regMsg, err := socket.MarshalEnvelope("register", payload)
	if err != nil { log.Println(err) }


	// Send over WebSocket
	if err := conn.WriteMessage(websocket.TextMessage, regMsg); err != nil {
		log.Fatal("failed to send register message:", err)
	}


	// Send metrics every 5s
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		// Collect real system metrics
		metrics := system.Metrics()

		msg, err := socket.MarshalEnvelope("metrics", metrics)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("failed to send metrics:", err)
			return
		}
	}
}
