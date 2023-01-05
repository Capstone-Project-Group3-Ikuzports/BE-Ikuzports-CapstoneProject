package service

import (
	"errors"
	"ikuzports/features/product"
	"ikuzports/features/transaction"
	"ikuzports/features/user"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"

	"github.com/labstack/gommon/log"
)

type transactionService struct {
	transactionRepository transaction.RepositoryInterface
	productRepository     product.RepositoryInterface
	userRepository        user.RepositoryInterface
}

func New(repo transaction.RepositoryInterface, repoProduct product.RepositoryInterface, repoUser user.RepositoryInterface) transaction.ServiceInterface {
	return &transactionService{
		transactionRepository: repo,
		productRepository:     repoProduct,
		userRepository:        repoUser,
	}
}

func (service *transactionService) Create(input transaction.TransactionCore) (data transaction.MidtransCore, err error) {
	dataProduct, errProduct := service.productRepository.GetByID(int(input.ProductID))
	if errProduct != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	dataUser, errUser := service.userRepository.GetById(int(input.UserID))
	if errUser != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	orderID := "Order-" + helper.CreateRandomCode(7)
	dataMidtr := thirdparty.PaymentCoreApi(orderID, input, dataProduct, dataUser)
	if dataMidtr.TransactionID == "" {
		return data, errors.New("payment rejected by midtrans")
	}

	input.TransactionID = dataMidtr.TransactionID
	input.StatusPayment = dataMidtr.TransactionStatus
	input.OrderID = orderID

	for _, v := range dataMidtr.VaNumbers {
		input.VirtualAccount = v.VANumber
	}

	_, err = service.transactionRepository.Create(input)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	MidtrResp := transaction.MidtransCore{
		OrderID:       dataMidtr.OrderID,
		GrossAmt:      dataMidtr.GrossAmount,
		StatusMessage: dataMidtr.StatusMessage,
	}

	for _, v := range dataMidtr.VaNumbers {
		MidtrResp.VANumbers.Bank = v.Bank
		MidtrResp.VANumbers.VANumber = v.VANumber
	}

	_, errDel := service.productRepository.Delete(int(input.ProductID))
	if errDel != nil {
		return data, errDel
	}

	return MidtrResp, nil
}

func (service *transactionService) GetAll() (data []transaction.TransactionCore, err error) {
	data, err = service.transactionRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

func (service *transactionService) GetByID(id int) (data transaction.TransactionCore, err error) {
	data, err = service.transactionRepository.GetByID(id)
	if err != nil {
		log.Error(err.Error())
		return transaction.TransactionCore{}, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

func (service *transactionService) Update(input transaction.TransactionCore) (err error) {
	_, err = service.transactionRepository.Update(input)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
