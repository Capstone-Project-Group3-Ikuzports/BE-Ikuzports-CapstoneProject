package thirdparty

import (
	"os"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	"ikuzports/features/product"
	"ikuzports/features/transaction"
	"ikuzports/features/user"
)

func OrderMidtrans(orderId string, price int64) *snap.Response {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)
	// orderId := "ORDER-103"

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: price,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapResp, _ := snap.CreateTransaction(req)
	return snapResp
}

func PaymentCoreApi(orderId string, transData transaction.TransactionCore, productdata product.ProductCore, userdata user.Core) *coreapi.ChargeResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	productID := strconv.Itoa(int(productdata.ID))

	req2 := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(transData.TotalPrice),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    productID,
				Name:  productdata.Name,
				Price: int64(productdata.Price),
				Qty:   int32(transData.TotalQuantity),
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: userdata.Name,
			Email: userdata.Email,
			Phone: userdata.PhoneNumber,
			BillAddr: &midtrans.CustomerAddress{
				Address: userdata.Address,
				City:    userdata.City,
			},
		},
		// CustomExpiry: &coreapi.CustomExpiry{
		// 	OrderTime:      orderTime,
		// 	ExpiryDuration: 60,
		// 	Unit:           "minute",
		// },
	}

	coreApiRes, _ := coreapi.ChargeTransaction(req2)
	return coreApiRes
}

func CheckMidtrans(orderId string) *coreapi.TransactionStatusResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	res, _ := c.CheckTransaction(orderId)
	return res
}
