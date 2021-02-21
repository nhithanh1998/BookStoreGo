package models

import "gorm.io/gorm"

// Book is ...
type Book struct {
	gorm.Model
	Name          string
	Author        Author
	Publisher     Publisher
	BookSize      int
	SKU           int64
	BookCoverType string
	BookTotalPage int
	description   string
}
