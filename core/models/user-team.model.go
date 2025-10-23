package models

import (
	"time"
)

type UserTeam struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	UserID     int64     `gorm:"not null;index" json:"user_id"`
	TeamID     int64     `gorm:"not null;index" json:"team_id"`
	InvitedBy  *int64    `json:"invited_by"`
	JoinedAt   time.Time  `gorm:"autoCreateTime" json:"joined_at"`
	IsDefault  bool       `gorm:"default:false" json:"is_default"`
	IsAdmin  bool       `gorm:"default:false" json:"is_admin"`

	User     User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Team     Team `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Inviter  *User `gorm:"foreignKey:InvitedBy" json:"inviter,omitempty"`
}
