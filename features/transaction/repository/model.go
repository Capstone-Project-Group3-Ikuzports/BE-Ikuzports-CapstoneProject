package repository

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint
	TotalQuantity   int
	TotalPrice      int
	ProductID       uint
	PaymentMethod   string
	TransactionID   uint
	StatusPayment   string
	VirtualAccount  string
	TransactionTime time.Time
}
