package models

import "gorm.io/gorm"

// Publisher is ...
type Publisher struct {
	gorm.Model
	Name string
}
