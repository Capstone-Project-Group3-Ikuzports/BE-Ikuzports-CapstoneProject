package repository

import (
	"time"

	"gorm.io/gorm"
)

type ClubActivity struct {
	gorm.Model
	ClubID         uint
	Name           string
	StartDate      time.Time
	EndDtae        time.Time
	Location       string
	ActivityDetail string
}
