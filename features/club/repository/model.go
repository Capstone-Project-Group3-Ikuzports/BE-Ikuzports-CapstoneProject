package repository

import (
	_chat "ikuzports/features/chat/repository"
	_activity "ikuzports/features/clubActivity/repository"
	_member "ikuzports/features/clubMember/repository"
	_galery "ikuzports/features/galery/repository"
	"time"

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
	ClubMember   []_member.ClubMember
	Chat         []_chat.Chat
	ClubActivity []_activity.ClubActivity
	Galery       []_galery.Galery
	Aggreement   Aggreement
}

type Aggreement struct {
	ClubID         uint
	TermsCondition string
	CreatedAt      time.Time
	UpdateAt       time.Time
	DeletedAt      time.Time
}
