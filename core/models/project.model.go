package models

import (
	"time"
)

type Project struct {
	BaseModel

	TeamID      int64 `gorm:"not null;index" json:"team_id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ArchivedAt  *time.Time `json:"archived_at"`

	Team      Team        `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Resources []Resource  `gorm:"foreignKey:ProjectID" json:"resources,omitempty"`
}
