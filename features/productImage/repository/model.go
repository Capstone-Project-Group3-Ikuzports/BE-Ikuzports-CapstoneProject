package repository

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	Url       string
	ProductID uint
}
