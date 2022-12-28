package repository

import (
	itemcategory "ikuzports/features/itemCategory"

	"gorm.io/gorm"
)

type itemCategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) itemcategory.RepositoryInterface {
	return &itemCategoryRepository{
		db: db,
	}
}

func (repo *itemCategoryRepository) GetAll() (data []itemcategory.ItemCategoryCore, err error) {
	var itemCategory []ItemCategory

	tx := repo.db.Find(&itemCategory)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(itemCategory)

	return dataCore, nil
}
