package models

import (
	"time"

	"github.com/google/uuid"
)

// User is ...
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()", gorm:"primary_key"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
