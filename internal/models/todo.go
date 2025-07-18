package models

import (
	"github.com/google/uuid"
)

type Todos struct {
	BaseModel
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedById uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedBy   User      `gorm:"not null" json:"created_by"`
	UpdatedById uuid.UUID `gorm:"type:uuid;not null" json:"updated_by_id"`
	UpdatedBy   User      `gorm:"not null" json:"updated_by"`
}
