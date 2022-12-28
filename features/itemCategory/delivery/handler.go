package delivery

import (
	itemcategory "ikuzports/features/itemCategory"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemCategoryDelivery struct {
	itemCategoryService itemcategory.ServiceInterface
}

func New(service itemcategory.ServiceInterface, e *echo.Echo) {
	handler := &ItemCategoryDelivery{
		itemCategoryService: service,
	}
	e.GET("/itemcategories", handler.GetAll, middlewares.JWTMiddleware())
}

func (delivery *ItemCategoryDelivery) GetAll(c echo.Context) error {
	res, err := delivery.itemCategoryService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(res)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all item category data", dataResp))
}
