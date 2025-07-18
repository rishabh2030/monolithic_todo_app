package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
	JoinedAt time.Time `gorm:"default:current_timestamp" json:"joined_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.JoinedAt.IsZero() {
		u.JoinedAt = time.Now()
	}
	return
}
