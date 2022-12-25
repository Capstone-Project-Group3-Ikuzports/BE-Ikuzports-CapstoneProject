package repository

import (
	"errors"
	"ikuzports/features/galery"

	"gorm.io/gorm"
)

type galeryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) galery.RepositoryInterface {
	return &galeryRepository{
		db: db,
	}
}

// GetAll implements galery.RepositoryInterface
func (repo *galeryRepository) GetAll() (data []galery.Core, err error) {
	var image []Galery

	// tx := repo.db.Find(&club)
	tx := repo.db.Find(&image)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(image)

	return dataCore, nil
}

// Create implements galery.RepositoryInterface
func (repo *galeryRepository) Create(input galery.Core, id int) error {
	imageGorm := fromCore(input)

	tx := repo.db.Create(&imageGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetById implements galery.RepositoryInterface
func (repo *galeryRepository) GetById(id int) (data galery.Core, err error) {
	var images Galery

	tx := repo.db.Find(&images, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = images.toCore()
	return dataCore, nil
}

// Delete implements galery.RepositoryInterface
func (repo *galeryRepository) Delete(id int) error {
	var galery Galery
	tx := repo.db.Delete(&galery, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// Update implements galery.RepositoryInterface
func (repo *galeryRepository) Update(input galery.Core, id int) error {
	galeryGorm := fromCore(input)
	tx := repo.db.Where("id= ?", id).Updates(galeryGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update club failed")
	}
	return nil
}
