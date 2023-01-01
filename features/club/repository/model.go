package repository

import (
	"ikuzports/features/chat"
	_chat "ikuzports/features/chat/repository"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	_activity "ikuzports/features/clubActivity/repository"
	_members "ikuzports/features/clubMember"
	_member "ikuzports/features/clubMember/repository"
	"ikuzports/features/galery"
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
	JoinedMember uint
	MemberTotal  uint
	Rule         string
	Requirement  string
	ClubMember   []_member.ClubMember
	Chat         []_chat.Chat
	ClubActivity []_activity.ClubActivity
	Galery       []_galery.Galery
	Category     Category
}

type ClubMember struct {
	_member.ClubMember
}

type Category struct {
	gorm.Model
	Name string
	Club []Club
}

type Chat struct {
	_chat.Chat
}

type Galery struct {
	_galery.Galery
}

type ClubActivity struct {
	_activity.ClubActivity
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
		User: club.User{
			Name: dataCore.User.Name,
		},
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}

func (dataCore *ClubMember) toCoreMembers() _members.Core {
	return _members.Core{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		User: _members.User{
			Name:         dataCore.User.Name,
			Gender:       dataCore.User.Gender,
			Phone_number: dataCore.User.Phone_number,
		},
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}

func toCoreMembersList(dataModel []ClubMember) []_members.Core {
	var dataCore []_members.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreMembers())
	}
	return dataCore
}

func (dataCore *Chat) toCoreChat() chat.Core {
	return chat.Core{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		User: chat.User{
			ID:   dataCore.User.ID,
			Name: dataCore.User.Name,
		},
		ClubID:    dataCore.ClubID,
		Message:   dataCore.Message,
		CreatedAt: dataCore.CreatedAt,
	}
}

func toCoreChatList(dataModel []Chat) []chat.Core {
	var dataCore []chat.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreChat())
	}
	return dataCore
}

func (dataCore *Galery) toCoreGalery() galery.Core {
	return galery.Core{
		ID:      dataCore.ID,
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}

func toCoreGaleryList(dataModel []Galery) []galery.Core {
	var dataCore []galery.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreGalery())
	}
	return dataCore
}

func (dataCore *ClubActivity) toCoreActivity() clubActivity.Core {
	return clubActivity.Core{
		ID:             dataCore.ID,
		ClubID:         dataCore.ClubID,
		Name:           dataCore.Name,
		StartTime:      dataCore.StartTime,
		EndTime:        dataCore.EndTime,
		Day:            dataCore.Day,
		Location:       dataCore.Location,
		ActivityDetail: dataCore.ActivityDetail,
	}
}

func toCoreActivityList(dataModel []ClubActivity) []clubActivity.Core {
	var dataCore []clubActivity.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreActivity())
	}
	return dataCore
}
