package user

import (
	"time"
)

type Core struct {
	ID          uint
	Name        string `valiidate:"required"`
	Email       string `valiidate:"required,email"`
	Password    string `valiidate:"required"`
	PhoneNumber string `valiidate:"required"`
	Address     string
	UserImage   string
	Gender      string
	Biodata     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Club struct {
	ID          uint
	Name        string
	Address     string
	City        string
	CategoryID  uint
	Description string
	Logo        string
	MemberTotal int
}

type Event struct {
	ID               uint
	Name             string
	UserID           uint
	Address          string
	City             string
	CategoryID       uint
	StartDate        time.Time
	EndDate          time.Time
	TotalParticipant int
	OrganizerName    string
}

type Transaction struct {
	ID              uint
	UserID          uint
	TotalQuantity   int
	TotalPrice      int
	ProductID       uint
	PaymentMethod   string
	TransactionID   uint
	StatusPayment   string
	VirtualAccount  string
	TransactionTime time.Time
}

type Product struct {
	ID             uint
	Name           string
	Price          int
	Quantity       int
	Description    string
	UserID         uint
	ItemCategoryID uint
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetClubs(id int) (data []Club, err error)
	GetProducts(id int) (data []Product, err error)
	GetEvents(id int) (data []Event, err error)
	GetTransactions(id int) (data []Transaction, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
	GetClubs(id int) (data []Club, err error)
	GetProducts(id int) (data []Product, err error)
	GetEvents(id int) (data []Event, err error)
	GetTransactions(id int) (data []Transaction, err error)
}
