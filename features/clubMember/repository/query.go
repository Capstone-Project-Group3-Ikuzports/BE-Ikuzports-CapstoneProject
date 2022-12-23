package repository

import (
	"errors"
	"fmt"
	"ikuzports/features/clubMember"

	"gorm.io/gorm"
)

type clubMemberRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) clubMember.RepositoryInterface {
	return &clubMemberRepository{
		db: db,
	}
}

// Create implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) Create(input clubMember.Core, id int) error {
	memberGorm := fromCore(input)
	tx := repo.db.Create(&memberGorm)
	fmt.Println(tx)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) GetAll() (data []clubMember.Core, err error) {
	var member []ClubMember

	tx := repo.db.Find(&member)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(member)

	return dataCore, nil
}

// GetById implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) GetById(id int) (data clubMember.Core, err error) {
	var member ClubMember

	tx := repo.db.First(&member, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = member.toCore()
	return dataCore, nil
}
