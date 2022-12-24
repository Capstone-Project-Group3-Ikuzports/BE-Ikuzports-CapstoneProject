package club

import "time"

type Core struct {
	ID           uint
	Name         string `validate:"required"`
	Address      string
	City         string `validate:"required"`
	CategoryID   uint   `validate:"required"`
	Description  string
	Logo         string
	JoinedMember int
	MemberTotal  int `validate:"required"`
	Rule         string
	Requirement  string `validate:"required"`
	Category     Category
	CreatedAt    time.Time
	UpdateAt     time.Time
}
type Category struct {
	ID   uint
	Name string
}
type Status struct {
	ID     uint
	UserID uint
	ClubID uint
	Status string
}

type ServiceInterface interface {
	GetAll(queryName string, queryCity string, queryCategoryID int) (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, userId int) error
	Delete(id int, userId int) error
	// GetChats(id int) (data []Club, err error)
	// GetGaleries(id int) (data []Product, err error)
	// GetActivities(id int) (data []Event, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(queryName string, queryCity string, queryCategoryID int) (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetLastID() (id int, err error)
	UpdateMember(id int) (rows int, err error)
	// GetChats(id int) (data []Club, err error)
	// GetGaleries(id int) (data []Product, err error)
	// GetActivities(id int) (data []Event, err error)
	GetStatus(id int, userId int) (data Status, err error)
}
