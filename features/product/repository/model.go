package repository

import (
	"ikuzports/features/product"
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
	ProductImage   []ProductImage
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

type ProductImage struct {
	gorm.Model
	URL       string
	ProductID uint
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
		City:           dataModel.City,
		ProductImage:   toCoreListImage(dataModel.ProductImage),
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

func (dataModel *ProductImage) toCoreProductImage() product.ProductImage {
	return product.ProductImage{
		ID:  dataModel.ID,
		URL: dataModel.URL,
	}
}

func toCoreListImage(dataModel []ProductImage) []product.ProductImage {
	var dataCoreImage []product.ProductImage
	for _, v := range dataModel {
		dataCoreImage = append(dataCoreImage, v.toCoreProductImage())
	}
	return dataCoreImage
}

func toCoreList(dataModel []Product) []product.ProductCore {
	var dataCore []product.ProductCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
