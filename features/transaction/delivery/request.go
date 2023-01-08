package delivery

import "ikuzports/features/transaction"

type TransactionRequest struct {
	TotalPrice      int    `json:"total_price" form:"total_price"`
	ProductID       uint   `json:"product_id" form:"product_id"`
	PaymentMethod   string `json:"payment_method" form:"payment_method"`
	TotalQuantity   int    `json:"product_quantity" form:"product_quantity"`
	TransactionTime string `json:"transaction_time" form:"transaction_time"`
}

type NotificationRequest struct {
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	OrderID           string `json:"order_id" form:"order_id"`
}

func toCore(transactionInput TransactionRequest) transaction.TransactionCore {
	transactionCoreData := transaction.TransactionCore{
		TotalPrice:      transactionInput.TotalPrice,
		ProductID:       transactionInput.ProductID,
		PaymentMethod:   transactionInput.PaymentMethod,
		TotalQuantity:   transactionInput.TotalQuantity,
		TransactionTime: transactionInput.TransactionTime,
	}
	return transactionCoreData
}

func toCoreNotif(notifInput NotificationRequest) transaction.TransactionCore {
	transactionCoreData := transaction.TransactionCore{
		StatusPayment: notifInput.TransactionStatus,
		OrderID:       notifInput.OrderID,
	}
	return transactionCoreData
}
