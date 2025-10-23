package models

type Team struct {
	BaseModel

	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"uniqueIndex;not null" json:"slug"`
	Description string `json:"description"`
	Logo        string `json:"logo"`

	UserTeams []UserTeam `gorm:"foreignKey:TeamID" json:"user_teams,omitempty"`
	Projects  []Project  `gorm:"foreignKey:TeamID" json:"projects,omitempty"`
	Resources []Resource `gorm:"foreignKey:TeamID" json:"resources,omitempty"`
}



// ---

type NewTeamPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Logo            string      `json:"logo"`
}

type UpdateTeamPayload struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
}
