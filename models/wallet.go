package models

import "github.com/google/uuid"

// Wallet is ...
type Wallet struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()", gorm:"primary_key"`
	UserID uuid.UUID
	User   `gorm:"foreignkey:UserID", gorm:"references:ID"`
}
