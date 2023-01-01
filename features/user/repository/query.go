package repository

import (
	"errors"
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	"ikuzports/features/transaction"
	"ikuzports/features/user"
	"ikuzports/utils/helper"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) error {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements user.Repository
func (repo *userRepository) GetAll() (data []user.Core, err error) {
	var users []User

	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(users)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *userRepository) GetAllWithSearch(query string) (data []user.Core, err error) {
	var users []User

	tx := repo.db.Where("full_name LIKE ?", "%"+query+"%").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(users)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *userRepository) GetById(id int) (data user.Core, err error) {
	var user User

	tx := repo.db.First(&user, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = user.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *userRepository) Update(input user.Core, id int) error {
	userGorm := fromCore(input)
	var user User
	tx := repo.db.Model(&user).Where("ID = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *userRepository) Delete(id int) error {
	var user User
	tx := repo.db.Delete(&user, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *userRepository) FindUser(email string) (result user.Core, err error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	result = userData.toCore()

	return result, nil
}

// GetClubs implements user.RepositoryInterface
func (repo *userRepository) GetClubs(id int) (data []clubMember.Core, err error) {
	var club []ClubMember
	tx := repo.db.Preload("Club").Preload("Club.Category").Where("user_id = ?", id).Find(&club)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toClubList(club)
	return dataCore, nil
}

// GetEvents implements user.RepositoryInterface
func (repo *userRepository) GetEvents(id int) (data []event.EventCore, err error) {
	var event []Event
	tx := repo.db.Preload("Category").Preload("User").Where("user_id = ?", id).Find(&event)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toEventList(event)
	return dataCore, nil
}

// GetProducts implements user.RepositoryInterface
func (repo *userRepository) GetProducts(id int) (data []user.ProductCore, err error) {
	var products []Product
	tx := repo.db.Preload("User").Preload("ProductImage").Preload("ItemCategory").Where("user_id = ?", id).Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	helper.LogDebug(products[1].ProductImage)
	helper.LogDebug("==========")
	var dataCore = toProductList(products)
	return dataCore, nil
}

// GetTransactions implements user.RepositoryInterface
func (repo *userRepository) GetTransactions(id int) (data []transaction.TransactionCore, err error) {
	var transaction []Transaction
	tx := repo.db.Where("user_id = ?", id).Find(&transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toTransactionList(transaction)
	return dataCore, nil
}
