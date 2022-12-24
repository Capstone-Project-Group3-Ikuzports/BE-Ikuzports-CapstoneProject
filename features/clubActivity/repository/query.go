package repository

import (
	"errors"
	"ikuzports/features/clubActivity"

	"gorm.io/gorm"
)

type clubActivityRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) clubActivity.RepositoryInterface {
	return &clubActivityRepository{
		db: db,
	}
}

// Create implements clubActivity.RepositoryInterface
func (repo *clubActivityRepository) Create(input clubActivity.Core, id int) error {
	activityGorm := fromCore(input)
	tx := repo.db.Create(&activityGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert club failed")
	}
	return nil
}

// GetAll implements clubActivity.RepositoryInterface
func (repo *clubActivityRepository) GetAll() (data []clubActivity.Core, err error) {
	var activity []ClubActivity

	tx := repo.db.Find(&activity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(activity)
	return dataCore, nil
}

// GetAllByClubId implements clubActivity.RepositoryInterface
func (repo *clubActivityRepository) GetAllByClubId(queryClubId uint) (data []clubActivity.Core, err error) {
	var activity []ClubActivity

	tx := repo.db.Where("club_id = ?", queryClubId).Find(&activity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(activity)
	return dataCore, nil
}

// Update implements clubActivity.RepositoryInterface
func (repo *clubActivityRepository) Update(input clubActivity.Core, id int) error {
	activityGorm := fromCore(input)
	tx := repo.db.Where("id= ?", id).Updates(activityGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update club failed")
	}
	return nil
}
