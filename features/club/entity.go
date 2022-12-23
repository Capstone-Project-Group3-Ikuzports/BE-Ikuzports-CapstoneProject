package club

import "time"

type Core struct {
	ID          uint
	Name        string
	Address     string
	City        string
	CategoryID  uint
	Description string
	Logo        string
	MemberTotal int
	Rule        string
	Requirement string
	Aggreement  Aggreement
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type Aggreement struct {
	ClubID         uint
	TermsCondition string
}

type ServiceInterface interface {
	GetAll(queryName, queryCity, queryCategory string) (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
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
	// GetChats(id int) (data []Club, err error)
	// GetGaleries(id int) (data []Product, err error)
	// GetActivities(id int) (data []Event, err error)
}
