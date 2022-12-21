package repository

import (
	"ikuzports/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint
	FullName        string
	Email           string
	Password        string
	Phone           string
	Gender          string
	ProfileImageUrl string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

//DTO

func (dataModel User) toCore() auth.Core {
	return auth.Core{
		ID:              dataModel.ID,
		FullName:        dataModel.FullName,
		Email:           dataModel.Email,
		Password:        dataModel.Password,
		Phone:           dataModel.Phone,
		Gender:          dataModel.Gender,
		ProfileImageUrl: dataModel.ProfileImageUrl,
		Role:            dataModel.Role,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}
