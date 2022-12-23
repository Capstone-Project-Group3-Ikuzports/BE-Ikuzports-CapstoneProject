package repository

import (
	"errors"
	"ikuzports/features/club"

	"gorm.io/gorm"
)

type clubRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) club.RepositoryInterface {
	return &clubRepository{
		db: db,
	}
}

// Create implements club.RepositoryInterface
func (repo *clubRepository) Create(input club.Core, id int) error {
	clubGorm := fromCore(input)
	tx := repo.db.Create(&clubGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	var idClub int
	ty := repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&idClub)
	if ty.Error != nil {
		return errors.New("failed create data")
	}
	status := "Owner"
	tz := repo.db.Exec("insert into club_members (user_id, club_id, status) Values(?, ?, ?)", id, idClub, status)
	if tz.Error != nil {
		return errors.New("failed create data")
	}
	return nil
}

// GetAll implements club.RepositoryInterface
func (repo *clubRepository) GetAll() (data []club.Core, err error) {
	var club []Club

	tx := repo.db.Find(&club)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(club)

	return dataCore, nil
}

// GetAllWithSearch implements club.RepositoryInterface
func (repo *clubRepository) GetAllWithSearch(queryName string, queryCity string, queryCategoryID int) (data []club.Core, err error) {
	var club []Club

	tx := repo.db.Where("name LIKE ?", "%"+queryName+"%").Where(&Club{City: queryCity, CategoryID: uint(queryCategoryID)}).Find(&club) //masih salah
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(club)

	return dataCore, nil
}

// GetById implements club.RepositoryInterface
func (repo *clubRepository) GetById(id int) (data club.Core, err error) {
	var club Club

	tx := repo.db.First(&club, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = club.toCore()
	return dataCore, nil
}

// Update implements club.RepositoryInterface
func (repo *clubRepository) Update(input club.Core, id int) error {
	clubGorm := fromCore(input)
	tx := repo.db.Where("id= ?", id).Updates(clubGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}

// Delete implements club.RepositoryInterface
func (repo *clubRepository) Delete(id int) error {
	var club Club
	tx := repo.db.Delete(&club, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
