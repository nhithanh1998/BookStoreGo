package models

import (
	"time"

	"github.com/google/uuid"
)

// Book is ...
type Book struct {
	ID            uuid.UUID `gorm:"type:uuid; default=uuid_generate_v4()", gorm:"primary_key"`
	Name          string    `gorm:"unique", gorm:"not null"`
	AuthorID      uuid.UUID
	Author        Author `gorm:"foreignkey:AuthorID", gorm:"references:ID`
	PublisherID   uuid.UUID
	Publisher     Publisher `gorm:"foreignkey:PublisherID", gorm:"references:ID"`
	BookSize      int       `gorm:"default:0"`
	SKU           int64
	BookCoverType string
	BookTotalPage int    `gorm:"default:0"`
	description   string `gorm:"default:No description"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
