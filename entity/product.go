package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title      string
	Price      int
	Stock      int
	CategoryID int
}