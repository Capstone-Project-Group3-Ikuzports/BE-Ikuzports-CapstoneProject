package user

import (
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	"ikuzports/features/product"
	"ikuzports/features/transaction"
	"time"
)

type Core struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Address     string
	City        string
	UserImage   string
	Gender      string
	Biodata     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

type Product struct {
	product.ProductCore
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetClubs(id int) (data []clubMember.Core, err error)
	GetProducts(id int) (data []product.ProductCore, err error)
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
	GetProducts(id int) (data []product.ProductCore, err error)
	GetEvents(id int) (data []event.EventCore, err error)
	GetTransactions(id int) (data []transaction.TransactionCore, err error)
}
