package repository

import (
	"ikuzports/features/transaction"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          int
	TotalPrice      int
	TotalQuantity   int
	ProductID       uint
	PaymentMethod   string
	TransactionID   string
	StatusPayment   string
	VirtualAccount  string
	TransactionTime time.Time
	OrderID         string
}

func fromCore(dataModel transaction.TransactionCore) Transaction {
	transactionGorm := Transaction{
		UserID:          dataModel.UserID,
		TotalPrice:      dataModel.TotalPrice,
		TotalQuantity:   dataModel.TotalQuantity,
		ProductID:       dataModel.ProductID,
		TransactionID:   dataModel.TransactionID,
		StatusPayment:   dataModel.StatusPayment,
		VirtualAccount:  dataModel.VirtualAccount,
		TransactionTime: timeParsing(dataModel.TransactionTime),
		OrderID:         dataModel.OrderID,
	}
	return transactionGorm
}

func (dataModel *Transaction) toCore() transaction.TransactionCore {
	return transaction.TransactionCore{
		ID:              dataModel.ID,
		UserID:          dataModel.UserID,
		TotalPrice:      dataModel.TotalPrice,
		TotalQuantity:   dataModel.TotalQuantity,
		ProductID:       dataModel.ProductID,
		TransactionID:   dataModel.TransactionID,
		StatusPayment:   dataModel.StatusPayment,
		VirtualAccount:  dataModel.VirtualAccount,
		TransactionTime: dataModel.TransactionTime.String(),
		OrderID:         dataModel.OrderID,
	}
}

func toCoreList(dataModel []Transaction) []transaction.TransactionCore {
	var dataCore []transaction.TransactionCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func timeParsing(date string) (dateParsed time.Time) {
	layoutFormat := "2006-01-02 15:04:05"

	dateParsed, _ = time.Parse(layoutFormat, date)

	return dateParsed
}
