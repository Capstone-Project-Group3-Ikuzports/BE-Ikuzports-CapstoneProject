package repository

import (
	_club "ikuzports/features/club/repository"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name             string
	UserID           uint
	Address          string
	City             string
	CategoryID       uint
	StartDate        time.Time
	EndDate          time.Time
	TotalParticipant int
	Status           string
	EventMember      []EventParticipant
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
