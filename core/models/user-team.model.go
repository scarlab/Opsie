package models

import (
	"time"
)

type UserTeam struct {
	ID         int64     `gorm:"primaryKey" json:"id,string"`
	UserID     int64     `gorm:"not null;index" json:"user_id,string"`
	TeamID     int64     `gorm:"not null;index" json:"team_id,string"`
	JoinedAt   time.Time  `gorm:"autoCreateTime" json:"joined_at"`
	IsDefault  bool       `gorm:"default:false" json:"is_default"`

	User     User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Team     Team `gorm:"foreignKey:TeamID" json:"team,omitempty"`
}


type TeamMember struct {
	ID  				int64    		`json:"id,string"`
	TeamId  			int64    		`json:"team_id,string"`
	DisplayName  		string    		`json:"display_name"`
	Email  				string    		`json:"email"`
	Avatar  			string    		`json:"avatar"`
	SystemRole  		string    		`json:"system_role"`
	IsActive  			bool    		`json:"is_active"`
	JoinedAt   			time.Time  		`json:"joined_at"`
}

 


type AddUserToTeamPayload struct {
	UserID    int64   `json:"user_id,string" validate:"required"`
	TeamID    int64   `json:"team_id,string" validate:"required"`
	IsDefault bool    `json:"is_default"`
}


type RemoveUserToTeamPayload struct {
	UserID    int64   `json:"user_id,string" validate:"required"`
	TeamID    int64   `json:"team_id,string" validate:"required"`
}
