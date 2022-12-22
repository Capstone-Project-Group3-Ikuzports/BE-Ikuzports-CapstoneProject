package repository

import (
	"ikuzports/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	City        string
	UserImage   string
	Gender      string
	Biodata     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

//DTO

func (dataModel User) toCore() auth.Core {
	return auth.Core{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		PhoneNumber: dataModel.PhoneNumber,
		Address:     dataModel.Address,
		City:        dataModel.City,
		Gender:      dataModel.Gender,
		UserImage:   dataModel.UserImage,
		Biodata:     dataModel.Biodata,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}
