package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	Product Product
	User User
}

type TransactionProduct struct {
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
}

type TransactionProductUser struct {
	User
	Product
	Transaction
}