package repository

import (
	_productImage "ikuzports/features/productImage/repository"
	_transaction "ikuzports/features/transaction/repository"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string
	Price          int
	Description    string
	UserID         uint
	ItemCategoryID uint
	Transaction    []_transaction.Transaction
	ProductImage   []_productImage.ProductImage
}

type ItemCategory struct {
	gorm.Model
	Name    string
	Product []Product
}
