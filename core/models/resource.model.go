package models

import (
	"time"

	"gorm.io/datatypes"
)

type Resource struct {
	BaseModel

	TeamID     int64          	`gorm:"not null;index" json:"team_id"`
	ProjectID  *int64         	`gorm:"index" json:"project_id"`
	Name       string          		`gorm:"not null" json:"name"`
	Description string         		`json:"description"`
	Type       string          		`json:"type"`
	Ports      datatypes.JSON  		`gorm:"type:jsonb" json:"ports"`
	Env        datatypes.JSON  		`gorm:"type:jsonb" json:"env"`
	Replicas   int             		`gorm:"default:1" json:"replicas"`
	Status     string				`gorm:"default:'stopped'" json:"status"`
	IsArchived bool            		`gorm:"default:false" json:"is_archived"`
	ArchivedAt *time.Time      		`json:"archived_at"`

	Team          Team           	`gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Project       *Project       	`gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	ResourceNodes []ResourceNode 	`gorm:"foreignKey:ResourceID" json:"resource_nodes,omitempty"`
}
