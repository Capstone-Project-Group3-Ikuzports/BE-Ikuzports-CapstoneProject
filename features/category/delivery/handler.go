package delivery

import (
	"ikuzports/features/category"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryDelivery struct {
	categoryService category.ServiceInterface
}

func New(service category.ServiceInterface, e *echo.Echo) {
	handler := &CategoryDelivery{
		categoryService: service,
	}
	e.GET("/categories", handler.GetAll, middlewares.JWTMiddleware())
}

func (delivery *CategoryDelivery) GetAll(c echo.Context) error {
	res, err := delivery.categoryService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all category data", res))
}
