package models

import "gorm.io/gorm"

// Genre is ...
type Genre struct {
	gorm.Model
	Name string
}
