package repository

import (
	"ikuzports/features/clubMember"

	"gorm.io/gorm"
)

type ClubMember struct {
	gorm.Model
	UserID uint
	ClubID uint
	Status string
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
