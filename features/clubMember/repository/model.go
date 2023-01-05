package repository

import (
	"ikuzports/features/clubMember"

	"gorm.io/gorm"
)

type ClubMember struct {
	gorm.Model
	UserID uint
	User   User
	ClubID uint
	Status string
}

type User struct {
	gorm.Model
	Name         string
	Gender       string
	UserImage    string
	Phone_number string
	ClubMember   []ClubMember
}

func fromCore(dataCore clubMember.Core) ClubMember {
	return ClubMember{
		UserID: dataCore.UserID,
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}
func (dataCore *ClubMember) toCore() clubMember.Core {
	return clubMember.Core{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		User: clubMember.User{
			ID:           dataCore.User.ID,
			Name:         dataCore.User.Name,
			Gender:       dataCore.User.Gender,
			Phone_number: dataCore.User.Phone_number,
		},
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}

func toCoreList(dataModel []ClubMember) []clubMember.Core {
	var dataCore []clubMember.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
