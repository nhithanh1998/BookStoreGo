package models

import (
	"time"

	"github.com/google/uuid"
)

// Publisher is ...
type Publisher struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()", gorm:"primary_key"`
	Name      string    `gorm:"unique", gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
