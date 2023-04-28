package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Type              string  
	SoldProductAmount int `gorm:"default:0"`
	Products []Product    
}