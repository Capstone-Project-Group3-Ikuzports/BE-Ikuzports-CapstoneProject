package delivery

import (
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
	Biodata     string `json:"biodata"`
}
type ClubResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Category    uint   `json:"category"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	MemberTotal int    `json:"member_total"`
}
type EventResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	UserID           uint      `json:"user_id"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	CategoryID       uint      `json:"category_id"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	TotalParticipant int       `json:"total_participants"`
	OrganizerName    string    `json:"organizer_name"`
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

type ProductResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Quantity       int    `json:"quantity"`
	Description    string `json:"description"`
	UserID         uint   `json:"user_id"`
	ItemCategoryID uint   `json:"item_category"`
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
		Biodata:     dataCore.Biodata,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromClub(dataCore user.Club) ClubResponse {
	return ClubResponse{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Address:     dataCore.Address,
		City:        dataCore.City,
		Category:    dataCore.CategoryID,
		Description: dataCore.Description,
		Logo:        dataCore.Logo,
		MemberTotal: dataCore.MemberTotal,
	}
}

func fromClubList(dataCore []user.Club) []ClubResponse {
	var dataResponse []ClubResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromClub(v))
	}
	return dataResponse
}

func fromEvent(dataCore user.Event) EventResponse {
	return EventResponse{
		ID:               dataCore.ID,
		UserID:           dataCore.UserID,
		Name:             dataCore.Name,
		Address:          dataCore.Address,
		City:             dataCore.City,
		CategoryID:       dataCore.CategoryID,
		StartDate:        dataCore.StartDate,
		EndDate:          dataCore.EndDate,
		TotalParticipant: dataCore.TotalParticipant,
		OrganizerName:    dataCore.OrganizerName,
	}
}

func fromEventList(dataCore []user.Event) []EventResponse {
	var dataResponse []EventResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromEvent(v))
	}
	return dataResponse
}

func fromProduct(dataCore user.Product) ProductResponse {
	return ProductResponse{
		ID:             dataCore.ID,
		Name:           dataCore.Name,
		Price:          dataCore.Price,
		Quantity:       dataCore.Quantity,
		Description:    dataCore.Description,
		UserID:         dataCore.UserID,
		ItemCategoryID: dataCore.ItemCategoryID,
	}
}

func fromProductList(dataCore []user.Product) []ProductResponse {
	var dataResponse []ProductResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromProduct(v))
	}
	return dataResponse
}

func fromTransaction(dataCore user.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:              dataCore.ID,
		UserID:          dataCore.UserID,
		TotalQuantity:   dataCore.TotalQuantity,
		TotalPrice:      dataCore.TotalPrice,
		ProductID:       dataCore.ProductID,
		PaymentMethod:   dataCore.PaymentMethod,
		TransactionID:   dataCore.TransactionID,
		StatusPayment:   dataCore.StatusPayment,
		VirtualAccount:  dataCore.VirtualAccount,
		TransactionTime: dataCore.TransactionTime,
	}
}

func fromTransactionList(dataCore []user.Transaction) []TransactionResponse {
	var dataResponse []TransactionResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromTransaction(v))
	}
	return dataResponse
}
