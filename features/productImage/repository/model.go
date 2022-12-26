package repository

import (
	"ikuzports/features/productImage"

	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	Url       string
	ProductID uint
	Product   Product
}

type Product struct {
	gorm.Model
	Name         string
	ProductImage []ProductImage
}

func fromCore(dataModel productImage.ProductImageCore) ProductImage {
	productImageGorm := ProductImage{
		Url:       dataModel.URL,
		ProductID: dataModel.ProductID,
	}
	return productImageGorm
}

func (dataModel *ProductImage) toCore() productImage.ProductImageCore {
	return productImage.ProductImageCore{
		ID:        dataModel.ID,
		URL:       dataModel.Url,
		ProductID: dataModel.ProductID,
	}
}

func toCoreList(dataModel []ProductImage) []productImage.ProductImageCore {
	var dataCore []productImage.ProductImageCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
