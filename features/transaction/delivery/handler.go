package delivery

import (
	"ikuzports/features/transaction"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionDelivery struct {
	transactionService transaction.ServiceInterface
}

func New(service transaction.ServiceInterface, e *echo.Echo) {
	handler := &TransactionDelivery{
		transactionService: service,
	}
	e.POST("/transactions", handler.Create, middlewares.JWTMiddleware())
	e.POST("/transactions/notif", handler.Update)
	e.GET("/transactions", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/transactions/:id", handler.GetByID, middlewares.JWTMiddleware())
}

func (delivery *TransactionDelivery) Create(c echo.Context) error {
	transactionInput := TransactionRequest{}
	errBind := c.Bind(&transactionInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(transactionInput)

	dataCore.UserID = middlewares.ExtractTokenUserId(c)

	dataMidtr, err := delivery.transactionService.Create(dataCore)

	dataResp := fromCoreMidtr(dataMidtr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create data", dataResp))
}

func (delivery *TransactionDelivery) GetAll(c echo.Context) error {
	res, err := delivery.transactionService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed get all transaction data "+err.Error()))
	}

	dataResp := fromCoreList(res)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all transaction data", dataResp))
}

func (delivery *TransactionDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := delivery.transactionService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed get all transaction data "+err.Error()))
	}

	dataResp := fromCore(res)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get transaction data by id", dataResp))
}

func (delivery *TransactionDelivery) Update(c echo.Context) error {
	notifInput := NotificationRequest{}
	errBind := c.Bind(&notifInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	dataCore := toCoreNotif(notifInput)

	errGet := delivery.transactionService.Update(dataCore)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Notification received"))
}
