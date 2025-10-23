package models

import "time"

type Node struct {
	BaseModel

	Name       string    `gorm:"not null" json:"name"`
	Hostname   string    `json:"hostname"`
	IPAddress  string    `json:"ip_address"`
	OS         string    `json:"os"`
	Kernel     string    `json:"kernel"`
	Arch       string    `json:"arch"`
	Cores      int       `json:"cores"`
	Threads    int       `json:"threads"`
	Memory     int64     `json:"memory"`
	Online     bool      `gorm:"default:false" json:"online"`
	Token      string    `json:"token"`
	Verified   bool      `gorm:"default:false" json:"verified"`
	LastSeen   time.Time `json:"last_seen"`

	ResourceNodes []ResourceNode `gorm:"foreignKey:NodeID" json:"resource_nodes,omitempty"`
}
