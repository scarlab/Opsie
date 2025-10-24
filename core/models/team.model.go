package models

type Team struct {
	BaseModel

	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"uniqueIndex;not null" json:"slug"`
	Description string `json:"description"`
	Icon        string `json:"icon"`

	UserTeams []UserTeam `gorm:"foreignKey:TeamID" json:"user_teams,omitempty"`
	Projects  []Project  `gorm:"foreignKey:TeamID" json:"projects,omitempty"`
	Resources []Resource `gorm:"foreignKey:TeamID" json:"resources,omitempty"`
}


type TeamWithMeta struct {
	Team
	IsDefault bool `json:"is_default"`
	IsAdmin bool 	`json:"is_admin"`
}


type NewTeamPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Icon            string      `json:"icon"`
}

type UpdateTeamPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Icon            string      `json:"icon"`
}
