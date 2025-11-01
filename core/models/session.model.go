package models

import (
	"time"
)

type Session struct {
	ID        int64    `gorm:"primaryKey" json:"id,string"`
	UserID    int64    `gorm:"not null;index" json:"user_id,string"`
	Key       string    `gorm:"uniqueIndex;not null" json:"key"`
	IP        string    `json:"ip"`
	OS        string    `json:"os"`
	Device    string    `json:"device"`
	Browser   string    `json:"browser"`
	IsValid   bool      `gorm:"default:true" json:"is_valid"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
