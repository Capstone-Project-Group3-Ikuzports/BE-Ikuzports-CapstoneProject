package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	FullName        string `valiidate:"required"`
	Email           string `valiidate:"required,email"`
	Password        string `valiidate:"required"`
	Phone           string `valiidate:"required"`
	Gender          string
	ProfileImageUrl string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, c echo.Context) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
}
