package user

import (
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	"ikuzports/features/transaction"
	"time"
)

type Core struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required"`
	Password    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Address     string
	City        string
	UserImage   string
	Gender      string
	LoginMethod string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GoogleCore struct {
	Email   string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Picture string `json:"picture" form:"picture"`
}

type ClubMember struct {
	clubMember.Core
	Club Club
}

type Club struct {
	club.Core
}
type Event struct {
	event.EventCore
}

type Transaction struct {
	transaction.TransactionCore
}

type ProductCore struct {
	ID             uint
	Name           string
	Price          uint
	Description    string
	UserID         int
	User           Core
	ItemCategoryID uint
	ItemCategory   ItemCategory
	City           string
	Thumbnail      string
	ProductImage   []ProductImage
}

type ItemCategory struct {
	ID   uint
	Name string
}

type ProductImage struct {
	ID  uint
	URL string
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetClubs(id int) (data []clubMember.Core, err error)
	GetProducts(id int) (data []ProductCore, err error)
	GetEvents(id int) (data []event.EventCore, err error)
	GetTransactions(id int) (data []transaction.TransactionCore, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
	GetClubs(id int) (data []clubMember.Core, err error)
	GetProducts(id int) (data []ProductCore, err error)
	GetEvents(id int) (data []event.EventCore, err error)
	GetTransactions(id int) (data []transaction.TransactionCore, err error)
}
