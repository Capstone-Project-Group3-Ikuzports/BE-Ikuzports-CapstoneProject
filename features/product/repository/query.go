package repository

import (
	"errors"
	"ikuzports/features/product"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.RepositoryInterface {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) GetAll() (data []product.ProductCore, err error) {
	var product []Product
	tx := repo.db.Preload("User").Preload("ItemCategory").Order("updated_at desc").Find(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(product)

	return dataCore, nil

}

func (repo *productRepository) GetAllFilter(queryItemCategoryID int, queryCity, queryName string) (data []product.ProductCore, err error) {
	var product []Product
	tx := repo.db.Where("name LIKE ?", "%"+queryName+"%").Where(&Product{ItemCategoryID: uint(queryItemCategoryID), City: queryCity}).Preload("User").Preload("ItemCategory").Order("updated_at desc").Find(&product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(product)

	return dataCore, nil

}

func (repo *productRepository) Create(input product.ProductCore) (row int, err error) {
	productGorm := fromCore(input)
	tx := repo.db.Create(&productGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *productRepository) GetByID(id int) (data product.ProductCore, err error) {
	var product Product
	tx := repo.db.Preload("User").Preload("ItemCategory").Preload("ProductImage").First(&product, id)
	if tx.Error != nil {
		return data, tx.Error
	}

	dataCore := product.toCore()

	return dataCore, nil

}

func (repo *productRepository) Update(id int, input product.ProductCore) (rows int, err error) {
	productGorm := fromCore(input)
	var product Product

	tx := repo.db.Model(&product).Where("ID = ?", id).Updates(&productGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("update failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *productRepository) Delete(id int) (rows int, err error) {
	var product Product
	tx := repo.db.Delete(&product, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed")
	}
	return int(tx.RowsAffected), nil
}

// GetImages implements product.RepositoryInterface
func (repo *productRepository) GetImages(id int) (data []product.ProductImage, err error) {
	var images []ProductImage
	tx := repo.db.Where("product_id = ?", id).Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreListImage(images)
	return dataCore, nil
}
