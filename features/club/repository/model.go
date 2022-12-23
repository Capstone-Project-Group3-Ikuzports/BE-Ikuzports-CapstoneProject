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
}

func fromCore(dataCore club.Core) Club {
	clubGorm := Club{
		Name:        dataCore.Name,
		Address:     dataCore.Address,
		City:        dataCore.City,
		CategoryID:  dataCore.CategoryID,
		Description: dataCore.Description,
		Logo:        dataCore.Logo,
		MemberTotal: dataCore.MemberTotal,
		Rule:        dataCore.Rule,
		Requirement: dataCore.Requirement,
	}
	return clubGorm
}
func (dataCore *Club) toCore() club.Core {
	return club.Core{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Address:     dataCore.Address,
		City:        dataCore.City,
		CategoryID:  dataCore.CategoryID,
		Description: dataCore.Description,
		Logo:        dataCore.Logo,
		MemberTotal: dataCore.MemberTotal,
		Rule:        dataCore.Rule,
		Requirement: dataCore.Requirement,
		CreatedAt:   dataCore.CreatedAt,
		UpdateAt:    dataCore.UpdatedAt,
	}
}

func toCoreList(dataModel []Club) []club.Core {
	var dataCore []club.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
