package repository

import "gorm.io/gorm"

type Galery struct {
	gorm.Model
	Url     string
	ClubID  uint
	Caption string
}
