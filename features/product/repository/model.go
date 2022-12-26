package repository

import (
	"ikuzports/features/product"
	_productImage "ikuzports/features/productImage/repository"
	_transaction "ikuzports/features/transaction/repository"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string
	Price          int
	Description    string
	UserID         int
	User           User
	ItemCategoryID uint
	City           string
	ItemCategory   ItemCategory
	Transaction    []_transaction.Transaction
	ProductImage   []_productImage.ProductImage
}

type ItemCategory struct {
	gorm.Model
	Name    string
	Product []Product
}

type User struct {
	gorm.Model
	Name string
	City string
}

func fromCore(dataModel product.ProductCore) Product {
	productGorm := Product{
		Name:           dataModel.Name,
		Price:          int(dataModel.Price),
		Description:    dataModel.Description,
		UserID:         dataModel.UserID,
		ItemCategoryID: dataModel.ItemCategoryID,
		City:           dataModel.City,
	}
	return productGorm
}

func (dataModel *Product) toCore() product.ProductCore {
	return product.ProductCore{
		ID:             dataModel.ID,
		Name:           dataModel.Name,
		Price:          uint(dataModel.Price),
		Description:    dataModel.Description,
		UserID:         dataModel.UserID,
		User:           dataModel.User.toCoreUser(),
		ItemCategoryID: dataModel.ItemCategoryID,
		ItemCategory:   dataModel.ItemCategory.toCoreItemCategory(),
	}
}

func (dataModel *ItemCategory) toCoreItemCategory() product.ItemCategory {
	return product.ItemCategory{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func (dataModel *User) toCoreUser() product.User {
	return product.User{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func toCoreList(dataModel []Product) []product.ProductCore {
	var dataCore []product.ProductCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
