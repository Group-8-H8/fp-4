package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model   
	FullName  string    
	Email     string `gorm:"uniqueIndex"`
	Password  string    
	Role      string `gorm:"default:'customer'"`
	Balance   int `gorm:"default:0"`
	Transactions []Transaction       
}