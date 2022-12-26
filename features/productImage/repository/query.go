package repository

import (
	"ikuzports/features/productImage"

	"gorm.io/gorm"
)

type productImageRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) productImage.RepositoryInterface {
	return &productImageRepository{
		db: db,
	}
}

func (repo *productImageRepository) Create(data productImage.ProductImageCore) (rows int, err error) {
	imageProductGorm := fromCore(data)
	tx := repo.db.Create(&imageProductGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *productImageRepository) GetAll() (data []productImage.ProductImageCore, err error) {
	var productImage []ProductImage
	tx := repo.db.Find(&productImage)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(productImage)

	return dataCore, nil
}

func (repo *productImageRepository) GetByID(id int) (data productImage.ProductImageCore, err error) {
	var productImage ProductImage
	tx := repo.db.First(&productImage, id)
	if tx.Error != nil {
		return data, tx.Error
	}

	dataCore := productImage.toCore()

	return dataCore, nil
}
