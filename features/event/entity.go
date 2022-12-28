package event

import (
	"time"
)

type EventCore struct {
	ID               uint
	Name             string `validate:"required"`
	UserID           int
	User             User
	Address          string    `validate:"required"`
	City             string    `validate:"required"`
	CategoryID       uint      `validate:"required"`
	StartDate        time.Time `validate:"required"`
	EndDate          time.Time `validate:"required"`
	TotalParticipant int
	ImageEvent       string
	Status           string
	Category         Category
	MaximumPeople    int
}

type User struct {
	ID   uint
	Name string
}

type Category struct {
	ID   uint
	Name string
}

type RepositoryInterface interface {
	GetAll() (data []EventCore, err error)
	GetAllFilter(queryCategoryID int, queryCity, queryStatus string) (data []EventCore, err error)
	Create(input EventCore) (row int, err error)
	GetByID(id int) (data EventCore, err error)
	Delete(id int) (row int, err error)
	GetDate() (data []EventCore, err error)
	UpdateStatus(id int, status string) (rows int, err error)
	UpdateTotal(id int) (rows int, err error)
	GetLastID() (id int, err error)
}

type ServiceInterface interface {
	GetAll(queryCategoryID int, queryCity, queryStatus string) (data []EventCore, err error)
	Create(input EventCore) (err error)
	GetByID(id int) (data EventCore, err error)
	Delete(id int) (err error)
	// UpdateStatus() (err error)
}
