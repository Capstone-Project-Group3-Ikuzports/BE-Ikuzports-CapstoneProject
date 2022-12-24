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
func (repo *clubMemberRepository) Create(input clubMember.Core) error {
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

// FindMember implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) FindMember(id int, idUser int) (data clubMember.Core, err error) {
	var member ClubMember

	tx := repo.db.Where("club_id = ?", id).Where("user_id = ?", idUser).First(&member)
	if tx.Error != nil {
		return data, tx.Error
	}
	var dataCore = member.toCore()

	return dataCore, nil
}

// UpdateMember implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) UpdateMember(id int) error {
	tx := repo.db.Exec("UPDATE clubs SET joined_member = (SELECT COUNT(user_id) FROM club_members WHERE club_id = ?) WHERE id = ?", id, id)
	if tx.Error != nil {
		return errors.New("failed update total_participant data")
	}

	if tx.RowsAffected == 0 {
		return errors.New("failed update total_participant data")
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

// GetAllByStatus implements clubMember.RepositoryInterface
func (repo *clubMemberRepository) GetAllByStatus(queryStatus string) (data []clubMember.Core, err error) {
	var member []ClubMember

	tx := repo.db.Where("status = ?", queryStatus).Find(&member)
	if tx.Error != nil {
		return data, tx.Error
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
