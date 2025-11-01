package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64         `gorm:"primaryKey;autoIncrement:false" json:"id,string"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // optional soft delete
}
