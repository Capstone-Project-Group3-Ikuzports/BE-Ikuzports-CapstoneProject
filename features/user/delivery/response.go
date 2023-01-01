package delivery

import (
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	_event "ikuzports/features/event/delivery"
	_product "ikuzports/features/product/delivery"
	"ikuzports/features/transaction"
	_transaction "ikuzports/features/transaction/delivery"
	"ikuzports/features/user"

	"time"
)

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	UserImage   string `json:"user_image"`
	Gender      string `json:"gender"`
}
type ClubResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	CategoryID   uint   `json:"category_id"`
	Category     string `json:"category"`
	City         string `json:"city"`
	JoinedMember uint   `json:"joined_member"`
	Logo         string `json:"logo"`
	MemberTotal  int    `json:"member_total"`
	Status       string `json:"status"`
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	TotalQuantity   int       `json:"total_quantity"`
	TotalPrice      int       `json:"total_price"`
	ProductID       uint      `json:"product_id"`
	PaymentMethod   string    `json:"payment_method"`
	TransactionID   uint      `json:"transaction_id"`
	StatusPayment   string    `json:"status_payment"`
	VirtualAccount  string    `json:"virtual_account"`
	TransactionTime time.Time `json:"transaction_time"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		PhoneNumber: dataCore.PhoneNumber,
		Address:     dataCore.Address,
		UserImage:   dataCore.UserImage,
		Gender:      dataCore.Gender,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromClub(dataCore clubMember.Core) ClubResponse {
	return ClubResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Club.Name,
		CategoryID:   dataCore.Club.CategoryID,
		Category:     dataCore.Club.Category,
		City:         dataCore.Club.City,
		JoinedMember: dataCore.Club.JoinedMember,
		MemberTotal:  int(dataCore.Club.MemberTotal),
		Logo:         dataCore.Club.Logo,
		Status:       dataCore.Status,
	}
}

// func fromClub(dataCore club.Core) ClubResponse {
// 	var arrMember []_member.MemberResponse
// 	for _, val := range dataCore.Member {
// 		arrMember = append(arrMember, _member.MemberResponse{
// 			ID:     val.ID,
// 			UserID: val.UserID,
// 			ClubID: val.ClubID,
// 			Status: val.Status,
// 		})
// 	}
// 	return ClubResponse{
// 		ID:           dataCore.ID,
// 		Name:         dataCore.Name,
// 		Category:     dataCore.Category.Name,
// 		City:         dataCore.City,
// 		JoinedMember: uint(dataCore.JoinedMember),
// 		MemberTotal:  int(dataCore.MemberTotal),
// 		Logo:         dataCore.Logo,
// 		Member:       arrMember,
// 	}
// }

func fromClubList(dataCore []clubMember.Core) []ClubResponse {
	var dataResponse []ClubResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromClub(v))
	}
	return dataResponse
}

func fromEvent(dataCore event.EventCore) _event.EventResponse {
	return _event.EventResponse{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Address:          dataCore.Address,
		City:             dataCore.City,
		OrganizerName:    dataCore.User.Name,
		CategoryName:     dataCore.Category.Name,
		StartDate:        dataCore.StartDate,
		EndDate:          dataCore.EndDate,
		TotalParticipant: uint(dataCore.TotalParticipant),
		ImageEvent:       dataCore.ImageEvent,
		Status:           dataCore.Status,
	}
}

func fromEventList(dataCore []event.EventCore) []_event.EventResponse {
	var dataResponse []_event.EventResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromEvent(v))
	}
	return dataResponse
}

func fromProduct(dataCore user.ProductCore) _product.ProductResponse {
	return _product.ProductResponse{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Price:            dataCore.Price,
		Description:      dataCore.Description,
		UserName:         dataCore.User.Name,
		ItemCategoryName: dataCore.ItemCategory.Name,
		City:             dataCore.City,
		ProductImage:     fromCoreListImage(dataCore.ProductImage),
	}
}

func fromProductList(dataCore []user.ProductCore) []_product.ProductResponse {
	var dataResponse []_product.ProductResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromProduct(v))
	}
	return dataResponse
}

func fromCoreImage(data user.ProductImage) _product.ProductImage {
	return _product.ProductImage{
		ID:  data.ID,
		Url: data.URL,
	}
}

func fromCoreListImage(dataCore []user.ProductImage) []_product.ProductImage {
	var dataResponse []_product.ProductImage
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreImage(v))
	}
	return dataResponse
}

func fromTransaction(dataCore transaction.TransactionCore) _transaction.TransactionResp {
	return _transaction.TransactionResp{
		ID:              dataCore.ID,
		UserID:          dataCore.UserID,
		TotalPrice:      dataCore.TotalPrice,
		TotalQuantity:   dataCore.TotalQuantity,
		ProductID:       dataCore.ProductID,
		TransactionID:   dataCore.TransactionID,
		StatusPayment:   dataCore.StatusPayment,
		VirtualAccount:  dataCore.VirtualAccount,
		TransactionTime: dataCore.TransactionTime,
		OrderID:         dataCore.OrderID,
	}
}

func fromTransactionList(dataCore []transaction.TransactionCore) []_transaction.TransactionResp {
	var dataResponse []_transaction.TransactionResp
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromTransaction(v))
	}
	return dataResponse
}
