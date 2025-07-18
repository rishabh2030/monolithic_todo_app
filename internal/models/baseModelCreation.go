package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	IsDeletedAt bool      `gorm:"default:false" json:"is_deleted_at"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}
