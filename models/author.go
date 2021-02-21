package models

import "gorm.io/gorm"

// Author is ...
type Author struct {
	gorm.Model
	Name        string
	Nationality string
	Gender      bool
	Biography   string
}
