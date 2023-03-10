package repository

import (
	_chat "ikuzports/features/chat/repository"
	_club "ikuzports/features/club/repository"
	"ikuzports/features/clubMember"
	_member "ikuzports/features/clubMember/repository"
	"ikuzports/features/event"
	_event "ikuzports/features/event/repository"
	_product "ikuzports/features/product/repository"
	"ikuzports/features/transaction"
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
	UserImage        string
	Gender           string
	LoginMethod      string
	ClubMember       []_member.ClubMember
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
	Member   []_member.ClubMember
	Category Category
}

type Event struct {
	_event.Event
}

type Transaction struct {
	_transaction.Transaction
}

type Product struct {
	gorm.Model
	Name           string
	Price          int
	Description    string
	UserID         int
	User           User
	ItemCategoryID uint
	City           string
	Thumbnail      string
	ItemCategory   ItemCategory
	Transaction    []_transaction.Transaction
	ProductImage   []ProductImage
}
type ItemCategory struct {
	gorm.Model
	Name    string
	Product []Product
}

type Category struct {
	gorm.Model
	Name string
	Club []Club
}

type ProductImage struct {
	gorm.Model
	URL       string
	ProductID uint
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
		UserImage:   dataCore.UserImage,
		Gender:      dataCore.Gender,
		LoginMethod: dataCore.LoginMethod,
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
		UserImage:   dataModel.UserImage,
		Gender:      dataModel.Gender,
		LoginMethod: dataModel.LoginMethod,
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
			CategoryID:   dataModel.Club.CategoryID,
			Category:     dataModel.Club.Category.Name,
			City:         dataModel.Club.City,
			Logo:         dataModel.Club.Logo,
			JoinedMember: dataModel.Club.JoinedMember,
			MemberTotal:  dataModel.Club.MemberTotal,
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

// func (dataModel *ClubMember) toCoreClub() club.Core {
// 	var arrMember []club.ClubMember
// 	for _, val := range dataModel.Member {
// 		arrMember = append(arrMember, club.ClubMember{
// 			ID:     val.ID,
// 			UserID: val.UserID,
// 			ClubID: val.ClubID,
// 			Status: val.Status,
// 		})
// 	}
// 	return club.Core{
// 		ID:   dataModel.Club.ID,
// 		Name: dataModel.Club.Name,
// 		Category: club.Category{
// 			Name: dataModel.Category.Name,
// 		},
// 		City:         dataModel.Club.City,
// 		JoinedMember: dataModel.Club.JoinedMember,
// 		MemberTotal:  dataModel.Club.MemberTotal,
// 		Logo:         dataModel.Club.Logo,
// 		Member:       arrMember,
// 	}
// }

// func toClubList(dataModel []Club) []club.Core {
// 	var dataCore []club.Core
// 	for _, v := range dataModel {
// 		dataCore = append(dataCore, v.toCoreClub())
// 	}
// 	return dataCore
// }

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
		Category: event.Category{
			Name: dataModel.Category.Name,
		},
		ImageEvent: dataModel.ImageEvent,
		Status:     dataModel.Status,
	}
}

func toEventList(dataModel []Event) []event.EventCore {
	var dataCore []event.EventCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreEvent())
	}
	return dataCore
}

func (dataModel *Product) toCoreProduct() user.ProductCore {
	return user.ProductCore{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Price:       uint(dataModel.Price),
		Description: dataModel.Description,
		UserID:      dataModel.UserID,
		User: user.Core{
			Name: dataModel.User.Name,
		},
		ItemCategoryID: dataModel.ItemCategoryID,
		ItemCategory: user.ItemCategory{
			Name: dataModel.ItemCategory.Name,
		},
		City:         dataModel.City,
		Thumbnail:    dataModel.Thumbnail,
		ProductImage: toProductCoreListImage(dataModel.ProductImage),
	}
}

func toProductList(dataModel []Product) []user.ProductCore {
	var dataCore []user.ProductCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreProduct())
	}
	return dataCore
}
func (dataModel *ProductImage) toCoreProductImage() user.ProductImage {
	return user.ProductImage{
		ID:  dataModel.ID,
		URL: dataModel.URL,
	}
}

func toProductCoreListImage(dataModel []ProductImage) []user.ProductImage {
	var dataCoreImage []user.ProductImage
	for _, v := range dataModel {
		dataCoreImage = append(dataCoreImage, v.toCoreProductImage())
	}
	return dataCoreImage
}

func (dataModel *Transaction) toCoreTransaction() transaction.TransactionCore {
	return transaction.TransactionCore{
		ID:              dataModel.ID,
		UserID:          dataModel.UserID,
		TotalQuantity:   dataModel.TotalQuantity,
		TotalPrice:      dataModel.TotalPrice,
		ProductID:       dataModel.ProductID,
		TransactionID:   dataModel.TransactionID,
		StatusPayment:   dataModel.StatusPayment,
		VirtualAccount:  dataModel.VirtualAccount,
		TransactionTime: dataModel.TransactionTime.String(),
		OrderID:         dataModel.OrderID,
	}
}

func toTransactionList(dataModel []Transaction) []transaction.TransactionCore {
	var dataCore []transaction.TransactionCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreTransaction())
	}
	return dataCore
}
