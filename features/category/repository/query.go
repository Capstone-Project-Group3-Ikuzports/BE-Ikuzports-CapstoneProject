package repository

import (
	"ikuzports/features/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) category.RepositoryInterface {
	return &categoryRepository{
		db: db,
	}
}

func (repo *categoryRepository) GetAll() (data []category.CategoryCore, err error) {
	var category []Category

	tx := repo.db.Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(category)

	return dataCore, nil
}
