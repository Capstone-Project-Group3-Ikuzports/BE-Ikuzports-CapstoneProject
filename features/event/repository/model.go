package repository

import (
	_club "ikuzports/features/club/repository"
	_eventCore "ikuzports/features/event"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name             string
	UserID           uint
	User             User
	Address          string
	City             string
	CategoryID       uint
	StartDate        time.Time
	EndDate          time.Time
	TotalParticipant int
	ImageEvent       string
	Status           string
	EventMember      []EventParticipant
	Category         Category
	MaximumPeople    int
	Description      string
}

type User struct {
	gorm.Model
	Name   string
	Events []Event
}

type EventParticipant struct {
	gorm.Model
	UserID  uint
	EventID uint
	Status  string
}

type Category struct {
	gorm.Model
	Name   string
	Clubs  []_club.Club
	Events []Event
}

func (dataModel *Event) toCore() _eventCore.EventCore {
	return _eventCore.EventCore{
		ID:               dataModel.ID,
		Name:             dataModel.Name,
		Address:          dataModel.Address,
		City:             dataModel.City,
		User:             dataModel.User.toCoreUser(),
		StartDate:        dataModel.StartDate,
		EndDate:          dataModel.EndDate,
		TotalParticipant: dataModel.TotalParticipant,
		ImageEvent:       dataModel.ImageEvent,
		Status:           dataModel.Status,
		Category:         dataModel.Category.toCoreCategory(),
		MaximumPeople:    dataModel.MaximumPeople,
		Description:      dataModel.Description,
	}
}

func fromCore(dataModel _eventCore.EventCore) Event {
	eventGorm := Event{
		Name:          dataModel.Name,
		UserID:        uint(dataModel.UserID),
		Address:       dataModel.Address,
		City:          dataModel.City,
		CategoryID:    dataModel.CategoryID,
		StartDate:     dataModel.StartDate,
		EndDate:       dataModel.EndDate,
		Status:        dataModel.Status,
		ImageEvent:    dataModel.ImageEvent,
		MaximumPeople: dataModel.MaximumPeople,
		Description:   dataModel.Description,
	}
	return eventGorm
}

func toCoreList(dataModel []Event) []_eventCore.EventCore {
	var dataCore []_eventCore.EventCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *Category) toCoreCategory() _eventCore.Category {
	return _eventCore.Category{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func (dataModel *User) toCoreUser() _eventCore.User {
	return _eventCore.User{
		Name: dataModel.Name,
	}
}
