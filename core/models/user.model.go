package models

import (
	"gorm.io/datatypes"
)



type User struct {
	BaseModel

	DisplayName string         `gorm:"not null" json:"display_name"`
	Email       string         `gorm:"uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"not null" json:"-"`
	Avatar      string         `json:"avatar"`
	SystemRole  string 			`gorm:"type:user_system_role;default:'staff'" json:"system_role"`
	Preference  datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"preference"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`

	// Relations
	Sessions   []Session   `gorm:"foreignKey:UserID" json:"sessions,omitempty"`
	UserTeams  []UserTeam  `gorm:"foreignKey:UserID" json:"user_teams,omitempty"`
	Invitations []UserTeam `gorm:"foreignKey:InvitedBy" json:"invitations,omitempty"`
}




// ---
type NewOwnerPayload struct {
    DisplayName string          `json:"display_name"`
    Email string                `json:"email"`
    Password string             `json:"password"`
}


type UpdateAccountNamePayload struct {
    DisplayName string          `json:"display_name"`
}
type UpdateAccountPasswordPayload struct {
    Password string             `json:"password"`
    NewPassword string          `json:"new_password"`
}


type LoginPayload struct {
    Email string                `json:"email"`
    Password string             `json:"password"`
}


type AuthUser struct {
    ID int64                       `json:"id"`
    DisplayName string          `json:"display_name"`
    Email string                `json:"email"`
    Avatar string               `json:"avatar"`
    SystemRole string           `json:"system_role"`
    IsActive bool               `json:"is_active"`
    Preference map[string]any   `json:"preference"`
}


type SessionWithUser struct {
    Session Session
    AuthUser    AuthUser
}