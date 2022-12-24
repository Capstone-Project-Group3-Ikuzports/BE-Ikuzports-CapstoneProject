package delivery

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClubDelivery struct {
	clubService club.ServiceInterface
}

func New(service club.ServiceInterface, e *echo.Echo) {
	handler := &ClubDelivery{
		clubService: service,
	}

	e.GET("/clubs", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/clubs/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/clubs", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/clubs/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/clubs/:id", handler.Delete, middlewares.JWTMiddleware())
	// e.GET("/clubs/:id/chats", handler.GetChats, middlewares.JWTMiddleware())
	// e.GET("/clubs/:id/galleries", handler.GetGalleries, middlewares.JWTMiddleware())
	// e.GET("/clubs/:id/activities", handler.GetActivities, middlewares.JWTMiddleware())

}

func (delivery *ClubDelivery) Create(c echo.Context) error {
	clubInput := InsertRequest{}
	errBind := c.Bind(&clubInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	dataCore := toCore(clubInput)
	file, _ := c.FormFile("logo")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "logo")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.Logo = res
	} else {
		dataCore.Logo = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfA4x4hFqzMJRG8mkELzikjEXLgNu-ImEzEA&usqp=CAU"
	}

	err := delivery.clubService.Create(dataCore, userId)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot empty. Details : "+err.Error()))
		}
		if strings.Contains(err.Error(), "Please pick another email.") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert data "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ClubDelivery) GetAll(c echo.Context) error {
	queryName := c.QueryParam("name")
	queryCity := c.QueryParam("city")
	queryCategory := c.QueryParam("category")

	helper.LogDebug("\n\n\nULALA")

	// debug cek query param masuk
	helper.LogDebug("\n isi queryName = ", queryName)
	helper.LogDebug("\n isi queryCity= ", queryCity)
	helper.LogDebug("\n isi queryCategory = ", queryCategory)

	results, err := delivery.clubService.GetAll(queryName, queryCity, queryCategory)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *ClubDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.clubService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *ClubDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	clubInput := InsertRequest{}
	errBind := c.Bind(&clubInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	dataCore := toCore(clubInput)
	file, _ := c.FormFile("logo")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "logo")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.Logo = res
	} else {
		dataCore.Logo = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfA4x4hFqzMJRG8mkELzikjEXLgNu-ImEzEA&usqp=CAU"
	}

	err := delivery.clubService.Update(dataCore, id, userId)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *ClubDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	err := delivery.clubService.Delete(id, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
