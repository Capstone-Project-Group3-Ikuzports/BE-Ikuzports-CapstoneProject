package repository

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	UserID  uint
	ClubID  uint
	Message string
}
