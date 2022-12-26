package delivery

import (
	"ikuzports/features/clubActivity"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClubActivityDelivery struct {
	clubActivityService clubActivity.ServiceInterface
}

func New(service clubActivity.ServiceInterface, e *echo.Echo) {
	handler := &ClubActivityDelivery{
		clubActivityService: service,
	}

	e.GET("/activities", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/activities", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/activities/:id", handler.Update, middlewares.JWTMiddleware())

}

func (delivery *ClubActivityDelivery) Create(c echo.Context) error {
	activityInput := InsertRequest{}
	errBind := c.Bind(&activityInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	dataCore := toCore(activityInput)
	err := delivery.clubActivityService.Create(dataCore, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ClubActivityDelivery) GetAll(c echo.Context) error {
	queryClubId, _ := strconv.Atoi(c.QueryParam("club_id"))
	results, err := delivery.clubActivityService.GetAll(queryClubId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	dataRespon := fromCoreList(results)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all activity", dataRespon))
}

func (delivery *ClubActivityDelivery) Update(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	inputData := InsertRequest{}
	errBind := c.Bind(&inputData)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}

	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	dataUpdateCore := toCore(inputData)
	err := delivery.clubActivityService.Update(dataUpdateCore, idParam, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success update activity"))
}
