package repository

import (
	"time"

	"gorm.io/gorm"
)

type ClubActivity struct {
	gorm.Model
	ClubID         uint
	Name           string
	StartHour      time.Time
	EndHour        time.Time
	Day            string
	Location       string
	ActivityDetail string
}
