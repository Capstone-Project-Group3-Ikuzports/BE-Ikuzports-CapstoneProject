package delivery

import "ikuzports/features/transaction"

type TransactionResp struct {
	ID              uint   `json:"id"`
	UserID          int    `json:"user_id"`
	TotalPrice      int    `json:"total_price"`
	TotalQuantity   int    `json:"total_quantity"`
	ProductID       uint   `json:"product_id"`
	TransactionID   string `json:"transaction_id"`
	StatusPayment   string `json:"status_payment"`
	VirtualAccount  string `json:"virtual_account"`
	TransactionTime string `json:"transaction_time"`
	OrderID         string `json:"order_id"`
}

type MidtransResp struct {
	OrderID       string            `json:"order_id"`
	GrossAmt      string            `json:"gross_amount"`
	StatusMessage string            `json:"status_message"`
	VANumbers     VANumbersResponse `json:"va_numbers"`
}

type VANumbersResponse struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

func fromCore(dataCore transaction.TransactionCore) TransactionResp {
	return TransactionResp{
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

func fromCoreList(dataCore []transaction.TransactionCore) []TransactionResp {
	var dataResponse []TransactionResp
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreMidtr(dataCore transaction.MidtransCore) MidtransResp {
	return MidtransResp{
		OrderID:       dataCore.OrderID,
		GrossAmt:      dataCore.GrossAmt,
		StatusMessage: dataCore.StatusMessage,
		VANumbers:     fromCoreVA(dataCore.VANumbers),
	}
}

func fromCoreVA(dataCore transaction.VANumbersCore) VANumbersResponse {
	return VANumbersResponse{
		Bank:     dataCore.Bank,
		VANumber: dataCore.VANumber,
	}
}
