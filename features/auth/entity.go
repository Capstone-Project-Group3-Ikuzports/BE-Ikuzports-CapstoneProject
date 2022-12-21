package auth

import "time"

type Core struct {
	ID              uint
	FullName        string
	Email           string `validate:"required,email"`
	Password        string `validate:"required"`
	Phone           string
	Gender          string
	ProfileImageUrl string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	Login(input Core) (data Core, token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result Core, err error)
}
