package repository

import (
	"errors"
	"ikuzports/features/transaction"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.RepositoryInterface {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) Create(input transaction.TransactionCore) (rows int, err error) {
	transactionGorm := fromCore(input)
	tx := repo.db.Create(&transactionGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}

	return int(tx.RowsAffected), nil
}

func (repo *transactionRepository) GetAll() (data []transaction.TransactionCore, err error) {
	var transaction []Transaction
	tx := repo.db.Order("updated_at desc").Find(&transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(transaction)
	return dataCore, nil
}

func (repo *transactionRepository) GetByID(id int) (data transaction.TransactionCore, err error) {
	var transaction Transaction
	tx := repo.db.First(&transaction, id)
	if tx.Error != nil {
		return data, tx.Error
	}

	var dataCore = transaction.toCore()
	return dataCore, nil
}

func (repo *transactionRepository) GetByOrderID(orderID string) (data transaction.TransactionCore, err error) {
	var transaction Transaction
	tx := repo.db.Where("order_id = ?", orderID).First(&transaction)
	if tx.Error != nil {
		return data, tx.Error
	}

	var dataCore = transaction.toCore()
	return dataCore, nil
}

func (repo *transactionRepository) Update(input transaction.TransactionCore) (rows int, err error) {
	transactionGorm := fromCore(input)
	var transaction Transaction

	tx := repo.db.Model(&transaction).Where("order_id = ?", transactionGorm.OrderID).Updates(&transactionGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("update failed")
	}
	return int(tx.RowsAffected), nil
}
