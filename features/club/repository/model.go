package repository

import (
	_chat "ikuzports/features/chat/repository"
	"ikuzports/features/club"
	_activity "ikuzports/features/clubActivity/repository"
	_member "ikuzports/features/clubMember/repository"
	_galery "ikuzports/features/galery/repository"

	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name         string
	Address      string
	City         string
	CategoryID   uint
	Description  string
	Logo         string
	JoinedMember int
	MemberTotal  int
	Rule         string
	Requirement  string
	ClubMember   []_member.ClubMember
	Chat         []_chat.Chat
	ClubActivity []_activity.ClubActivity
	Galery       []_galery.Galery
	Category     Category
}

type ClubMember struct {
	gorm.Model
	UserID uint
	ClubID uint
	Status string
}
type Category struct {
	gorm.Model
	Name string
	Club []Club
}

func fromCore(dataCore club.Core) Club {
	clubGorm := Club{
		Name:         dataCore.Name,
		Address:      dataCore.Address,
		City:         dataCore.City,
		CategoryID:   dataCore.CategoryID,
		Description:  dataCore.Description,
		Logo:         dataCore.Logo,
		JoinedMember: dataCore.JoinedMember,
		MemberTotal:  dataCore.MemberTotal,
		Rule:         dataCore.Rule,
		Requirement:  dataCore.Requirement,
	}
	return clubGorm
}
func (dataCore *Club) toCore() club.Core {
	dataModel := club.Core{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Address:      dataCore.Address,
		City:         dataCore.City,
		CategoryID:   dataCore.CategoryID,
		Description:  dataCore.Description,
		Logo:         dataCore.Logo,
		JoinedMember: dataCore.JoinedMember,
		MemberTotal:  dataCore.MemberTotal,
		Rule:         dataCore.Rule,
		Requirement:  dataCore.Requirement,
		Category: club.Category{
			ID:   dataCore.Category.ID,
			Name: dataCore.Category.Name,
		},
	}
	return dataModel
}

func toCoreList(dataModel []Club) []club.Core {
	var dataCore []club.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataCore *ClubMember) toCoreMember() club.Status {
	return club.Status{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}

func toCoreMemberList(dataModel []ClubMember) []club.Status {
	var dataCore []club.Status
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreMember())
	}
	return dataCore
}
