package repository

import (
	_chat "ikuzports/features/chat/repository"
	_club "ikuzports/features/club/repository"
	"ikuzports/features/clubMember"
	_member "ikuzports/features/clubMember/repository"
	"ikuzports/features/event"
	_event "ikuzports/features/event/repository"
	"ikuzports/features/product"
	_product "ikuzports/features/product/repository"
	_transaction "ikuzports/features/transaction/repository"
	"ikuzports/features/user"

	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	Name             string `validate:"required"`
	Email            string `validate:"required,email,unique"`
	Password         string `validate:"required"`
	PhoneNumber      string `validate:"required"`
	Address          string
	City             string
	UserImage        string
	Gender           string
	Biodata          string
	Club             []_member.ClubMember
	EventParticipant []_event.EventParticipant
	Event            []_event.Event
	Transaction      []_transaction.Transaction
	Product          []_product.Product
	Chat             []_chat.Chat
}

type ClubMember struct {
	_member.ClubMember
	Club _club.Club
}

type Club struct {
	_club.Club
	Member []_member.ClubMember
}

type Event struct {
	_event.Event
}

type Transaction struct {
	_transaction.Transaction
}

type Product struct {
	_product.Product
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore user.Core) User {
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
func (dataModel *User) toCore() user.Core {
	return user.Core{
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
func toCoreList(dataModel []User) []user.Core {
	var dataCore []user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *ClubMember) toCoreClub() clubMember.Core {
	return clubMember.Core{
		ID:     dataModel.ID,
		UserID: dataModel.UserID,
		ClubID: dataModel.ClubID,
		Club: clubMember.Club{
			ID:           dataModel.Club.ID,
			Name:         dataModel.Club.Name,
			Category:     dataModel.Club.Category.Name,
			City:         dataModel.Club.City,
			JoinedMember: uint(dataModel.Club.JoinedMember),
			MemberTotal:  uint(dataModel.Club.MemberTotal),
			Logo:         dataModel.Club.Logo,
		},
		Status: dataModel.Status,
	}
}

func toClubList(dataModel []ClubMember) []clubMember.Core {
	var dataCore []clubMember.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreClub())
	}
	return dataCore
}

func (dataModel *Event) toCoreEvent() event.EventCore {
	return event.EventCore{
		ID:               dataModel.ID,
		Name:             dataModel.Name,
		Address:          dataModel.Address,
		City:             dataModel.City,
		CategoryID:       dataModel.CategoryID,
		StartDate:        dataModel.StartDate,
		EndDate:          dataModel.EndDate,
		TotalParticipant: dataModel.TotalParticipant,
		User: event.User{
			Name: dataModel.User.Name,
		},
	}
}

func toEventList(dataModel []Event) []event.EventCore {
	var dataCore []event.EventCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreEvent())
	}
	return dataCore
}

func (dataModel *Product) toCoreProduct() product.ProductCore {
	return product.ProductCore{
		ID:             dataModel.ID,
		Name:           dataModel.Name,
		Price:          uint(dataModel.Price),
		Description:    dataModel.Description,
		UserID:         dataModel.UserID,
		ItemCategoryID: dataModel.ItemCategoryID,
		ItemCategory: product.ItemCategory{
			Name: dataModel.ItemCategory.Name,
		},
	}
}

func toProductList(dataModel []Product) []product.ProductCore {
	var dataCore []product.ProductCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreProduct())
	}
	return dataCore
}

// func (dataModel *Transaction) toCoreTransaction() user.Transaction {
// 	return user.Transaction{
// 		ID:             dataModel.ID,
// 		UserID:         dataModel.UserID,
// 		TotalQuantity:  dataModel.TotalQuantity,
// 		TotalPrice:     dataModel.TotalPrice,
// 		ProductID:      dataModel.ProductID,
// 		PaymentMethod:  dataModel.PaymentMethod,
// 		TransactionID:  dataModel.TransactionID,
// 		StatusPayment:  dataModel.StatusPayment,
// 		VirtualAccount: dataModel.VirtualAccount,
// 	}
// }

// func toTransactionList(dataModel []Transaction) []user.Transaction {
// 	var dataCore []user.Transaction
// 	for _, v := range dataModel {
// 		dataCore = append(dataCore, v.toCoreTransaction())
// 	}
// 	return dataCore
// }
