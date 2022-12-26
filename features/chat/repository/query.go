package repository

import (
	"errors"
	"ikuzports/features/chat"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) chat.RepositoryInterface {
	return &chatRepository{
		db: db,
	}
}

// Create implements chat.RepositoryInterface
func (repo *chatRepository) Create(input chat.Core) error {
	chatGorm := fromCore(input)
	tx := repo.db.Create(&chatGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// Delete implements chat.RepositoryInterface
func (repo *chatRepository) Delete(id int) error {
	var chat Chat
	tx := repo.db.Delete(&chat, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// FindMember implements chat.RepositoryInterface
func (repo *chatRepository) FindMember(id int, idUser int) (data chat.Core, err error) {
	panic("unimplemented")
}

// GetAll implements chat.RepositoryInterface
func (repo *chatRepository) GetAll() (data []chat.Core, err error) {
	var chat []Chat

	tx := repo.db.Preload("User").Find(&chat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(chat)
	return dataCore, nil
}

// GetById implements chat.RepositoryInterface
func (repo *chatRepository) GetById(id int) (data chat.Core, err error) {
	var chat Chat

	tx := repo.db.Preload("User").Find(&chat, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = chat.toCore()
	return dataCore, nil
}
