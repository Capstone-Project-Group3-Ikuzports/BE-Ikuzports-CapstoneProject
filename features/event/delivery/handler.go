package delivery

import (
	"errors"
	"ikuzports/features/event"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventDelivery struct {
	eventService event.ServiceInterface
}

func New(service event.ServiceInterface, e *echo.Echo) {
	handler := &EventDelivery{
		eventService: service,
	}
	e.POST("/events", handler.Create, middlewares.JWTMiddleware())
	e.GET("/events", handler.GetAll)
	e.GET("/events/:id", handler.GetByID, middlewares.JWTMiddleware())
}

func (delivery *EventDelivery) Create(c echo.Context) error {
	eventInput := EventRequest{}
	errBind := c.Bind(&eventInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(eventInput)

	file, _ := c.FormFile("image_event")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "image_event")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.ImageEvent = res
	} else {
		dataCore.ImageEvent = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}
	dataCore.UserID = middlewares.ExtractTokenUserId(c)

	err := delivery.eventService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Succes create data"))
}

func (delivery *EventDelivery) GetAll(c echo.Context) error {
	queryCategoryID, _ := strconv.Atoi(c.QueryParam("category_id"))
	queryCity := c.QueryParam("city")
	queryStatus := c.QueryParam("status")
	queryPage, _ := strconv.Atoi(c.QueryParam("page"))

	helper.LogDebug("\n isi queryCategoryID = ", queryCategoryID)
	helper.LogDebug("\n isi queryCity = ", queryCity)
	helper.LogDebug("\n isi queryStatus= ", queryStatus)

	result, err := delivery.eventService.GetAll(queryCategoryID, queryPage, queryCity, queryStatus)

	dataResp := fromCoreList(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", dataResp))
}

func (delivery *EventDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.eventService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCore(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read event", dataResp))
}

func (delivery *EventDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := delivery.eventService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
