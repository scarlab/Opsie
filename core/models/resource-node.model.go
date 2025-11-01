package models

type ResourceNode struct {
	ID         int64         `gorm:"primaryKey" json:"id,string"`
	ResourceID int64         `gorm:"not null;index" json:"resource_id,string"`
	NodeID     int64         `gorm:"not null;index" json:"node_id,string"`
	Status     string `gorm:"default:'stopped'" json:"status"`

	Resource Resource `gorm:"foreignKey:ResourceID" json:"resource,omitempty"`
	Node     Node     `gorm:"foreignKey:NodeID" json:"node,omitempty"`
}
