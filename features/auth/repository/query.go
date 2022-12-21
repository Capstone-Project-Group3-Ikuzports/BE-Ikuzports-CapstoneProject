package repository

import (
	"errors"
	"ikuzports/features/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email string) (result auth.Core, err error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return auth.Core{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return auth.Core{}, errors.New("login failed")
	}

	result = userData.toCore()

	return result, nil
}
