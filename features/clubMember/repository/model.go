package repository

import "gorm.io/gorm"

type ClubMember struct {
	gorm.Model
	UserID uint
	ClubID uint
	Status string
}
