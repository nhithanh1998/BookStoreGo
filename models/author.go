package models

import (
	"time"

	"github.com/google/uuid"
)

// Author is ...
type Author struct {
	ID          uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(), gorm:"primary_key"`
	Name        string    `gorm:"unique", gorm:"not null"`
	Nationality string
	Gender      bool   `gorm:"default:false"`
	Biography   string `gorm:"default:No biography yet"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
