package delivery

import (
	"errors"
	"ikuzports/features/galery"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type GaleryDelivery struct {
	galeryService galery.ServiceInterface
}

func New(service galery.ServiceInterface, e *echo.Echo) {
	handler := &GaleryDelivery{
		galeryService: service,
	}

	e.GET("/galeries", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/galeries/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/galeries", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/galeries/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/galeries/:id", handler.Delete, middlewares.JWTMiddleware())

}

func (delivery *GaleryDelivery) GetAll(c echo.Context) error {

	results, err := delivery.galeryService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *GaleryDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.galeryService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *GaleryDelivery) Create(c echo.Context) error {
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
	file, _ := c.FormFile("url")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "url")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.Url = res
	} else {
		dataCore.Url = "https://trilogi.ac.id/universitas/wp-content/uploads/2017/07/dummy-img.png"
	}

	err := delivery.galeryService.Create(dataCore, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *GaleryDelivery) Update(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	clubInput := UpdateRequest{}
	errBind := c.Bind(&clubInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	dataCore := toCoreUpdate(clubInput)
	file, _ := c.FormFile("url")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "url")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.Url = res
	}

	err := delivery.galeryService.Update(dataCore, idParam, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data"))
}

func (delivery *GaleryDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	err := delivery.galeryService.Delete(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success delete image"))
}
