package repository

import (
	_user "ikuzports/features/user"
	"time"

	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	Name        string `validate:"required"`
	Email       string `validate:"required,email,unique"`
	Password    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Address     string
	City        string
	UserImage   string
	Gender      string
	Biodata     string
	Club        []ClubMember
	Event       []EventParticipant
	Transaction []Transaction
	Product     []Product
	Chat        []Chat
}

type ClubMember struct {
	gorm.Model
	UserID uint
	ClubID uint
	Status string
}

type Chat struct {
	gorm.Model
	UserID  uint
	ClubID  uint
	Message string
}

type Club struct {
	gorm.Model
	Name         string
	Address      string
	City         string
	CategoryID   uint
	Description  string
	Logo         string
	MemberTotal  int
	ClubMember   []ClubMember
	Chat         []Chat
	ClubActivity []ClubActivity
	Galery       []Galery
	Aggreement   []Aggreement
}

type ClubActivity struct {
	gorm.Model
	ClubID         uint
	Name           string
	StartDate      time.Time
	EndDtae        time.Time
	Location       string
	ActivityDetail string
}
type Galery struct {
	gorm.Model
	Url     string
	ClubID  uint
	Caption string
}

type Aggreement struct {
	ClubID         uint
	TermsCondition string
	CreatedAt      time.Time
	UpdateAt       time.Time
	DeletedAt      time.Time
}
type Category struct {
	gorm.Model
	Name   string
	Clubs  []Club
	Events []Event
}

type EventParticipant struct {
	gorm.Model
	UserID  uint
	EventID uint
	Status  string
}
type Event struct {
	gorm.Model
	Name             string
	UserID           uint
	Address          string
	City             string
	CategoryID       uint
	StartDate        time.Time
	EndDate          time.Time
	TotalParticipant int
	OrganizerName    string
	Status           string
	EventMember      []EventParticipant
}

type Transaction struct {
	gorm.Model
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
	gorm.Model
	Name           string
	Price          int
	Quantity       int
	Description    string
	UserID         uint
	ItemCategoryID uint
	Transaction    []Transaction
	ProductImage   []ProductImage
}

type ProductImage struct {
	gorm.Model
	Url       string
	ProductID uint
}

type ItemCategory struct {
	gorm.Model
	Name    string
	Product []Product
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _user.Core) User {
	userGorm := User{
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		Password:    dataCore.Password,
		PhoneNumber: dataCore.PhoneNumber,
		Address:     dataCore.Address,
		City:        dataCore.City,
		UserImage:   dataCore.UserImage,
		Gender:      dataCore.Gender,
		Biodata:     dataCore.Biodata,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		PhoneNumber: dataModel.PhoneNumber,
		Address:     dataModel.Address,
		City:        dataModel.City,
		UserImage:   dataModel.UserImage,
		Gender:      dataModel.Gender,
		Biodata:     dataModel.Biodata,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *Club) toCoreClub() _user.Club {
	return _user.Club{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Address:     dataModel.Address,
		City:        dataModel.City,
		CategoryID:  dataModel.CategoryID,
		Description: dataModel.Description,
		Logo:        dataModel.Logo,
		MemberTotal: dataModel.MemberTotal,
	}
}

func toClubList(dataModel []Club) []_user.Club {
	var dataCore []_user.Club
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreClub())
	}
	return dataCore
}

func (dataModel *Event) toCoreEvent() _user.Event {
	return _user.Event{
		ID:               dataModel.ID,
		Name:             dataModel.Name,
		Address:          dataModel.Address,
		City:             dataModel.City,
		CategoryID:       dataModel.CategoryID,
		StartDate:        dataModel.StartDate,
		EndDate:          dataModel.EndDate,
		TotalParticipant: dataModel.TotalParticipant,
		OrganizerName:    dataModel.OrganizerName,
	}
}

func toEventList(dataModel []Event) []_user.Event {
	var dataCore []_user.Event
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreEvent())
	}
	return dataCore
}

func (dataModel *Product) toCoreProduct() _user.Product {
	return _user.Product{
		ID:             dataModel.ID,
		Name:           dataModel.Name,
		Price:          dataModel.Price,
		Quantity:       dataModel.Quantity,
		Description:    dataModel.Description,
		UserID:         dataModel.UserID,
		ItemCategoryID: dataModel.ItemCategoryID,
	}
}

func toProductList(dataModel []Product) []_user.Product {
	var dataCore []_user.Product
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreProduct())
	}
	return dataCore
}

func (dataModel *Transaction) toCoreTransaction() _user.Transaction {
	return _user.Transaction{
		ID:             dataModel.ID,
		UserID:         dataModel.UserID,
		TotalQuantity:  dataModel.TotalQuantity,
		TotalPrice:     dataModel.TotalPrice,
		ProductID:      dataModel.ProductID,
		PaymentMethod:  dataModel.PaymentMethod,
		TransactionID:  dataModel.TransactionID,
		StatusPayment:  dataModel.StatusPayment,
		VirtualAccount: dataModel.VirtualAccount,
	}
}

func toTransactionList(dataModel []Transaction) []_user.Transaction {
	var dataCore []_user.Transaction
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreTransaction())
	}
	return dataCore
}
