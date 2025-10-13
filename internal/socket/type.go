package socket

import "encoding/json"

type TEnvelope struct {
    Type  string          `json:"type"`
    Payload  json.RawMessage `json:"payload,omitempty"`
}

type TRegisterAgentPayload struct {
	Hostname    string     	`json:"hostname"`
	OS          string     	`json:"os"`
	Kernel      string     	`json:"kernel"`
	Arch      	string     	`json:"arch"`
	IPAddress   string		`json:"ip_address"`
	Cores    	uint16 	   	`json:"cores"`
	Threads   	uint16   	`json:"threads"`	
	Memory 		uint64 		`json:"memory"`
}

type TConnectPayload struct {
    Token string `json:"token"`
    NodeId string `json:"node_id"`
}